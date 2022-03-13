package main

import (
	"goHttpFramework/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	foo, _ := c.QueryString("foo", "def")
	time.Sleep(20 * time.Second)
	c.SetOkStatus().Json("ok,UserLoginController:" + foo)
	return nil
}
