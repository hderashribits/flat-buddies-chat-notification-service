package main

import (
	"log"

	"notification-service/internal/consumer"
)

func main() {
	log.Println("🚀 Starting Notification Service...")
	consumer.Start()
}
