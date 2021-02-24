package chat

// WsServer define
type WsServer struct {
	rooms map[*Room]bool
}

// NewWebsocketServer creates a new WsServer type
func NewWebsocketServer() *WsServer {
	return &WsServer{
		rooms: make(map[*Room]bool),
	}
}

// FindRoomByName method
func (server *WsServer) FindRoomByName(name string) *Room {
	for room := range server.rooms {
		if room.GetName() == name {
			return room
		}
	}
	return nil
}

// CreateRoom method
func (server *WsServer) CreateRoom(name string) *Room {
	room := NewRoom(name)
	go room.RunRoom()
	server.rooms[room] = true
	return room
}
