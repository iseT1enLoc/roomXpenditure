package main

import (
	"log"
	"time"

	"703room/703room.com/api/middlewares"
	"703room/703room.com/api/routes"
	"703room/703room.com/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Initialize Gin engine with default middleware (logger and recovery)
	r := gin.Default()
	err := godotenv.Load()
	if err != nil {
		log.Println("Không thể load file .env (có thể bạn đang chạy production).")
	}
	// Apply CORS middleware
	r.Use(middlewares.CORSMiddleware())

	// Uncomment if you need JWT middleware

	// Connect to the database
	db, err := config.ConnectToDatabase()
	if err != nil {
		log.Fatalf("[ERROR] Connecting to database failed: %v", err)
	}

	// err = db.AutoMigrate(
	// 	&models.User{},
	// 	&models.Room{},
	// 	&models.RoomMember{},
	// 	&models.Expense{},
	// 	&models.Credits{}, // if defined in your models
	// 	&models.UserHasPayment{},
	// 	&models.RoomExpenseInvitationRecipient{},
	// 	&models.RoomExpenseInvitationRequest{},
	// )
	// if err != nil {
	// 	log.Fatalf("[ERROR]: %v", err)
	// }
	// Setup application routes with a timeout of 50 seconds
	routes.SetUp(50*time.Second, db, r)

	// Run the server on port 8080 and check for errors
	r.Run()

}
