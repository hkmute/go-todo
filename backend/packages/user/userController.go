package user

import (
	"go-todo/packages/util"
	"go-todo/packages/util/res"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type userService interface {
	GetUser(params GetUserParams) (AppUserEntity, error)
	GetUserToken(params GetUserParams) (string, error)
	CreateUser(newUser NewUser) (AppUserEntity, error)
	EditUser(id int, editUser EditUser) (AppUserEntity, error)
	DeleteUser(id int) bool
}

type userController struct {
	userService userService
}

func NewUserController(userService userService) userController {
	return userController{userService: userService}
}

func (controller userController) GetMe(c *gin.Context) {
	var params GetUserParams

	if user, exists := c.Get("user"); exists {
		user := user.(AppUserEntity)
		userId := int(user.Id)
		params = GetUserParams{Id: &userId}
	}

	user, err := controller.userService.GetUser(params)

	if err != nil {
		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	}

	res.JsonSuccess(c, user)
}

func (controller userController) Login(c *gin.Context) {
	var loginUser LoginUser

	if err := c.ShouldBind(&loginUser); err != nil {
		res.JsonError(c, res.ErrorParams{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	user, err := controller.userService.GetUser(GetUserParams{Username: &loginUser.Username, WithPassword: true})

	if err != nil {
		if err == pgx.ErrNoRows {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized, Message: "Invalid username or password"})
			return
		}

		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	}

	if util.CheckPasswordHash(loginUser.Password, user.Password) == false {
		res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized, Message: "Invalid username or password"})
		return
	}

	userId := int(user.Id)
	token, err := controller.userService.GetUserToken(GetUserParams{Id: &userId})
	if err != nil {
		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	}
	res.JsonSuccess(c, token)
}

func (controller userController) Register(c *gin.Context) {
	var newUser NewUser

	if err := c.ShouldBind(&newUser); err != nil {
		res.JsonError(c, res.ErrorParams{StatusCode: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if user, err := controller.userService.GetUser(GetUserParams{Username: &newUser.Username}); err != nil && err != pgx.ErrNoRows {
		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	} else if user.Username != "" {
		res.JsonError(c, res.ErrorParams{StatusCode: http.StatusBadRequest, Message: "Username already exists"})
		return
	}

	user, err := controller.userService.CreateUser(newUser)
	if err != nil {
		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	}

	userId := int(user.Id)
	token, err := controller.userService.GetUserToken(GetUserParams{Id: &userId})
	if err != nil {
		res.JsonError(c, res.ErrorParams{Message: err.Error()})
		return
	}
	user.Token = token

	res.JsonSuccess(c, user)
}
