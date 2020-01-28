package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"KafkaConsumer/kafka"
	"KafkaConsumer/kafkaRestHandler"
	"github.com/gorilla/mux"
)

func main() {
	kafkaBrokerUrls := os.Getenv("KAFKA_BROKER_URLS")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	kafkaGroupId := os.Getenv("KAFKA_GROUP_ID")

	consumer := kafka.NewConsumer(
		strings.Split(kafkaBrokerUrls, ","),
		kafkaTopic,
		kafkaGroupId,
	)

	handler := kafkaRestHandler.NewKafkaRestHandler(consumer)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/consume", handler.ConsumeHandler)
	srv := &http.Server{
		Handler: router,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
