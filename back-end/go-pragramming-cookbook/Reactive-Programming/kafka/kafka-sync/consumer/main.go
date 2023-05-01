package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = consumer.Close() }()

	partitionConsumer, err := consumer.ConsumePartition("example-topic", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer func() { _ = partitionConsumer.Close() }()

	for {
		msg := <-partitionConsumer.Messages()
		fmt.Printf("Consumed Message: '%s' at offset %d\n", msg.Value, msg.Offset)
	}
}
