# Intermediary container for compilation
FROM golang:1.13-alpine
RUN apk update
WORKDIR /go/src/KafkaConsumer/
COPY ./src /go/src/KafkaConsumer/
RUN go build -o kafkaConsumer.bin

# Image with binary
FROM alpine:3.11
WORKDIR /
COPY --from=0 /go/src/KafkaConsumer/kafkaConsumer.bin .
CMD ["./kafkaConsumer.bin"]
