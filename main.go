package main

import (
	"log"
	"time"

	"703room/703room.com/api/middlewares"
	"703room/703room.com/api/routes"
	"703room/703room.com/config"
	"703room/703room.com/models"

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

	// Connect to the database
	db, err := config.ConnectToDatabase()

	err = db.AutoMigrate(
		&models.User{},
		&models.Room{},
		&models.RoomMember{},
		&models.Expense{},
		&models.Credits{},
		&models.UserHasPayment{},
		&models.RoomExpenseInvitationRecipient{},
		&models.RoomExpenseInvitationRequest{},
	)

	routes.SetUp(50*time.Second, db, r)

	// Run the server on port 8080 and check for errors
	r.Run()

}
