package models

type Profile struct {
	ID  int    `db:"id"`
	URL string `db:"url"`
	// Add other fields specific to the "profile" db structure
}
