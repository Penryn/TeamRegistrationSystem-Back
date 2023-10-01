package userController

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type uinfo struct{
	Name string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	Motto string   `json:"motto"`

}


func Updateinfodata(c *gin.Context){
	n,er:=c.Get("UserID")
	if !er{
		utils.JsonErrorResponse(c, 200400, "token获取失败")
		return
	}
	v,ok :=n.(int)
	if !ok{
		utils.JsonErrorResponse(c, 200400, "token断言失败")
		return
	}
	var data uinfo
	err := c.ShouldBindJSON(&data)
	if err !=nil {
		utils.JsonErrorResponse(c, 200400, "参数错误")
		return
	}
	//查询手机号是否重复
	err = userService.CheckUserinfoExistByPhone(data.Phone)
	if err == nil {
		utils.JsonErrorResponse(c, 400, "手机号已存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//查询邮箱是否重复注册
	err = userService.CheckUserinfoExistByEmail(data.Email)
	if err == nil {
		utils.JsonErrorResponse(c, 400, "邮箱已存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	//查询用户名是否存在
	err = userService.CheckUserinfoExistByName(data.Name)
	if err == nil {
		utils.JsonErrorResponse(c, 400, "用户名已存在")
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	var user *models.Userinfo
	user,err=userService.CheckUserinfoExistByUserid(v)
	if err !=nil{
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	

	err = userService.Updatainfo(models.Userinfo{
		ID: user.ID,
		Name: data.Name,
		Phone: data.Phone,
		Email: data.Email,
		Birthday: data.Birthday,
		Address: data.Address,
		Motto: data.Motto,
	})
	if err !=nil{
		utils.JsonInternalServerErrorResponse(c)
		fmt.Println(3)
		return
	}
	utils.JsonSuccessResponse(c,nil)
}