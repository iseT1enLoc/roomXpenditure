package main

import (
	"log"
	"time"

	"703room/703room.com/api/middlewares"
	"703room/703room.com/api/routes"
	"703room/703room.com/config"
	"703room/703room.com/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Gin engine with default middleware (logger and recovery)
	r := gin.Default()

	// Apply CORS middleware
	r.Use(middlewares.CORSMiddleware())
	// Uncomment if you need JWT middleware
	// r.Use(middlewares.JWTMiddleware("hjgrtjtbun"))

	// Connect to the database
	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatalf("[ERROR] Connecting to database failed: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.RoomMember{},
		&models.Expense{},
		&models.Credits{}, // if defined in your models
	)
	if err != nil {
		log.Fatalf("[ERROR]: %v", err)
	}
	// Setup application routes with a timeout of 50 seconds
	routes.SetUp(50*time.Second, db, r)

	// Run the server on port 8080 and check for errors
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("[ERROR] Server failed to start: %v", err)
	}
}
