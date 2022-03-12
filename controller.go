package main

import (
	"goHttpFramework/framework"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	time.Sleep(10 * time.Second)
	c.SetOkStatus().Json("ok")
	return nil
}
