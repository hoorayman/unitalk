package chat

import (
	"context"
	"time"
	"unitalk/broker"
	"unitalk/logger"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	// Max wait time when writing message to peer
	writeWait = 10 * time.Second
	// Max time till next pong from peer
	pongWait = 60 * time.Second
	// Send ping interval, must be less then pong wait time
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 10000
)

var ctx = context.Background()

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client represents the websocket client at the server
type Client struct {
	clientID string
	// The actual websocket connection.
	conn *websocket.Conn
	room string
}

// NewClient constructor
func NewClient(conn *websocket.Conn, room string, clientID string) *Client {
	return &Client{
		clientID: clientID,
		conn:     conn,
		room:     room,
	}
}

// ReadPump method
func (client *Client) ReadPump() {
	defer func() {
		client.conn.Close()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// Start endless read loop, waiting for messages from client
	for {
		_, message, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Writer.Error(err.Error(), zap.String("ws", "unexpected close error"))
			}
			break
		}
		err = broker.REDIS.Publish(ctx, client.room, message).Err()
		if err != nil {
			logger.Writer.Error(err.Error(), zap.String("redis", "pub"))
		}
	}
}

// WritePump method
func (client *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	pubsub := broker.REDIS.Subscribe(ctx, client.room)
	defer func() {
		ticker.Stop()
		client.conn.Close()
		pubsub.Unsubscribe(ctx, client.room)
	}()
	receiveFromRoom := pubsub.Channel()

	for {
		select {
		case message, ok := <-receiveFromRoom:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write([]byte(message.Payload))

			// Attach queued chat messages(if multi msgs) to the current websocket message to reduce system calls
			n := len(receiveFromRoom)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write([]byte((<-receiveFromRoom).Payload))
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
