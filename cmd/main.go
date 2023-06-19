package main

import (
	"fmt"
	"log"

	"github.com/skiliyani/go-app/internal/config"
	"github.com/skiliyani/go-app/internal/handler"
	"github.com/skiliyani/go-app/internal/messaging"
	"github.com/skiliyani/go-app/internal/storage"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize messaging component
	sqsClient := messaging.NewSQSClient(cfg.AWSRegion)
	queue := messaging.NewQueue(cfg.SQSQueueURL, sqsClient)

	// Initialize storage component
	db, err := storage.NewPostgresDB(cfg.DBConnectionString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize message handler
	msgHandler := handler.NewMessageHandler(db)

	// Start listening for messages
	err = queue.StartListening(msgHandler.HandleMessage)
	if err != nil {
		log.Fatalf("Failed to start listening for messages: %v", err)
	}

	fmt.Println("Listening for incoming messages...")
	select {}
}
