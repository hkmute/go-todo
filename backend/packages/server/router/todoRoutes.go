package router

import (
	"go-todo/packages/db"
	"go-todo/packages/middleware"
	"go-todo/packages/todo"

	"github.com/gin-gonic/gin"
)

type todoController interface {
	GetTodoList(c *gin.Context)
	GetTodoById(c *gin.Context)
	PostTodo(c *gin.Context)
	PutTodoById(c *gin.Context)
	ReorderTodoById(c *gin.Context)
	DeleteTodoById(c *gin.Context)
}

func todoRoutes(r *gin.Engine, path string, handlers ...gin.HandlerFunc) {
	var controller todoController = todo.NewTodoController(todo.NewTodoService(db.DB.Conn()))
	authMiddleware := middleware.AuthMiddleware()

	routes := r.Group(path, append(handlers, authMiddleware)...)

	routes.GET("/", controller.GetTodoList)
	routes.GET("/:id", controller.GetTodoById)
	routes.POST("/", controller.PostTodo)
	routes.PUT("/:id", controller.PutTodoById)
	routes.PUT("/:id/reorder", controller.ReorderTodoById)
	routes.DELETE("/:id", controller.DeleteTodoById)
}
