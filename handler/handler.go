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
func ServeWs(wsServer *chat.WsServer, w http.ResponseWriter, r *http.Request) {
	roomName := r.URL.Query().Get("room")
	if roomName == "" {
		roomName = "default"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Writer.Error(err.Error(), zap.String("ws", "upgrade fail"))
		return
	}

	var room *chat.Room
	room = wsServer.FindRoomByName(roomName)
	if room == nil {
		room = wsServer.CreateRoom(roomName)
	}

	client := chat.NewClient(conn, room)

	go client.WritePump()
	go client.ReadPump()

	room.TriggerRegisterClientInRoome(client)
}
