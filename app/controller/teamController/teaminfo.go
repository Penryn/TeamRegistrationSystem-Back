package teamController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/teamService"
	"TeamRegistrationSystem-Back/app/utils"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// 获取团队信息
type Getteaminfodata struct {
	ID int `form:"id"  binding:"required"`
}

func GetTeamInfo(c *gin.Context) {
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	v, _ := n.(int)
	var user *models.User
	user, terr := teamService.GetUserByUserID(v)
	if terr != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	//接受数据
	var data Getteaminfodata
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	var signed int
	if user.TeamID == 0 {
		signed = 0
	} else {
		var team *models.Team
		team, terrr := teamService.GetTeamByTeamID(user.TeamID)
		if terrr != nil {
			utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
			return
		}
		if data.ID == user.TeamID && v == team.CaptainID {
			signed = 2
		} else if data.ID == user.TeamID && v != team.CaptainID {
			signed = 1
		} else {
			signed = 0
		}
	}
	var TeamInfoList *models.Team
	TeamInfoList, err = teamService.GetTeamMoreByTeamID(data.ID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200, "队伍为空")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	type Team_info struct {
		ID          int           `json:"id"  gorm:"foreignkey:TeamID"`
		Signed      int           `json:"signed"`
		TeamName    string        `json:"team_name"`
		CaptainName string        `json:"captain_name"`
		Slogan      string        `json:"slogan"`
		Avatar      string        `json:"avatar"`
		Confirm     int           `json:"confirm"`
		Number      int           `json:"number"`
		Users       []models.User `json:"users"`
	}
	utils.JsonResponse(c, 200, 200, "ok", gin.H{
		"team_info": Team_info{
			ID:          TeamInfoList.ID,
			Signed:      signed,
			TeamName:    TeamInfoList.TeamName,
			CaptainName: TeamInfoList.CaptainName,
			Slogan:      TeamInfoList.Slogan,
			Avatar:      TeamInfoList.Avatar,
			Confirm:     TeamInfoList.Confirm,
			Number:      TeamInfoList.Number,
			Users:       TeamInfoList.Users,
		}})

}

// 更新基本信息
type UpdateInfoData struct {
	ID       int    `json:"id"  binding:"required"`
	TeamName string `json:"team_name"  binding:"required"`
	Slogan   string `json:"slogan"  binding:"required"`
}

func UpdateTeamInfo(c *gin.Context) {
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	v, _ := n.(int)
	terr := teamService.CheckUserExistByUID(v)
	if terr != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	//获取参数
	var data UpdateInfoData
	err := c.ShouldBindJSON(&data)
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
		utils.JsonErrorResponse(c, 200, "你不是队长，权限不足")
		return
	}
	//判断是否符合格式
	name_sample := regexp.MustCompile(`^.*\D.*$`)
	if !name_sample.MatchString(data.TeamName) || len(data.TeamName) > 10 {
		utils.JsonErrorResponse(c, 200, "队伍名称格式错误")
		return
	}
	if len(data.Slogan) > 1000 {
		utils.JsonErrorResponse(c, 200, "队伍口号格式错误")
		return
	}
	//更新信息
	err = teamService.Updateteaminfo(models.Team{
		ID:       data.ID,
		TeamName: data.TeamName,
		Slogan:   data.Slogan,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}

// 更新头像
func TeamAvatarUpload(c *gin.Context) {
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	v, _ := n.(int)
	terr := teamService.CheckUserExistByUID(v)
	if terr != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	//获取用户信息
	var user *models.User
	user, err := teamService.GetUserByUserID(v)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//获取团队信息
	var team *models.Team
	team, err = teamService.GetTeamByTeamID(user.TeamID)
	if err != nil {
		utils.JsonErrorResponse(c, 200, "团队不存在")
		return
	}
	//判断是否为队长
	flag := teamService.ComPaRe(team.CaptainID, v)
	if !flag {
		utils.JsonErrorResponse(c, 200, "你不是队长，权限不足")
		return
	}
	//保存图片文件
	file, err := c.FormFile("image")
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 创建临时目录
	tempDir, err := os.MkdirTemp("", "tempdir")
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	defer os.RemoveAll(tempDir) // 在处理完之后删除临时目录及其中的文件
	// 在临时目录中创建临时文件
	tempFile := filepath.Join(tempDir, file.Filename)
	f, err := os.Create(tempFile)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	defer f.Close()

	// 将上传的文件保存到临时文件中
	src, err := file.Open()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	defer src.Close()

	_, err = io.Copy(f, src)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 判断文件的MIME类型是否为图片
	mime, err := mimetype.DetectFile(tempFile)
	if err != nil || !strings.HasPrefix(mime.String(), "image/") {
		utils.JsonErrorResponse(c, 200, "仅允许上传图片文件")
		return
	}

	filename := uuid.New().String() + filepath.Ext(file.Filename)
	dst := "./uploads/" + filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	url := "47.115.209.120:8080" + "/uploads/" + filename
	err = teamService.UpdataTeamAvatar(models.Team{
		ID:     user.TeamID,
		Avatar: "http://"+url,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
