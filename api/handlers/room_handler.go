package handlers

import (
	"703room/703room.com/models"
	"703room/703room.com/services"
	"703room/703room.com/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoomHandler struct {
	room_service services.RoomService
}

func NewRoomHandler(room_service services.RoomService) *RoomHandler {
	return &RoomHandler{
		room_service: room_service,
	}
}

func (r *RoomHandler) CreateNewRoom() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get user_id from context and return early on failure
		id, ok := ctx.Get("user_id")
		log.Println(id)
		if !ok {
			utils.Error(ctx, 400, "No user ID found in context", nil)
			return
		}

		userID, ok := id.(uuid.UUID)
		if !ok {
			utils.Error(ctx, 500, "Invalid user ID format", nil)
			return
		}
		log.Println(userID)
		var room_name_req struct {
			Room_name string `json:"room_name"`
		}
		if err := ctx.ShouldBindBodyWithJSON(&room_name_req); err != nil {
			utils.Error(ctx, 404, "Can not parsing", err)
			return
		}
		room := &models.Room{
			RoomID:    uuid.New(),
			RoomName:  room_name_req.Room_name,
			CreatedBy: userID,
			CreatedAt: time.Now(),
		}
		err := r.room_service.CreateRoom(ctx, room)
		if err != nil {
			utils.Error(ctx, 500, "Error happened while create new room", err)
			return
		}
		//first member
		headMember := models.RoomMember{
			ID:        uuid.New(),
			RoomID:    room.RoomID,
			UserID:    userID,
			CreatedAt: time.Now(),
			Role:      "truongphong",
			JoinedAt:  time.Now(),
		}
		err = r.room_service.AddMember(ctx, &headMember)
		if err != nil {
			utils.Error(ctx, 500, "Error while add new member to the room", err)
			return
		}
		utils.Created(ctx, "Successfully create new room", room)
	}
}

func (r *RoomHandler) GetAllRoomsOfUserByUserID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Attempt to get user_id from context
		rawID, exists := ctx.Get("user_id")
		log.Println("[Handler] Context user_id:", rawID)

		if !exists {
			utils.Error(ctx, http.StatusBadRequest, "No user ID found in context", nil)
			return
		}

		userID, ok := rawID.(uuid.UUID)
		if !ok {
			utils.Error(ctx, http.StatusInternalServerError, "Invalid user ID format", nil)
			return
		}

		// Fetch rooms from service
		rooms, err := r.room_service.ListRoomsByUserID(ctx, userID.String())
		if err != nil {
			log.Println("[Handler] Failed to get rooms:", err)
			utils.Error(ctx, http.StatusInternalServerError, "Failed to get rooms", nil)
			return
		}

		log.Println("[Handler] Successfully retrieved rooms for user:", userID)
		utils.Success(ctx, "Get rooms by user ID successfully", rooms)
	}
}
