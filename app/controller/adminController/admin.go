package adminController

import (
	"TeamRegistrationSystem-Back/app/apiExpection"
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/app/services/adminService"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type deletedata struct{
	UserName  string `form:"user_name"`
}


func DeleteUserAndMessages(c *gin.Context) {
	//获取用户身份token
	n, er := c.Get("UserID")
	if !er {
		utils.JsonErrorResponse(c, 200, "token获取失败")
		return
	}
	m, _ := n.(int)
	var auser *models.User
	auser,terr := adminService.GetUserByUid(m)
	if terr != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	permission := auser.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}

	//获取要删除的用户
	var data deletedata
	err := c.ShouldBindQuery(&data)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	var user *models.User
	user,ter := adminService.GetUserByUserName(data.UserName)
	if ter != nil {
		utils.JsonErrorResponse(c, 200, "there's no such one person")
		return
	}
	if user.Permission==1{
		utils.JsonErrorResponse(c, 200, "No!!!I don't want to kill myself!")
		return
	}
	//判断用户有无队伍
	if user.TeamID!=0{
		//查询所在团队，是否存在已报名的团体
		var team *models.Team
		team,por := adminService.GetTeamByTeamID(user.TeamID)
		if por != nil {
			utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
			return
		}
		if team.Confirm == 1 {
			utils.JsonErrorResponse(c, 200, "already registed")
			return
		}
		if team.CaptainID==user.UserID{
			err :=adminService.DeleteTeam(user.TeamID,user.Name)
			if err != nil {
				utils.JsonInternalServerErrorResponse(c)
				return
			}
			return
		}
	}
	//删除用户相关信息
	err = adminService.DeleteInfoByUserID(user.UserID)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.RequestError.Msg)
		return
	}
	err =adminService.CheckMessageByUserID(user.UserID)
	if err==nil{
		err = adminService.DeleteMessageByUserID(user.UserID)
		if err != nil {
			utils.JsonErrorResponse(c, 200, apiExpection.RequestError.Msg)
			return
		}
	}
	utils.JsonSuccessResponse(c,nil)
}

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
	m, _ := n.(int)
	var auser *models.User
	auser,terr := adminService.GetUserByUid(m)
	if terr != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	permission := auser.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}
	allUser_list, err := adminService.GetUserList()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.JsonErrorResponse(c, 200, "队伍不存在")
			return
		} else {
			utils.JsonInternalServerErrorResponse(c)
			return
		}
	}

	utils.JsonSuccessResponse(c, gin.H{
		"user_list": allUser_list,
	})

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

type teamConfirmdata struct{
	Confirm int `form:"confirm"`
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
	user,err:=adminService.GetUserByUid(m)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	permission := user.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}

	var op teamConfirmdata
	err = c.ShouldBindQuery(&op)
	if op.Confirm!=0&&err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	allTeamInfo, err := adminService.GetAllTeamInfo(op.Confirm)
	if err != nil {
		utils.JsonInternalServerErrorResponse(c)
		return
	}
	utils.JsonSuccessResponse(c, gin.H{
		"team_list": allTeamInfo,
	})
	// utils.JsonSuccessResponse(c, gin.H{
	// 	"team_info": allTeamInfo,
	// })

}

type GetMessageData struct {
	Information string `form:"information" binding:"required"`
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
	user,err:=adminService.GetUserByUid(m)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}
	permission := user.Permission
	if permission == 0 {
		utils.JsonErrorResponse(c, 200, "insufficient privileges to perform the operation")
		return
	}

	//获取信息
	var msg GetMessageData
	err = c.ShouldBindQuery(&msg)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.ParamError.Msg)
		return
	}

	err = adminService.CreateMessage(msg.Information)
	if err != nil {
		utils.JsonErrorResponse(c, 200, apiExpection.Unknown.Msg)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}