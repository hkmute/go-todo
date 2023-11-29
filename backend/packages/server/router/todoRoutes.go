package router

import (
	"go-todo/packages/db"
	"go-todo/packages/todo"

	"github.com/gin-gonic/gin"
)

type todoController interface {
	GetTodoList(c *gin.Context)
	GetTodoListWithGoRoutine(c *gin.Context)
	GetTodoById(c *gin.Context)
	PostTodo(c *gin.Context)
	PutTodoById(c *gin.Context)
	DeleteTodoById(c *gin.Context)
}

func todoRoutes(r *gin.Engine, path string) {
	var controller todoController = todo.NewTodoController(todo.NewTodoService(db.DB.Conn()))

	routes := r.Group(path)

	routes.GET("/", controller.GetTodoList)
	routes.GET("/go", controller.GetTodoListWithGoRoutine)
	routes.GET("/:id", controller.GetTodoById)
	routes.POST("/", controller.PostTodo)
	routes.PUT("/:id", controller.PutTodoById)
	routes.DELETE("/:id", controller.DeleteTodoById)
}
