package todo

import (
	"go-todo/packages/util/appError"
	"go-todo/packages/util/res"
	_type "go-todo/packages/util/type"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type todoService interface {
	GetTodoList() []todoEntity
	GetTodoById(id int) (todoEntity, error)
	InsertTodo(newTodo NewTodo) (todoEntity, error)
	EditTodoById(id int, editTodo EditTodo) (todoEntity, error)
	DeleteTodoById(id int) bool
}

type todoController struct {
	todoService todoService
}

func NewTodoController(todoService todoService) todoController {
	return todoController{todoService: todoService}
}

func (controller todoController) GetTodoList(c *gin.Context) {
	var params GetTodoListParams
	if err := c.ShouldBind(&params); err != nil {
		errorMessage := appError.Message(err)
		res.JsonError(c, res.ErrorParams{Message: errorMessage})
		return
	}

	todoList := controller.todoService.GetTodoList()
	res.JsonSuccess(c, todoList)
}

func (controller todoController) GetTodoById(c *gin.Context) {
	var uri _type.Id

	if err := c.ShouldBindUri(&uri); err != nil {
		res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound})
		return
	}

	todo, err := controller.todoService.GetTodoById(uri.Id)
	if err != nil {
		if err == pgx.ErrNoRows {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound})
		} else {
			res.JsonError(c, res.ErrorParams{Message: err.Error()})
		}
		return
	}
	res.JsonSuccess(c, todo)
}

func (controller todoController) PostTodo(c *gin.Context) {
	var newTodo NewTodo
	if err := c.ShouldBind(&newTodo); err != nil {
		errorMessage := appError.Message(err)
		res.JsonError(c, res.ErrorParams{Message: errorMessage})
		return
	}
	todo, err := controller.todoService.InsertTodo(newTodo)
	if err != nil {
		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	}
	res.JsonSuccess(c, todo)
}

func (controller todoController) PutTodoById(c *gin.Context) {
	var uri _type.Id
	var editTodo EditTodo
	if err := c.ShouldBindUri(&uri); err != nil {
		res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound})
		return
	}
	if err := c.ShouldBind(&editTodo); err != nil {
		errorMessage := appError.Message(err)
		res.JsonError(c, res.ErrorParams{Message: errorMessage})
		return
	}

	todo, err := controller.todoService.EditTodoById(uri.Id, editTodo)
	if err != nil {
		if err == pgx.ErrNoRows {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound})
			return
		}
		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	}
	res.JsonSuccess(c, todo)
}

func (controller todoController) DeleteTodoById(c *gin.Context) {
	var uri _type.Id
	if err := c.ShouldBindUri(&uri); err != nil {
		res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound})
		return
	}

	deleted := controller.todoService.DeleteTodoById(uri.Id)
	if deleted {
		res.JsonSuccess(c, nil)
		return
	}
	res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound})
}
