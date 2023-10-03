package router

import (
	"TeamRegistrationSystem-Back/app/controller/teamController"
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
		team:=api.Group("/team").Use(midwares.JWTAuthMiddleware())
		{
			team.POST("/create",teamController.CreateTeam)
			team.DELETE("/delete",teamController.BreakTeam)
			team.DELETE("/dismiss",teamController.DismissUser)
			team.DELETE("",teamController.LeaveTeam)
			team.POST("",teamController.JoinTeam)
			team.GET("",teamController.SearchTeam)
			team.GET("/info",teamController.GetTeamInfo)
			team.PUT("/info",teamController.UpdateTeamInfo)
			team.PUT("/cancel",teamController.CancelTeam)
			team.PUT("/submit",teamController.SubmitTeam)
			team.PUT("/avatar",teamController.TeamAvatarUpload)
			
		}

	}
}