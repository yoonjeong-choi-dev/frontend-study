package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func produceMessage(producer sarama.SyncProducer, value string) {
	msg := &sarama.ProducerMessage{
		Topic: "example-topic",
		Value: sarama.StringEncoder(value),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message: %s\n", err)
		return
	}

	log.Printf("> message sent to partioin %d at offset %d\n", partition, offset)
}

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = producer.Close() }()

	for i := 0; i < 10; i++ {
		produceMessage(producer, fmt.Sprintf("Sample Message - %d", i))
	}

}
