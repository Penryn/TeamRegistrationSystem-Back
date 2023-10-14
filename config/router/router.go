package router

import (
	"TeamRegistrationSystem-Back/app/controller/adminController"
	"TeamRegistrationSystem-Back/app/controller/messageController"
	"TeamRegistrationSystem-Back/app/controller/teamController"
	"TeamRegistrationSystem-Back/app/controller/userController"
	"TeamRegistrationSystem-Back/app/midwares"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine){
	const pre = "/api"

	api:=r.Group(pre)
	{	api.POST("/login",userController.Login)
		api.POST("/reg",userController.Register)
		api.PUT("/ret",userController.Retrieve)
		api.POST("/email",userController.Sendmail)
		user:=api.Group("/user").Use(midwares.JWTAuthMiddleware())
		{
			user.PUT("/info",userController.Updateinfodata)
			user.POST("/avatar",userController.AvatarUpload)
			user.GET("/info",userController.GetUserInfo)
			user.GET("/message",messageController.GetUserInformation)
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
			team.POST("/avatar",teamController.TeamAvatarUpload)	
		}
		admin:=api.Group("/admin").Use(midwares.JWTAuthMiddleware())
		{
			admin.GET("/user", adminController.AdminInterface)
			admin.GET("/team", adminController.AdminGetTeam)
			admin.DELETE("/delete", adminController.DeleteUserAndMessages)
			admin.POST("/message", adminController.AdminMessage)
		}

	}
}