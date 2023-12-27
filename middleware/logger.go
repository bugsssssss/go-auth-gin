package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		clientIP := c.ClientIP()
		log.Printf("Request - IP: %s | Method: %s | Status: %d | Duration: %v", clientIP, c.Request.Method, c.Writer.Status(), duration)
	}
}
