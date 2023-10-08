package userController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
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
	Code     string `json:"code"`
}

func Retrieve(c *gin.Context) {
	//接受参数
	var data Userpassword
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utils.JsonErrorResponse(c,200,apiExpection.ParamError.Msg)
		return
	}
	//获取用户信息
	var user *models.User
	user,err=userService.GetUserByName(data.Name)
	if err !=nil{
		utils.JsonErrorResponse(c, 200, "用户不存在")
		return
	}
	if user.Permission==1{
		utils.JsonErrorResponse(c, 200, "权限不足")
		return
	}
	//判断密码是否符合格式
	if !userService.IsValidPassword(data.Password){
		utils.JsonErrorResponse(c, 200, "密码格式错误")
		return
	}
	flag1:=userService.Compare(data.Email,user.Email)
	flag2:=userService.Compare(data.Phone,user.Phone)
	flag3:=userService.Compare(data.Code,user.Code)
	flag4:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(data.Password))
	if !flag1||!flag2{
		utils.JsonErrorResponse(c,200,"手机号或邮箱错误")
		return
	}
	if flag4 ==nil{
		utils.JsonErrorResponse(c,200,"密码与前一次相同")
		return
	}
	if !flag3{
		utils.JsonErrorResponse(c,200,"验证码错误")
		return
	}
	pwd, err := userService.Encryption(data.Password)
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