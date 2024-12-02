package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientUrl := os.Getenv("CLIENT_URL")
		c.Header("Access-Control-Allow-Origin", clientUrl)
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}
