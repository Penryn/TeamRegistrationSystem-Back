package userController

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	fmt.Println("ok")
	var data RegisterData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "参数错误")
		return
	}

	// 判断手机号是否已经注册
	err = userService.CheckUserExistByPhone(data.Phone)
	if err == nil {
		utils.JsonErrorResponse(c, 400, "手机号已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 判断用户名是否已经注册
	err = userService.CheckUserExistByName(data.Name)
	if err == nil {
		utils.JsonErrorResponse(c, 400, "用户名已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	// 判断邮箱是否已经注册
	err = userService.CheckUserExistByEmail(data.Email)
	if err == nil {
		utils.JsonErrorResponse(c, 400, "邮箱已注册")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//加密
	pwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
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
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
