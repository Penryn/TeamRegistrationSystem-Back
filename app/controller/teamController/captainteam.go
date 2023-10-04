package teamController

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/teamService"
	"TeamRegistrationSystem-Back/app/utils"

	"github.com/gin-gonic/gin"
)

// 创建队伍
type teamdata struct {
	TeamName     string `json:"team_name" binding:"required"`
	Slogan       string `json:"slogan"  binding:"required"`
	TeamPassword string `json:"team_password"  binding:"required"`
}

func CreateTeam(c *gin.Context) {
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
	var data teamdata
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}
	//判断有无团队
	var user *models.User
	user, err = teamService.GetUserByUserID(v)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	if user.TeamID != 0 {
		utils.JsonErrorResponse(c, 200, "你已有团队")
		return
	}

	err = teamService.CreateTeam(models.Team{
		TeamName:     data.TeamName,
		Slogan:       data.Slogan,
		TeamPassword: data.TeamPassword,
		Confirm:      0,
		CaptainID:    v,
		Number:       1,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}

// 踢出队员
type dissmissdata struct {
	ID     int `json:"id"  binding:"required"`
	UserID int `json:"user_id"  binding:"required"`
}

func DismissUser(c *gin.Context) {
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
	//获取参数
	var data dissmissdata
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
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
		utils.JsonErrorResponse(c, 200204, "你不是队长，权限不足")
		return
	}
	if team.Confirm != 0 {
		utils.JsonErrorResponse(c, 200204, "该团队已报名,请不要踢队友")
		return
	}
	//解除用户与团队的关联
	err = teamService.LeaveTeam(data.UserID, data.ID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}

// 解散队伍
type breakteamdata struct {
	ID int `json:"id"  binding:"required"`
}

func BreakTeam(c *gin.Context) {
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
	//获取参数
	var data breakteamdata
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
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
		utils.JsonErrorResponse(c, 200204, "你不是队长，权限不足")
		return
	}
	if team.Confirm != 0 {
		utils.JsonErrorResponse(c, 200204, "该团队已报名,请先取消报名")
		return
	}
	//删除队伍
	err = teamService.DeleteTeam(data.ID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
