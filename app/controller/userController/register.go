package userController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	//"encoding/json"

	//"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RegisterData struct {
	Name       string `json:"name" binding:"required"`
	Phone      string `json:"phone" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

// 注册
func Register(c *gin.Context) {
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	//判断手机号是否符合格式
	name_sample:=regexp.MustCompile(`^[\w\s\p{Han}]{1,25}$`)
	if !name_sample.MatchString(data.Name) {
		utils.JsonErrorResponse(c, 200, "用户名格式错误")
		return
	}
	//判断邮箱是否符合格式
	email_sample:=regexp.MustCompile(`^[\w-]+(\.[\w-]+)*@[\w-]+(\.[\w-]+)+$`)
	if !email_sample.MatchString(data.Email) {
		utils.JsonErrorResponse(c, 200, "邮箱格式错误")
		return
	}
	//判断电话是否符合格式
	phone_sample:=regexp.MustCompile(`^1[3456789]\d{9}$`)
	if !phone_sample.MatchString(data.Phone) {
		utils.JsonErrorResponse(c, 200, "电话格式错误")
		return
	}
	//判断密码是否符合格式
	if !userService.IsValidPassword(data.Password){
		utils.JsonErrorResponse(c, 200, "密码格式错误")
		return
	}
	// 判断手机号是否已经注册
	err = userService.CheckUserExistByPhone(data.Phone)
	if err == nil {
		utils.JsonErrorResponse(c, 200, "手机号已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 判断用户名是否已经注册
	err = userService.CheckUserExistByName(data.Name)
	if err == nil {
		utils.JsonErrorResponse(c, 200, "用户名已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 判断邮箱是否已经注册
	err = userService.CheckUserExistByEmail(data.Email)
	if err == nil {
		utils.JsonErrorResponse(c, 200, "邮箱已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//加密
	pwd, err :=userService.Encryption(data.Password)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// 注册用户
	err = userService.Register(models.User{
		Name:     data.Name,
		Phone:    data.Phone,
		Email:    data.Email,
		Password: pwd,
		Permission: 0,
		Userinfo: models.Userinfo{
			Name: data.Name,
			Phone: data.Phone,
			Email: data.Email,
		},
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c,nil)
	// // 自动登录
	// c.Request.Header.Set("Content-Type", "application/json")
	// c.Request.ParseForm()
	// jsonName,_:=json.Marshal(data.Name,data.Name)
	// c.Request.PostForm.Set("account", string(jsonName))
	// c.Request.PostForm.Set("password",string(jsonPassword))
	// Login(c)
}
