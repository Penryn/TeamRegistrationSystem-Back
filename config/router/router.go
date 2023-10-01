package router

import (
	"TeamRegistrationSystem-Back/app/controller/userController"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre = "/api"

	api:=r.Group(pre)
	{
		user:=api.Group("/user")
		{
			user.POST("/login",userController.Login)
			user.POST("/reg",userController.Register)
			user.PUT("/info",userController.Updateinfodata)
		}

	}
}