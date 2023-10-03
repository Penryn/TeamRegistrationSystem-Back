package teamController

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/teamService"
	"TeamRegistrationSystem-Back/app/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type teamID struct{
	TeamData string `json:"team_data"`
}


func SearchTeam(c *gin.Context){
	//接受参数
	var data teamID
	err:=c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c,400,"参数错误")
		return
	}
	n,er := strconv.Atoi(data.TeamData)
	if er==nil{
		//输入为队伍id
		var teamList []models.Team
		teamList ,err =teamService.GetClassListByTeamID(n)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				utils.JsonErrorResponse(c, 404, "课程为空")
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
		teamList ,err =teamService.GetClassListByTeamName(data.TeamData)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				utils.JsonErrorResponse(c, 404, "队伍不存在")
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