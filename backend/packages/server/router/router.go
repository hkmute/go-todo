package router

import (
	"go-todo/packages/util/res"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	todoRoutes(r, "/todo")
	userRoutes(r, "/user")

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.NoRoute(func(c *gin.Context) {
		res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound, Message: "Not Found"})
	})
}
