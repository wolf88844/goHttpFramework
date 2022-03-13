package main

import (
	"github.com/gohade/hade/framework/gin"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "def")
	c.ISetOkStatus().IJson("ok,UserLoginController:" + foo)
}
