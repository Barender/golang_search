package logger

import (
	"log"
	"github.com/gin-gonic/gin"
)

// RequestLoggerMiddleware is a middleware to log network requests
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the incoming request
		log.Printf("Request: %s %s", c.Request.URL.Path , c.Request.Method)
		c.Next()
	}
}
