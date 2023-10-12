package adminController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/adminService"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	"TeamRegistrationSystem-Back/config/database"

	"github.com/gin-gonic/gin"
)

func deleteUserAndMessages(c *gin.Context) error {
	//获取用户身份token
	//判断是否为管理员
	//查询所在团队，是否存在正好5人的团体
	//是->删除用户相关信息
	//清空用户与团队的关联
	//删除用户
}

// type adminIdentify struct {
// 	Permission int `json:"permission" binding:"Required"`
// }

// 管理员界面
func adminInterface(c *gin.Context) {
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	m, ok := n.(int)
	if !ok {
		utils.JsonErrorResponse(c, 200, "invalid user")
		return
	}
	terr := userService.CheckUserExistByUID(m)
	if terr != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	/*
		var data adminIdentify
		err := c.ShouldBindJSON(&data)
		if err != nil {
			utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
			return
		}
		//判断操作权限
		if data.Permission == 0 {
			utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
			return
		}
	*/
	//我不知道他们有没有（
	//还是说我应该从数据库拿

	var user models.User
	database.DB.Where("uid = ?", m).First(&user)
	permission := user.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}
	//
	//?
	allUserInfo, err := adminService.GetAllUserInfo()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	allTeamInfo, err := adminService.GetAllTeamInfo()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	utils.JsonSuccessResponse(c, gin.H{
		"user_info": allUserInfo,
		"team_info": allTeamInfo,
	})
}
