package main

import (
	"github.com/Shopify/sarama"
	"net/http"
)

type KafkaController struct {
	producer sarama.AsyncProducer
}

// KafkaProduceHandler 카프카에 메시지를 발행하는 핸들러
// => http 요청을 통해 카프카에 메시지 발생
func (c *KafkaController) KafkaProduceHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	msg := r.FormValue("msg")
	if msg == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("msg must be set"))
		return
	}

	// 라이브러리에서 제공하는 input 채널을 이용하여 비동기적으로 처리
	c.producer.Input() <- &sarama.ProducerMessage{
		Topic: "async-example",
		Key:   nil,
		Value: sarama.StringEncoder(msg),
	}

	// 카프카에 메시지 발행 성공 시, 200 응답
	// => 발행된 메시지에 대한 소비 여부는 상관X
	w.WriteHeader(http.StatusOK)
}
