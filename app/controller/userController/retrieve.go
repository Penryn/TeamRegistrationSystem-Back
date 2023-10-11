package userController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Userpassword struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

func Retrieve(c *gin.Context) {
	//接受参数
	var data Userpassword
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
	if user.Permission == 1 {
		utils.JsonErrorResponse(c, 200, "权限不足")
		return
	}
	//判断密码是否符合格式
	if !userService.IsValidPassword(data.Password) {
		utils.JsonErrorResponse(c, 200, "密码格式错误")
		return
	}
	flag1 := userService.Compare(data.Code, user.Code)
	flag2 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if flag2 == nil {
		utils.JsonErrorResponse(c, 200, "密码与前一次相同")
		return
	}
	if !flag1 {
		utils.JsonErrorResponse(c, 200, "验证码错误")
		return
	}
	pwd, err := userService.Encryption(data.Password)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	err = userService.UpdataPassword(models.User{
		UserID:   user.UserID,
		Password: pwd,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}
