package teamController

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/teamService"
	"TeamRegistrationSystem-Back/app/utils"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)


//获取团队信息
type Getteaminfodata struct {
	ID int `json:"id"`
}

func GetTeamInfo(c *gin.Context) {
	var data Getteaminfodata
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}
	var TeamInfoList []models.Team
	TeamInfoList, err = teamService.GetTeamMoreListByTeamID(data.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 404, "队伍为空")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	utils.JsonSuccessResponse(c, gin.H{
		"class_list": TeamInfoList,
	})

}

//更新基本信息
type UpdateInfoData struct {
	ID      int  `json:"id"`
	TeamName string `json:"team_name"`
	Slogan   string `json:"slogan"`
}

func UpdateTeamInfo(c *gin.Context) {
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
	var data UpdateInfoData
	err:=c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c,400,"参数错误")
		return
	}
	//获取团队信息
	var team *models.Team
	team,err=teamService.GetTeamByTeamID(data.ID)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断是否为队长
	flag:=teamService.ComPaRe(team.CaptainID,v)
	if !flag{
		utils.JsonErrorResponse(c,200204,"你不是队长，权限不足")
		return
	}
	//更新信息
	err=teamService.Updateteaminfo(models.Team{
		ID: data.ID,
		TeamName: data.TeamName,
		Slogan: data.Slogan,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}


//更新头像
func TeamAvatarUpload(c *gin.Context){
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
	//获取用户信息
	var user *models.User
	user,err :=teamService.GetUserByUserID(v)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//获取团队信息
	var team *models.Team
	team,err=teamService.GetTeamByTeamID(user.TeamID)
	if err!=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//判断是否为队长
	flag:=teamService.ComPaRe(team.CaptainID,v)
	if !flag{
		utils.JsonErrorResponse(c,200204,"你不是队长，权限不足")
		return
	}
	//保存图片文件
	file, err := c.FormFile("image")
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	dst := "./uploads/" + filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	url := c.Request.Host + "/uploads/" + filename
	err =teamService.UpdataTeamAvatar(models.Team{
		ID: user.TeamID,
		Avatar: url,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}