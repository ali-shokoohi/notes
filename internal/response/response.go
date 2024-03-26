// internal/response/response.go
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApiResponse defines the structure of your standard API response
type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// ApiError defines the structure of your standard API error response
type ApiError struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// SendData sends a standard API response with data
func SendData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, ApiResponse{Success: true, Data: data})
}

// SendMessage sends a standard API response with a message
func SendMessage(c *gin.Context, message string) {
	c.JSON(http.StatusOK, ApiResponse{Success: true, Message: message})
}

// SendError sends a standard API error response
func SendError(c *gin.Context, statusCode int, errMessage string) {
	c.JSON(statusCode, ApiError{Success: false, Error: errMessage})
}
