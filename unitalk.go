package main

import (
	"net/http"
	"unitalk/config"
	"unitalk/handler"
	"unitalk/logger"
)

func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWs(w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	logger.Writer.Error(http.ListenAndServe(config.Config["listen"].(string), nil).Error())
}
