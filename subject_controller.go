package main

import (
	"github.com/gohade/hade/framework/gin"
	"github.com/gohade/hade/provider/demo"
)

func SubjectAddController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok,SubjectAddController")
}

func SubjectListController(c *gin.Context) {
	demoService := c.MustMake(demo.Key).(demo.Service)
	foo := demoService.GetFoo()
	c.ISetOkStatus().IJson(foo)
}

func SubjectDelController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok,SubjectDelController")
}

func SubjectUpdateController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok,SubjectUpdateController")
}

func SubjectGetController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok,SubjectGetController")
}

func SubjectNameController(c *gin.Context) {
	c.ISetOkStatus().IJson("ok,SubjectNameController")
}
