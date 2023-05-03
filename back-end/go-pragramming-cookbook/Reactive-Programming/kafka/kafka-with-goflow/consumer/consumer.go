package main

import (
	"github.com/Shopify/sarama"
	"github.com/trustmaster/goflow"
	kafka_with_goflow "kafka-with-goflow"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer func() { _ = consumer.Close() }()

	partitionConsumer, err := consumer.ConsumePartition("goflow-example", 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	defer func() { _ = partitionConsumer.Close() }()

	network := kafka_with_goflow.NewUpperApp()

	in := make(chan string)
	network.SetInPort("In", in)

	wait := goflow.Run(network)
	defer func() {
		close(in)
		<-wait
	}()

	for {
		// 카프카 메시지 큐에서 메시지를 받아와서 goflow network 진입점에 전달
		msg := <-partitionConsumer.Messages()
		in <- string(msg.Value)
	}
}
