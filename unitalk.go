package main

import (
	"net/http"
	"unitalk/chat"
	"unitalk/config"
)

func main() {
	wsServer := chat.NewWebsocketServer()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.ListenAndServe(config.Config["listen"].(string), nil)
}
