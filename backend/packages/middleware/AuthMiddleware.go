package middleware

import (
	"go-todo/packages/db"
	"go-todo/packages/user"
	"go-todo/packages/util"
	"go-todo/packages/util/res"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		token := strings.Split(authHeader, "Bearer ")[1]
		if token == "" {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}

		claims, err := util.ParseJwt(token)

		if err != nil {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}

		userService := user.NewUserService(db.DB.Conn())

		userId, err := strconv.Atoi(claims.Subject)
		if err != nil {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}
		user, err := userService.GetUser(user.GetUserParams{Id: &userId})
		if err != nil {
			res.JsonError(c, res.ErrorParams{StatusCode: http.StatusUnauthorized})
			c.Abort()
			return
		}
		c.Set("user", user)
		c.Next()
	}
}
