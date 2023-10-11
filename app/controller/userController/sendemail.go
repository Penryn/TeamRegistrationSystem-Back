package userController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/emailService"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Emaildata struct {
	Account string `json:"account"`
}

var lastRequestTime time.Time

func Sendmail(c *gin.Context) {
	// 获取当前时间
	currentTime := time.Now()

	// 检查时间间隔是否满足限制
	if currentTime.Sub(lastRequestTime) < 60*time.Second {
		utils.JsonErrorResponse(c, 200, "请求频率过高，请稍后再试")
		return
	}
	//接收数据
	var data Emaildata
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	//判断用户是否存在
	err = userService.CheckUserExistByAccount(data.Account)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200, "用户不存在")
		} else {
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}
	//获取用户信息
	var user *models.User
	user, err = userService.GetUserByAccount(data.Account)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//获取验证码
	code := emailService.RandCode()
	//储存验证码
	err = userService.UpdataCode(models.User{
		UserID: user.UserID,
		Code:   code,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//发送验证码
	err = emailService.MailSendCode(user.Email, code)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	type puser struct {
		Name  string `json:"name"`
		Email string `json:"eamil"`
	}
	// 将当前时间设置为最新的请求时间
	lastRequestTime = currentTime
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "发送成功",
		"data": puser{
			Name: user.Name,
			Email: user.Email,
		},
	})

}
