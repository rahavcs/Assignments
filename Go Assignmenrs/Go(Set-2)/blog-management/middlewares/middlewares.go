package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("Request: %s %s, Duration: %v", c.Request.Method, c.Request.URL.Path, time.Since(start))
	}
}
