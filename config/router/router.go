package router

import (
	"TeamRegistrationSystem-Back/app/controller/userController"
	"TeamRegistrationSystem-Back/app/midwares"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	r.POST("/login",userController.Login)
	r.POST("/reg",userController.Register)
	r.PUT("/ret",userController.Retrieve)
	const pre = "/api"

	api:=r.Group(pre)
	{
		user:=api.Group("/user").Use(midwares.JWTAuthMiddleware())
		{
			user.PUT("/info",userController.Updateinfodata)
			user.PUT("/avatar",userController.AvatarUpload)
			user.GET("/info",userController.GetUserInfo)
		}

	}
}