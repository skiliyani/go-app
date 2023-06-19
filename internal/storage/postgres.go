package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// PostgresDB represents a PostgreSQL database connection
type PostgresDB struct {
	db *sql.DB
}

// NewPostgresDB creates a new instance of PostgresDB and establishes a connection to the database
func NewPostgresDB(connectionString string) (*PostgresDB, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	return &PostgresDB{
		db: db,
	}, nil
}

// StoreMessage stores the given message into the database
func (p *PostgresDB) StoreMessage(message string) error {
	// Implement your logic to store the message in the database
	// You can use the db.Exec or db.Query functions to execute SQL statements

	// Example: Insert the message into a "messages" table
	query := "INSERT INTO messages (content) VALUES ($1)"
	_, err := p.db.Exec(query, message)
	if err != nil {
		return fmt.Errorf("failed to store the message: %v", err)
	}

	return nil
}

// Close closes the connection to the database
func (p *PostgresDB) Close() error {
	return p.db.Close()
}
