package main

import (
	"log"
	"net/http"

	"chat-service/internal/handler"
	"chat-service/internal/producer"
)

func main() {
	kafkaBrokers := []string{"kafka:9092"}
	kafkaTopic := "chat-messages"
	producer.InitKafka(kafkaBrokers, kafkaTopic)

	// Register the /send endpoint
	http.HandleFunc("/send", handler.SendMessageHandler)

	// Register the /match endpoint
	http.HandleFunc("/notification", handler.HandleNotification)

	log.Println("🚀 Service running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
