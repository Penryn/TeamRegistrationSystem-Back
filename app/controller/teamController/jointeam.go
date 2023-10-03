package teamController

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/teamService"
	"TeamRegistrationSystem-Back/app/utils"

	"github.com/gin-gonic/gin"
)

type jointeamdata struct{
	TeamID int `json:"team_id"`
	TeamPassword string `json:"team_password"`

}

func JoinTeam(c *gin.Context){
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200400, "token获取失败")
		return
	}
	v, ok := n.(int)
	if !ok {
		utils.JsonErrorResponse(c, 200400, "token断言失败")
		return
	}

	//接受传参
	var data jointeamdata
	err:=c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c,400,"参数错误")
		return
	}
	//判断有无团队
	var user *models.User
	user,err =teamService.GetUserByUserID(v)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	if user.TeamID!=0{
		utils.JsonErrorResponse(c,200,"你已有团队")
		return
	}
	err= teamService.UserJoinTeam(v,data.TeamID)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c,nil)
}