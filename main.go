package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

func main() {

	// to consume messages

	kafkaHost := os.Getenv("kafka.bootstrap")
	if kafkaHost == "" {
		kafkaHost = "localhost:9092"
	}

	topic := os.Getenv("kafka.topic")
	if topic == "" {
		topic = "forestTopic"
	}

	partition := 0

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafkaHost, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	//batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	//b := make([]byte, 10e3) // 10KB max per message
	for {

		message, err := conn.ReadMessage(1e3)
		if err != nil {
			log.Println("Error while reading: ", err)
			break
		} else {
			fmt.Println("Received:", string(message.Value))

		}
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}

}
