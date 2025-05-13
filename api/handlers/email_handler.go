package handlers

import (
	"703room/703room.com/services"
	"703room/703room.com/utils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Emailhandler struct {
	Email_service services.EmailService
}

func NewEmailhandler(Email_service services.EmailService) *Emailhandler {
	return &Emailhandler{
		Email_service: Email_service,
	}
}

type SendReportRequest struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Message   string `json:"message"`
}

func (e *Emailhandler) SendEmailHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		type RoomID struct {
			RoomId  uuid.UUID `json:"room_id"`
			Year    string    `json:"year"`
			Month   string    `json:"month"`
			Day     string    `json:"day"`
			Message string    `json:"message"`
		}
		log.Println("Enter line 36")
		var room_id RoomID
		if err := ctx.ShouldBindBodyWithJSON(&room_id); err != nil {
			log.Println("Enter line", err)
			utils.Error(ctx, 401, "Can not parsing json data", err)
			return
		}
		log.Println("Enter line 41")
		err := e.Email_service.SendReportToRoomate(ctx, room_id.RoomId, room_id.Year, room_id.Month, room_id.Day, room_id.Message)
		if err != nil {
			log.Println("Enter line", err)
			utils.Error(ctx, 401, "Can not send email sucessfully", err)
			return
		}
		log.Println("Enter line 46")
		utils.Success(ctx, "call email successfully", "Email sent successfully to all roommates")
	}
}
