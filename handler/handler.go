package handler

import (
	"net/http"
	"unitalk/chat"
	"unitalk/logger"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// ServeWs handles websocket requests from clients requests.
func ServeWs(w http.ResponseWriter, r *http.Request) {
	room := r.URL.Query().Get("room")
	if room == "" {
		room = "default"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Writer.Error(err.Error(), zap.String("ws", "upgrade fail"))
		return
	}

	client := chat.NewClient(conn, room)

	go client.WritePump()
	go client.ReadPump()
}
