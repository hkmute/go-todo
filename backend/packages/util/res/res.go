package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type successResponse struct {
	Success bool                   `json:"success"`
	Data    interface{}            `json:"data,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ErrorParams struct {
	StatusCode int
	Message    string
}

type Meta struct {
	Data map[string]interface{}
}

func JsonSuccess(c *gin.Context, data interface{}, meta ...Meta) {
	var metaData map[string]interface{}
	if len(meta) > 0 {
		metaData = meta[0].Data
	}
	c.JSON(http.StatusOK, successResponse{
		Success: true,
		Data:    data,
		Meta:    metaData,
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

	c.JSON(code, errorResponse{
		Success: false,
		Message: message,
	})
}
