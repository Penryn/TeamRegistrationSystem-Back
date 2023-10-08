package userController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/emailService"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	"time"

	"github.com/gin-gonic/gin"
)


type Emaildata struct{
	Email string `json:"email"`
}

var lastRequestTime time.Time

func Sendmail(c *gin.Context){
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
	//获取用户信息
	var user *models.User
	user,err=userService.GetUserByEmail(data.Email)
	if err !=nil{
		utils.JsonErrorResponse(c, 200, "用户不存在")
		return
	}
	//获取验证码
	code:=emailService.RandCode()
	//储存验证码
	err=userService.UpdataCode(models.User{
		UserID: user.UserID,
		Code: code,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//发送验证码
	err=emailService.MailSendCode(data.Email,code)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 将当前时间设置为最新的请求时间
	lastRequestTime = currentTime
	utils.JsonSuccessResponse(c, nil)

}