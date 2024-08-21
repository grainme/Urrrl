package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// URL represents the model for a shortened URL
type URL struct {
	ID          string    `gorm:"primaryKey"`
	ShortURL    string    `gorm:"uniqueIndex;size:10"`
	OriginalURL string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

// InitDB initializes the database connection and migrates the schema
func InitDB() (*gorm.DB, error) {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("urlshortener.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&URL{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// Initialize the database
	db, err := InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	fmt.Println("Database connection and migration successful!")

	// Example usage: Create a new URL record
	newURL := URL{
		ID:          "1",
		ShortURL:    "abc123",
		OriginalURL: "https://example.com",
	}

	if err := db.Create(&newURL).Error; err != nil {
		log.Fatalf("Failed to create new URL record: %v", err)
	}

	fmt.Println("New URL record created successfully!")
}
