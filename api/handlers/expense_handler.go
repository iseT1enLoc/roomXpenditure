package handlers

import (
	"703room/703room.com/models"
	"703room/703room.com/services"
	"703room/703room.com/utils"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ExpenseHandler struct {
	expense_service services.ExpenseService
	user_service    services.UserService
}

func NewExpenseHandler(expense_service services.ExpenseService, user_service services.UserService) *ExpenseHandler {
	return &ExpenseHandler{
		expense_service: expense_service,
		user_service:    user_service,
	}
}

func (e *ExpenseHandler) CreateNewExpense() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var expenseBody struct {
			Title  string  `json:"title"` // Capitalized fields!
			Amount float64 `json:"amount"`
			Notes  string  `json:"notes"`
		}

		// Bind and return early on error
		if err := ctx.ShouldBindJSON(&expenseBody); err != nil {
			utils.Error(ctx, 400, "Cannot parse expense body", err)
			return
		}

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
		// Create the Expense object
		expense := models.Expense{
			ExpenseID: uuid.New(),
			UserID:    userID,
			Title:     expenseBody.Title,
			Amount:    expenseBody.Amount,
			Notes:     expenseBody.Notes,
			CreatedAt: time.Now(),
		}

		// Attempt to insert into database
		if err := e.expense_service.CreateExpense(ctx, &expense); err != nil {
			utils.Error(ctx, 400, err.Error(), err)
			return
		}

		// Respond with created success
		utils.Created(ctx, "Created expense successfully", expense)
	}
}
func (e *ExpenseHandler) GetExpenseByID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idStr := ctx.Param("id")
		if idStr == "" {
			utils.Error(ctx, 400, "Missing expense ID parameter", nil)
			return
		}

		expenseID, err := uuid.Parse(idStr)
		if err != nil {
			utils.Error(ctx, 400, "Invalid UUID format", err)
			return
		}

		expense, err := e.expense_service.GetExpenseByID(ctx, expenseID)
		if err != nil {
			utils.Error(ctx, 404, "Expense not found", err)
			return
		}

		utils.Success(ctx, "Successfully retrieved expense by ID", expense)
	}
}

type member_stat struct {
	Member_name string
	Money       float64
}

// http://localhost:8080/api/protected/expense/user?year=2025&month=4&day=22
func (h *ExpenseHandler) GetExpensesFiltered() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, exists := ctx.Get("user_id")
		if !exists {
			utils.Error(ctx, 401, "User ID not found in context", nil)
			return
		}

		userID, ok := id.(uuid.UUID)
		if !ok {
			utils.Error(ctx, 500, "User ID type assertion failed", nil)
			return
		}

		// Optional filter by mode, month, year, day
		year := ctx.Query("year")
		month := ctx.Query("month")
		day := ctx.Query("day")

		expenses, err := h.expense_service.GetExpensesFiltered(ctx, userID, year, month, day)
		if err != nil {
			utils.Error(ctx, 400, "Failed to fetch expenses", err)
			return
		}

		utils.Success(ctx, "Fetched expenses successfully", expenses)
	}
}
func (h *ExpenseHandler) CalculateMonthExpense() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Optional filter by mode, month, year, day
		year := ctx.Query("year")
		month := ctx.Query("month")
		day := ctx.Query("day")

		users, err := h.user_service.GetAllUsers(ctx)
		log.Println(users)
		if err != nil {
			utils.Error(ctx, 400, "Failed to fetch expenses", err)
			return
		}
		var response_data struct {
			Rooom_total_expense float64
			Member_Stat         []member_stat
		}
		var total_room_amount float64
		for i := 0; i < len(users); i += 1 {
			log.Println(users[i].Name)
			value, _ := h.expense_service.CalculateMemberExpenseByMemberId(ctx, users[i].UserID, year, month, day)
			var mem_stat member_stat
			mem_stat.Member_name = users[i].Name
			mem_stat.Money = value
			total_room_amount += value
			response_data.Member_Stat = append(response_data.Member_Stat, mem_stat)
		}
		response_data.Rooom_total_expense = total_room_amount
		log.Println(response_data.Rooom_total_expense)
		log.Println(response_data.Member_Stat)
		utils.Success(ctx, "Fetched expenses successfully", response_data)
	}
}
