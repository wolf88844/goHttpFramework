package middleware

import (
	"context"
	"fmt"
	"github.com/gohade/hade/framework/gin"
	"log"
	"net/http"
	"time"
)

func Timeout(d time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		finish := make(chan struct{}, 1)
		panicChan := make(chan interface{}, 1)

		durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
		defer cancel()

		go func() {
			defer func() {
				if p := recover(); p != nil {
					panicChan <- p
				}
			}()

			c.Next()

			finish <- struct{}{}
		}()

		select {
		case p := <-panicChan:
			log.Println(p)
			c.ISetStatus(http.StatusInternalServerError).IJson("panic")
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.ISetStatus(http.StatusInternalServerError).IJson("time out")
		}
	}
}
