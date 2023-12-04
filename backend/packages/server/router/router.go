package router

import (
	"go-todo/packages/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	authMiddleware := middleware.AuthMiddleware()

	todoRoutes(r, "/todo", authMiddleware)
	userRoutes(r, "/user")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
}
