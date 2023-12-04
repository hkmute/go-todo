package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	todoRoutes(r, "/todo")
	userRoutes(r, "/user")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})
}
