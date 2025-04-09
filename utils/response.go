package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  string      `json:"status"`            // "success" or "error"
	Message string      `json:"message,omitempty"` // Optional human-readable message
	Data    interface{} `json:"data,omitempty"`    // Payload
	Error   interface{} `json:"error,omitempty"`   // Error details if any
}

func Success(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, code int, message string, err interface{}) {
	c.JSON(code, APIResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	})
}
