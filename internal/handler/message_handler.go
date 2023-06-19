package handler

import (
	"encoding/json"
	"fmt"

	"github.com/skiliyani/go-app/internal/businesslogic"
	"github.com/skiliyani/go-app/internal/models"
	"github.com/skiliyani/go-app/internal/storage"
)

// MessageHandler represents the handler for incoming messages
type MessageHandler struct {
	messageLogic *businesslogic.MessageLogic
	db           *storage.PostgresDB
	// Add any other dependencies or state here
}

// NewMessageHandler creates a new instance of MessageHandler
func NewMessageHandler(db *storage.PostgresDB) *MessageHandler {
	// Initialize any dependencies or state variables here
	messageLogic := businesslogic.NewMessageLogic()

	return &MessageHandler{
		messageLogic: messageLogic,
		db:           db,
	}
}

// HandleMessage handles the incoming message
func (mh *MessageHandler) HandleMessage(messageJSON string) error {
	// Parse the message JSON into a Message struct
	var message models.Message
	err := json.Unmarshal([]byte(messageJSON), &message)
	if err != nil {
		return fmt.Errorf("failed to parse the JSON message: %v", err)
	}

	// Process the message using the message logic
	err = mh.messageLogic.ProcessMessage(&message)
	if err != nil {
		return fmt.Errorf("failed to process the message: %v", err)
	}

	// Handle any other operations related to the message, such as storing it in the database
	err = mh.storeMessageInDatabase(&message)
	if err != nil {
		return fmt.Errorf("failed to store the message in the database: %v", err)
	}

	return nil
}

// storeMessageInDatabase stores the message in the database
func (mh *MessageHandler) storeMessageInDatabase(message *models.Message) error {
	// Implement your code to store the message in the database using the provided db connection
	// Example:
	// err := mh.db.Save(message)
	// if err != nil {
	//   return fmt.Errorf("failed to store the message in the database: %v", err)
	// }

	return nil
}
