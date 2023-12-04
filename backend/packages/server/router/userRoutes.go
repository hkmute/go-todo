package router

import (
	"go-todo/packages/db"
	"go-todo/packages/middleware"
	"go-todo/packages/user"

	"github.com/gin-gonic/gin"
)

type userController interface {
	GetMe(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
}

func userRoutes(r *gin.Engine, path string, handlers ...gin.HandlerFunc) {
	var controller userController = user.NewUserController(user.NewUserService(db.DB.Conn()))
	authMiddleware := middleware.AuthMiddleware()

	routes := r.Group(path, handlers...)

	routes.GET("/me", authMiddleware, controller.GetMe)
	routes.POST("/register", controller.Register)
	routes.POST("/login", controller.Login)
}
