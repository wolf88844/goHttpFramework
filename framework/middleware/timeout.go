package middleware

import (
	"context"
	"fmt"
	"goHttpFramework/framework"
	"log"
	"time"
)

func Timeout(d time.Duration) framework.ControllerHandler {
	return func(c *framework.Context) error {
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
			c.WriteMux().Lock()
			defer c.WriteMux().Unlock()
			log.Println(p)
			c.Json(500, "panic")
		case <-finish:
			fmt.Println("finish")
		case <-durationCtx.Done():
			c.WriteMux().Lock()
			defer c.WriteMux().Unlock()
			c.Json(500, "time out")
			c.SetHasTimeout()
		}
		return nil
	}
}
