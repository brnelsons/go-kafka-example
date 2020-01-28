package kafkaRestHandler

import (
	"KafkaConsumer/kafka"
	"io/ioutil"
	"net/http"
)

type KafkaRestHandler struct {
	consumer *kafka.Consumer
	producer *kafka.Producer
}

func NewKafkaRestHandler(consumer *kafka.Consumer, producer *kafka.Producer) *KafkaRestHandler {
	return &KafkaRestHandler{
		consumer: consumer,
		producer: producer,
	}
}

func (KafkaRestHandler *KafkaRestHandler) ConsumeHandler(writer http.ResponseWriter, request *http.Request) {
	message, err := KafkaRestHandler.consumer.ConsumeOne()
	if KafkaRestHandler.handleError(err, writer) {
		return
	}

	writer.WriteHeader(200)
	_, err = writer.Write(message.Value)
	if KafkaRestHandler.handleError(err, writer) {
		return
	}
}

func (KafkaRestHandler *KafkaRestHandler) ProduceHandler(writer http.ResponseWriter, request *http.Request) {
	bytes, err := ioutil.ReadAll(request.Body)
	if KafkaRestHandler.handleError(err, writer) {
		return
	}
	err = KafkaRestHandler.producer.Produce(bytes)
	if KafkaRestHandler.handleError(err, writer) {
		return
	}
}

func (KafkaRestHandler *KafkaRestHandler) handleError(err error, writer http.ResponseWriter) bool {
	if err != nil {
		writer.WriteHeader(500)
		_, _ = writer.Write([]byte(err.Error()))
		return true
	}
	return false
}
