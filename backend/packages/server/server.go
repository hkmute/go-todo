package server

import (
	"go-todo/packages/middleware"
	"go-todo/packages/server/router"
	"go-todo/packages/util/res"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.New()

	// r.Use(cors.Default())
	// r.Use(middleware.CorsMiddleware())

	r.Use(
		middleware.CorsMiddleware(),
		gin.Logger(),
		gin.CustomRecovery(func(c *gin.Context, err interface{}) {
			if err, ok := err.(error); ok {
				res.JsonError(c, res.ErrorParams{Message: err.Error()})
			} else {
				res.JsonError(c, res.ErrorParams{Message: "An unexpected error occurred"})
			}
			c.Abort()
			return
		}))

	router.Init(r)

	r.Run()
}
