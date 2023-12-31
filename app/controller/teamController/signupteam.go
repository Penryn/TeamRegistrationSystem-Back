package teamController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/teamService"
	"TeamRegistrationSystem-Back/app/utils"

	"github.com/gin-gonic/gin"
)

type signupTeamdata struct{
	ID int `form:"id"  binding:"required"`
}

func SubmitTeam(c *gin.Context){
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
	//获取参数
	var data signupTeamdata
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	//获取团队信息
	var team *models.Team
	team, err = teamService.GetTeamByTeamID(data.ID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断是否为队长
	flag := teamService.ComPaRe(team.CaptainID, v)
	if !flag {
		utils.JsonErrorResponse(c, 200, "权限不足")
		return
	}
	//判断是否符合报名条件
	if team.Number <4 || team.Number>6{
		utils.JsonErrorResponse(c, 200, "人数不符合要求")
		return
	}
	if team.Confirm!=0{
		utils.JsonErrorResponse(c, 200, "你已报名,请不要重复报名")
		return
	}
	//报名
	err=teamService.Updatasubmit(models.Team{
		ID: data.ID,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}



func CancelTeam(c *gin.Context){
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
	//获取参数
	var data signupTeamdata
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	//获取团队信息
	var team *models.Team
	team, err = teamService.GetTeamByTeamID(data.ID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断是否为队长
	flag := teamService.ComPaRe(team.CaptainID, v)
	if !flag {
		utils.JsonErrorResponse(c, 200, "权限不足")
		return
	}
	//取消报名
	err=teamService.Updatacancel(models.Team{
		ID: data.ID,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}