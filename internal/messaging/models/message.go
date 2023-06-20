package models

type Message struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	// Add other fields specific to the "message" JSON structure
}
