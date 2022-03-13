package main

import (
	"github.com/gohade/hade/framework/gin"
)

func registerRouter(core *gin.Engine) {
	core.GET("/foo", FooControllerHandler)

	core.GET("/user/login", UserLoginController)

	subjectApi := core.Group("/subject")
	{
		subjectApi.DELETE("/:id", SubjectDelController)
		subjectApi.PUT("/:id", SubjectUpdateController)
		subjectApi.GET("/:id", SubjectGetController)
		subjectApi.GET("/list/all", SubjectListController)
		subjectApi.POST("/:Id", SubjectAddController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.GET("/name", SubjectNameController)
		}
	}
}
