package kafkaRestHandler

import (
	"net/http"

	"KafkaConsumer/kafka"
)

type KafkaRestHandler struct {
	consumer *kafka.Consumer
}

func NewKafkaRestHandler(consumer *kafka.Consumer) *KafkaRestHandler {
	return &KafkaRestHandler{
		consumer: consumer,
	}
}

func (KafkaRestHandler *KafkaRestHandler) ConsumeHandler(writer http.ResponseWriter, request *http.Request) {
	message, err := KafkaRestHandler.consumer.ConsumeOne()
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte("asd"))
		return
	}

	writer.WriteHeader(200)
	_, _ = writer.Write(message.Value)
}
