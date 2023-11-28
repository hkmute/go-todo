package server

import (
	"go-todo/packages/server/router"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	router.Init(r)

	r.Run()
}
