package main

import (
	"log"
	"net/http"
)

func main() {
	wsServer := NewWebsocketServer()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
