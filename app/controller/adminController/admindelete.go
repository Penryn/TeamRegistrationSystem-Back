package adminController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/adminService"
	"TeamRegistrationSystem-Back/app/utils"
	"TeamRegistrationSystem-Back/config/database"

	"github.com/gin-gonic/gin"
)

type GetInfoData struct {
	Name string `form:"name" binding:"required"`
}

func DeleteUserAndMessages(c *gin.Context) {
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
	terr := adminService.CheckUserExistByUID(m)
	if terr != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	var now models.User
	database.DB.Where("user_id = ?", m).First(&now)
	permission := now.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}

	//获取要删除的用户
	// var uid int
	var username GetInfoData
	err := c.ShouldBindQuery(&username)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	// var user models.User
	// var uid int
	user, qer := adminService.GetUserByName(username.Name)
	if qer != nil {
		utils.JsonErrorResponse(c, 200, "there's no such one person")
		return
	}
	uid := user.UserID

	//判断是否为管理员本人
	if uid == m {
		utils.JsonErrorResponse(c, 200, "No!!!I don't want to kill myself!")
		return
	}

	//查询所在团队，是否存在已报名的团体
	exs := adminService.CheckTeamExist(uid)
	if exs == 1 {
		por := adminService.CheckTeamByUserID(uid)
		if por == 1 {
			utils.JsonErrorResponse(c, 200, "already registed")
			return
		} else if por == 2 {
			utils.JsonErrorResponse(c, 200, "teamCaptain nonono")
			return
		}
	}

	//否->清空用户与团队的关联
	rerr := adminService.DeleteRelevantTeamInfo(uid)
	if rerr != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.RequestError.Msg)
		return
	}

	//删除用户相关信息
	err = adminService.DeleteInfoByUserID(uid)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.RequestError.Msg)
		return
	}
}
