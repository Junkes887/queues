package kafkaConsumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Run(topics []string, servers string, msgChan chan *kafka.Message) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": servers,
		"group.id":          "queues-go",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}
	consumer.SubscribeTopics(topics, nil)
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			msgChan <- msg
		}
	}
}
