package main

import (
	"github.com/Shopify/sarama"
	"log"
)

// ProcessResponse 비동기 sarama.AsyncProducer 객체로 부터 결과와 에러를 비동기로 받아옴
// => 비동기 처리를 위해 채널 및 select 사용
func ProcessResponse(producer sarama.AsyncProducer) {
	for {
		select {
		case result := <-producer.Successes():
			log.Printf("> message: '%s' sent to partition %d at offset %d\n",
				result.Value,
				result.Partition,
				result.Offset,
			)
		case err := <-producer.Errors():
			log.Println("Failed to produce message:", err)
		}
	}
}
