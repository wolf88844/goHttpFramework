package main

import "goHttpFramework/framework"

func UserLoginController(c *framework.Context) error {
	c.SetOkStatus().Json("ok,UserLoginController")
	return nil
}
