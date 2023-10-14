package adminController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/adminService"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	"TeamRegistrationSystem-Back/config/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

/*
type GetInfoData struct {
	Name string `form:"name" binding:"required"`
}


func deleteUserAndMessages(c *gin.Context) {
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

	var now models.User
	database.DB.Where("uid = ?", m).First(&now)
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
*/

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
	uid, ok := n.(int)
	if !ok {
		utils.JsonErrorResponse(c, 200, "invalid user")
		return
	}
	terr := userService.CheckUserExistByUID(uid)
	if terr != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	// var data adminIdentify
	// err := c.ShouldBindJSON(&data)
	// if err != nil {
	// 	utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
	// 	return
	// }
	// //判断操作权限
	// if data.Permission == 0 {
	// 	utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
	// 	return
	// }

	//我不知道他们有没有（
	//还是说我应该从数据库拿

	var user models.User
	database.DB.Where("uid = ?", uid).First(&user)
	permission := user.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}

	allUserInfo, err := adminService.GetAllUserInfo()
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	for _, j := range allUserInfo {
		fmt.Println(j)
	}

	// allTeamInfo, err := adminService.GetAllTeamInfo()
	// if err != nil {
	// 	utils.JsonInternalServerErrorResponse(c)
	// 	return
	// }

	// utils.JsonSuccessResponse(c, gin.H{
	// 	"user_info": allUserInfo,
	// 	// "team_info": allTeamInfo,
	// })

}

func adminGetTeam(c *gin.Context) {
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
	var user models.User
	database.DB.Where("uid = ?", m).First(&user)
	permission := user.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}

	var op int
	err := c.ShouldBindQuery(&op)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	allTeamInfo, err := adminService.GetAllTeamInfo(op)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}

	// utils.JsonSuccessResponse(c, gin.H{
	// 	"team_info": allTeamInfo,
	// })

	for _, j := range allTeamInfo {
		fmt.Println(j)
	}
}

func adminNotice(c *gin.Context) {

}
