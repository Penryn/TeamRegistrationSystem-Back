package userController

//登记
import (
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
		utils.JsonErrorResponse(c,400,"参数错误")
		return
	}


	//判断用户是否存在
	err = userService.CheckUserExistByAccount(data.Account)
	if err!=nil{
		if err ==gorm.ErrRecordNotFound{
			utils.JsonErrorResponse(c,401,"用户不存在")
		}else{
			utils.JsonInternalServerErrorResponse(c)
		}
		return
	}
	//获取用户信息
	var user *models.User
	user,err=userService.GetUserByAccount(data.Account)
	if err !=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}


	//判断密码是否正确
	flag :=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(data.Password))
	if flag !=nil{
		utils.JsonErrorResponse(c,402,"密码错误")
		return
	}
	token ,err:=utils.GenToken(user.UserID)
	if err !=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	type ulogin struct{
		UserID int `json:"user_id"`
		Name   string `json:"name"`
		Token   string `jsonL:"token"`
	}
	var nn ulogin

	nn.Name=user.Name
	nn.Token=token
	nn.UserID=user.UserID

	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"ok",
		"data":nn,

	})

}
