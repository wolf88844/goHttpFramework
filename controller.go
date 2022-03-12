package main

import (
	"goHttpFramework/framework"
	"time"
)

func FooControllerHandler(c *framework.Context) error {
	time.Sleep(10 * time.Second)
	c.Json(200, "ok")
	return nil
}
