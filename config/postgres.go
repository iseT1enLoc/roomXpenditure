package config

import (
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")

	var db *gorm.DB
	var err error

	// Retry timeout setup
	deadline := time.Now().Add(20 * time.Second)

	for time.Now().Before(deadline) {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Database connection established.")
			return db, nil
		}

		log.Println("Failed to connect to DB. Retrying in 2 seconds...")
		time.Sleep(2 * time.Second)
	}

	return nil, errors.New("could not connect to database within 20 seconds")
}

//Logger: logger.Default.LogMode(logger.Info), // Enable logging mode
//DryRun: true
