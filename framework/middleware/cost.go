package middleware

import (
	"github.com/gohade/hade/framework/gin"
	"log"
	"time"
)

func Cost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.Printf("api uri start: %v", c.Request.RequestURI)
		c.Next()
		end := time.Now()
		cost := end.Sub(start)
		log.Printf("api uri: %v,cost: %v", c.Request.RequestURI, cost.Seconds())
	}
}
