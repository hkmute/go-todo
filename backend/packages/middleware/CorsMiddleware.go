package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowHeaders = []string{"*"}
	config.AllowAllOrigins = true

	return cors.New(config)
}
