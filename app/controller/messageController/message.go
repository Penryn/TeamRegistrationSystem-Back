package messageController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/messageService"
	"TeamRegistrationSystem-Back/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserInformation(c *gin.Context){
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	v, _ := n.(int)
	terr :=messageService.CheckUserExistByUID(v)
	if terr !=nil{
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	var messageList []models.Message
	messageList,err:=messageService.GetClassList(v)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 404, "消息为空")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}
	utils.JsonSuccessResponse(c, gin.H{
		"class_list": messageList,
	})
}