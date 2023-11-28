package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrorParams struct {
	StatusCode int
	Message    string
}

func JsonSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func JsonError(c *gin.Context, err ErrorParams) {
	code := err.StatusCode
	message := err.Message
	if code == 0 {
		code = http.StatusInternalServerError
	}
	if message == "" {
		message = http.StatusText(code)
	}

	c.JSON(code, ErrorResponse{
		Success: false,
		Message: message,
	})
}
