package teamController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/teamService"
	"TeamRegistrationSystem-Back/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type teamID struct{
	TeamData string `json:"team_data"  binding:"required"`
}


func SearchTeam(c *gin.Context){
	//接受参数
	var data teamID
	err:=c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c,200,apiExpection.ParamError.Msg)
		return
	}
	n,er := strconv.Atoi(data.TeamData)
	if er==nil{
		//输入为队伍id
		var teamList []models.Team
		teamList ,err =teamService.GetTeamListByTeamID(n)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				utils.JsonErrorResponse(c, 200, "队伍为不存在")
				return
			} else {
				utils.JsonInternalServerErrorResponse(c)
				return
			}
		}
		utils.JsonSuccessResponse(c, gin.H{
			"team_list": teamList,
		})
	}else{
		//输入为队伍名称
		var teamList []models.Team
		teamList ,err =teamService.GetTeamListByTeamName(data.TeamData)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				utils.JsonErrorResponse(c, 200, "队伍不存在")
				return
			} else {
				utils.JsonInternalServerErrorResponse(c)
				return
			}
		}
		utils.JsonSuccessResponse(c, gin.H{
			"team_list": teamList,
		})
	}


}