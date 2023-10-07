package teamController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/teamService"
	"TeamRegistrationSystem-Back/app/utils"

	"github.com/gin-gonic/gin"
)

//加入团队
type jointeamdata struct{
	ID int `json:"id"  binding:"required"`
	TeamPassword string `json:"team_password"  binding:"required"`

}

func JoinTeam(c *gin.Context){
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	v, _ := n.(int)
	terr :=teamService.CheckUserExistByUID(v)
	if terr !=nil{
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	//接受传参
	var data jointeamdata
	err:=c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c,200,apiExpection.ParamError.Msg)
		return
	}
	//获取团队信息
	var team *models.Team
	team,err=teamService.GetTeamByTeamID(data.ID)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	if team.Confirm!=0{
		utils.JsonErrorResponse(c,200,"该团队已报名")
		return
	}
	flag:=teamService.ComPare(team.TeamPassword,data.TeamPassword)
	if !flag{
		utils.JsonErrorResponse(c,200,"密码错误")
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
	//把用户和团队关联起来
	err= teamService.UserJoinTeam(v,data.ID)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c,nil)
}

//退出团队

type leaveteamdata struct{
	ID int `json:"id"  binding:"required"`

}

func LeaveTeam(c *gin.Context){
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	v, _ := n.(int)
	terr :=teamService.CheckUserExistByUID(v)
	if terr !=nil{
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	//接受传参
	var data leaveteamdata
	err:=c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c,200,"参数错误")
		return
	}
	//获取团队信息
	var team *models.Team
	team,err=teamService.GetTeamByTeamID(data.ID)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	if team.Confirm!=0{
		utils.JsonErrorResponse(c,200,"团队已报名,请勿退出")
		return
	}
	if team.CaptainID==v{
		utils.JsonErrorResponse(c,200,"你是队长，请不要退出队伍")
		return
	}
	//解除用户与团队的关联
	err=teamService.LeaveTeam(v,data.ID)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c,nil)

}