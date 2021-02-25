package main

import (
	"net/http"
	"unitalk/config"
	"unitalk/handler"
)

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWs(w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	http.ListenAndServe(config.Config["listen"].(string), nil)
}
