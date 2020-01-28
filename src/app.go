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
	kafkaBrokerUrls := strings.Split(os.Getenv("KAFKA_BROKER_URLS"), ",")
	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	kafkaGroupId := os.Getenv("KAFKA_GROUP_ID")
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	// setup consumer
	consumer := kafka.NewConsumer(
		kafkaBrokerUrls,
		kafkaTopic,
		kafkaGroupId,
	)

	// setup producer
	producer := kafka.NewProducer(
		kafkaBrokerUrls,
		kafkaTopic,
	)

	handler := kafkaRestHandler.NewKafkaRestHandler(consumer, producer)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/consume", handler.ConsumeHandler)
	router.HandleFunc("/api/v1/produce", handler.ProduceHandler)
	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
