package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"net/http"
)

const PORT = ":7166"

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer func() { _ = producer.Close() }()

	// 카프카 메시지 발행 프로스세스를 고루틴으로 실행
	go ProcessResponse(producer)

	// 카프카 메시지 발행을 위한 API Handler 등록 및 서버 실행
	controller := KafkaController{producer: producer}
	http.HandleFunc("/", controller.KafkaProduceHandler)

	fmt.Printf("Listening on port %s\n", PORT)
	fmt.Println(http.ListenAndServe(PORT, nil))
}
