package adminController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/services/adminService"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

// type adminIdentify struct {
// 	Permission int `json:"permission" binding:"Required"`
// }

// 管理员界面
func AdminInterface(c *gin.Context) {

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

	user := adminService.GetUserByUserID(uid)
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

type GetTeamOpData struct {
	Op int `form:"team_confirm" binding:"required"`
}

func AdminGetTeam(c *gin.Context) {
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

	user := adminService.GetUserByUserID(m)
	permission := user.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}

	var op GetTeamOpData
	err := c.ShouldBindQuery(&op)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	allTeamInfo, err := adminService.GetAllTeamInfo(op.Op)
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

type GetMessageData struct {
	Name string `form:"information" binding:"required"`
}

func AdminMessage(c *gin.Context) {
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

	now := adminService.GetUserByUserID(m)
	permission := now.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}

	//获取信息
	var msg GetMessageData
	err := c.ShouldBindQuery(&msg)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	result := adminService.CreateMessage(msg.Name)
	if result != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.Unknown.Msg)
	}
	utils.JsonSuccessResponse(c, nil)
}
