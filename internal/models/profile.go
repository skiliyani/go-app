package models

type Profile struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
	// Add other fields specific to the "profile" JSON structure
}
