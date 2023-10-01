package userController

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Userpassword struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func Retrieve(c *gin.Context) {
	//接受参数
	var data Userpassword
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c,400,"参数错误")
		return
	}
	//判断用户是否存在
	err = userService.CheckUserExistByName(data.Name)
	if err != nil {
		utils.JsonErrorResponse(c, 400, "用户不存在")
		return
	} 
	//获取用户信息
	var user *models.User
	user,err=userService.GetUserByName(data.Name)
	if err !=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	flag1:=userService.Compare(data.Email,user.Email)
	flag2:=userService.Compare(data.Phone,user.Phone)
	flag3:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(data.Password))
	if !flag1||!flag2{
		utils.JsonErrorResponse(c,200220,"手机号或邮箱错误")
		return
	}
	if flag3 ==nil{
		utils.JsonErrorResponse(c,402,"密码与前一次相同")
		return
	}
	pwd, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	err =userService.UpdataPassword(models.User{
		UserID: user.UserID,
		Password: pwd,
	})
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, nil)

}