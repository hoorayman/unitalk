package main

import (
	"fmt"
	"net/http"
	"unitalk/broker"
	"unitalk/config"
	"unitalk/handler"
	"unitalk/logger"
	"unitalk/mq"
	"unitalk/reg"
)

func main() {
	defer broker.REDIS.Close()
	defer reg.ZK.Close()
	defer mq.KAFKAPRODUCER.Close()
	fmt.Println("start service on " + config.Config["listen"].(string))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handler.ServeWs(w, r)
	})

	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	logger.Writer.Error(http.ListenAndServe(config.Config["listen"].(string), nil).Error())
}
