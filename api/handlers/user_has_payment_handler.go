package handlers

import (
	"703room/703room.com/models"
	"703room/703room.com/services"
	"703room/703room.com/utils"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHasPaymentHandler struct {
	user_has_payment services.UserHasPaymentService
	user_service     services.UserService
}

func NewUserHasPaymentHandler(user_has_payment services.UserHasPaymentService, user_service services.UserService) *UserHasPaymentHandler {
	return &UserHasPaymentHandler{
		user_has_payment: user_has_payment,
		user_service:     user_service,
	}
}

func (e *UserHasPaymentHandler) CreateNewExpense() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var expenseBody struct {
			RoomID   uuid.UUID  `json:"room_id"` // Required to associate expense to a room
			Title    string     `json:"title"`
			Amount   float64    `json:"amount"`
			Quantity int        `json:"quantity"`
			UsedDate *time.Time `json:"used_date"`
			Notes    string     `json:"notes"`
		}

		// Bind JSON input
		if err := ctx.ShouldBindJSON(&expenseBody); err != nil {
			utils.Error(ctx, 400, "Cannot parse expense body", err)
			return
		}

		// Get user_id from context
		id, ok := ctx.Get("user_id")
		if !ok {
			utils.Error(ctx, 400, "No user ID found in context", nil)
			return
		}

		userID, ok := id.(uuid.UUID)
		if !ok {
			utils.Error(ctx, 500, "Invalid user ID format", nil)
			return
		}
		usedDate := expenseBody.UsedDate
		if usedDate == nil {
			now := time.Now()
			usedDate = &now
		}
		log.Println(expenseBody)

		// Create UserHasPayment object
		userHasPayment := &models.UserHasPayment{
			ID:        uuid.New(),
			RoomID:    expenseBody.RoomID,
			UserID:    userID,
			Title:     expenseBody.Title,
			Quantity:  expenseBody.Quantity,
			Amount:    expenseBody.Amount,
			Notes:     expenseBody.Notes,
			UsedDate:  expenseBody.UsedDate,
			CreatedAt: time.Now(),
		}

		// Call service to create expense
		if err := e.user_has_payment.CreateExpense(ctx, userHasPayment); err != nil {
			utils.Error(ctx, 400, "Failed to create expense", err)
			return
		}

		utils.Created(ctx, "Created expense successfully", userHasPayment)
	}
}

// http://localhost:8080/api/protected/expense/user?year=2025&month=4&day=22
func (h *UserHasPaymentHandler) GetExpensesFiltered() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get user_id from context
		id, exists := ctx.Get("user_id")
		if !exists {
			utils.Error(ctx, 401, "User ID not found in context", nil)
			return
		}

		_, ok := id.(uuid.UUID)
		if !ok {
			utils.Error(ctx, 500, "User ID type assertion failed", nil)
			return
		}

		// Get room_id from query
		roomIDStr := ctx.Query("room_id")
		if roomIDStr == "" {
			utils.Error(ctx, 400, "Missing required parameter: room_id", nil)
			return
		}

		roomID, err := uuid.Parse(roomIDStr)
		if err != nil {
			utils.Error(ctx, 400, "Invalid room_id format", err)
			return
		}

		// Optional filters
		year := ctx.Query("year")
		month := ctx.Query("month")
		day := ctx.Query("day")
		user_id := ctx.Query("user_id")
		uID, err := uuid.Parse(user_id)
		if err != nil {
			utils.Error(ctx, 400, "Invalid user_id format", err)
			return
		}
		// Call service
		expenses, err := h.user_has_payment.GetExpensesFiltered(ctx, uID, roomID, year, month, day)
		if err != nil {
			utils.Error(ctx, 400, "Failed to fetch expenses", err)
			return
		}

		utils.Success(ctx, "Fetched expenses successfully", expenses)
	}
}

func (h *UserHasPaymentHandler) CalculateMonthExpense() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Enter calculate expense")
		// Get required room_id
		roomIDStr := ctx.Query("room_id")
		if roomIDStr == "" {
			utils.Error(ctx, 400, "Missing required parameter: room_id", nil)
			return
		}

		roomID, err := uuid.Parse(roomIDStr)
		if err != nil {
			utils.Error(ctx, 400, "Invalid room_id format", err)
			return
		}
		log.Println("Enter line 161")
		// Optional filters
		year := ctx.Query("year")
		month := ctx.Query("month")
		day := ctx.Query("day")

		// Get all users in the room (adjust GetAllUsers to accept roomID if needed)
		users, err := h.user_service.GetUsersByRoomID(ctx, roomID)
		fmt.Println(users)
		if err != nil {
			utils.Error(ctx, 400, "Failed to fetch users", err)
			return
		}
		fmt.Println("Enter line 174")
		// Prepare response struct
		var response_data struct {
			RoomTotalExpense float64       `json:"room_total_expense"`
			MemberStat       []member_stat `json:"member_stat"`
		}
		var totalRoomAmount float64
		fmt.Println("Enter line 181")
		for _, user := range users {
			fmt.Println("Enter line 183")
			value, _ := h.user_has_payment.CalculateMemberExpenseByMemberId(ctx, user.UserID, roomID, year, month, day)

			memStat := member_stat{
				Member_name: user.Name,
				Money:       value,
				Member_Id:   user.UserID,
			}
			response_data.MemberStat = append(response_data.MemberStat, memStat)
			log.Println(memStat)
			totalRoomAmount += value
		}

		response_data.RoomTotalExpense = totalRoomAmount

		utils.Success(ctx, "Fetched expenses successfully", response_data)
	}
}
func (h *UserHasPaymentHandler) GetAllRoomMemberExpenseFilter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get user_id from context
		id, exists := ctx.Get("user_id")
		if !exists {
			utils.Error(ctx, 401, "User ID not found in context", nil)
			return
		}

		_, ok := id.(uuid.UUID)
		if !ok {
			utils.Error(ctx, 500, "User ID type assertion failed", nil)
			return
		}

		// Get room_id from query
		roomIDStr := ctx.Query("room_id")
		if roomIDStr == "" {
			utils.Error(ctx, 400, "Missing required parameter: room_id", nil)
			return
		}

		roomID, err := uuid.Parse(roomIDStr)
		if err != nil {
			utils.Error(ctx, 400, "Invalid room_id format", err)
			return
		}

		// Optional filters
		year := ctx.Query("year")
		month := ctx.Query("month")
		day := ctx.Query("day")

		// Call service
		expenses, err := h.user_has_payment.GetRoomExpenseDetails(ctx, roomID, year, month, day)
		if err != nil {
			utils.Error(ctx, 400, "Failed to fetch expenses", err)
			return
		}

		utils.Success(ctx, "Fetched expenses successfully", expenses)
	}
}
