package main

import (
	"github.com/gohade/hade/framework/gin"
	"time"
)

func FooControllerHandler(c *gin.Context) {
	time.Sleep(10 * time.Second)
	c.ISetOkStatus().IJson("ok")
}
