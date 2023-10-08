package userController

//登记
import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginDate struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	//接受参数
	var data LoginDate
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
	//获取用户详细信息
	var info *models.Userinfo
	info, err = userService.GetUserInfoByUserID(user.UserID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	//判断密码是否正确
	flag := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if flag != nil {
		utils.JsonErrorResponse(c, 200, "密码错误")
		return
	}
	token, err := utils.GenToken(user.UserID)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	type ulogin struct {
		Name   string `json:"name"`
		Token  string `json:"token"`
		Avatar string `json:"avatar"`
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": ulogin{
			Name:  user.Name,
			Token: token,
			Avatar: info.Avatar,
		},
	})

}
