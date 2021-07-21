package mq

import (
	"unitalk/config"

	"github.com/Shopify/sarama"
)

// KAFKAPRODUCER producer
var KAFKAPRODUCER sarama.SyncProducer

// TOPIC to send
var TOPIC string

func init() {
	kafkaConf := config.Config["kafka"].(map[string]interface{})
	TOPIC = kafkaConf["topic"].(string)
	var kafkaServers []string
	for _, v := range kafkaConf["serverList"].([]interface{}) {
		kafkaServers = append(kafkaServers, v.(string))
	}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // 0 1 all
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer(kafkaServers, config)
	if err != nil {
		panic(err)
	}
	KAFKAPRODUCER = producer
}
