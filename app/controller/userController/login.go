package userController

//登记
import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"

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
	utils.JsonSuccessResponse(c,user)

}
