package todo

import (
	"go-todo/packages/user"
	"go-todo/packages/util/appError"
	"go-todo/packages/util/res"
	_type "go-todo/packages/util/type"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type todoService interface {
	GetTodoList(params GetTodoListParams, userId int) ([]TodoEntity, error)
	GetTodoCount(userId int) (int, error)
	GetTodoById(id int) (TodoEntity, error)
	InsertTodo(newTodo NewTodo) (TodoEntity, error)
	EditTodoById(id int, editTodo EditTodo) (TodoEntity, error)
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
	user := c.MustGet("user").(user.AppUserEntity)
	userId := int(user.Id)

	todoList, err := controller.todoService.GetTodoList(params, userId)

	if err != nil {
		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	}

	if err != nil {
		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	}

	if todoList == nil {
		todoList = []TodoEntity{}
	}

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

	if appUser, exists := c.Get("user"); exists {
		appUser := appUser.(user.AppUserEntity)
		userId := int(appUser.Id)

		if err := c.ShouldBind(&newTodo); err != nil {
			errorMessage := appError.Message(err)
			res.JsonError(c, res.ErrorParams{Message: errorMessage})
			return
		}
		newTodo.UserId = userId
		log.Println(userId, newTodo)
		todo, err := controller.todoService.InsertTodo(newTodo)
		if err != nil {
			res.JsonError(c, res.ErrorParams{Message: err.Error()})
			return
		}
		res.JsonSuccess(c, todo)
		return
	}

	res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized})
}

func (controller todoController) PutTodoById(c *gin.Context) {
	if appUser, exists := c.Get("user"); exists {
		var uri _type.Id
		var editTodo EditTodo
		appUser := appUser.(user.AppUserEntity)

		if err := c.ShouldBindUri(&uri); err != nil {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound})
			return
		}
		if err := c.ShouldBind(&editTodo); err != nil {
			errorMessage := appError.Message(err)
			res.JsonError(c, res.ErrorParams{Message: errorMessage})
			return
		}

		toBeEdited, err := controller.todoService.GetTodoById(uri.Id)

		if err != nil || toBeEdited.User_id != int(appUser.Id) {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized})
			return
		}

		todo, err := controller.todoService.EditTodoById(uri.Id, editTodo)
		if err != nil {
			res.JsonError(c, res.ErrorParams{Message: err.Error()})
			return
		}
		res.JsonSuccess(c, todo)
		return
	}
	res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized})
}

func (controller todoController) DeleteTodoById(c *gin.Context) {
	if appUser, exists := c.Get("user"); exists {
		var uri _type.Id
		appUser := appUser.(user.AppUserEntity)

		if err := c.ShouldBindUri(&uri); err != nil {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound})
			return
		}

		toBeEdited, err := controller.todoService.GetTodoById(uri.Id)
		if err != nil || toBeEdited.User_id != int(appUser.Id) {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized})
			return
		}

		deleted := controller.todoService.DeleteTodoById(uri.Id)
		if deleted {
			res.JsonSuccess(c, nil)
			return
		}
		res.JsonError(c, res.ErrorParams{StatusCode: http.StatusNotFound})
		return
	}
	res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized})
}
