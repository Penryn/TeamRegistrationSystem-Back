package adminService

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/config/database"
	//"fmt"
	"time"
)


func GetUserList() ([]models.User, error) {
	result := database.DB.Find(&models.User{})
	if result.Error != nil {
		return nil, result.Error
	}
	var userList []models.User
	result = database.DB.Find(&userList)
	if result.Error != nil {
		return nil, result.Error
	}
	return userList, nil
}

func GetAllTeamInfo(op int) ([]models.Team, error) {
	var infoList []models.Team
	result := database.DB.Where("confirm = ?", op).Find(&infoList)
	if result.Error != nil {
		return nil, result.Error
	}
	return infoList, nil
}

func DeleteInfoByUserID(userID int) error {
	var user models.User
	database.DB.Take(&user,userID)
	// err := database.DB.Model(&user).Association("UserInfo").Clear()
	database.DB.Where("user_id=?",userID).Delete(&models.Userinfo{UserID: userID})
	database.DB.Delete(&models.User{UserID: userID})
	// if err != nil {
	// 	return err

	// }
	return nil
}

func DeleteMessageByUserID(userID int) error {
	result := database.DB.Where("user_id = ?", userID).Delete(&models.Message{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

/*
func CheckTeamByUserID(userID int) int {
	// var userTeam []models.Team
	var teamCon []int
	// result :=
	database.DB.Where("user_id = ?", userID).Select("confirm").Find(&teamCon)
	// if result.Error != nil {
	// 	return 0,result.Error
	// }
	// var flat bool
	flat := false
	for i := 0; i < len(teamCon); i++ {
		if teamCon[i] == 1 {
			flat = true
		}
	}
	if !flat || len(teamCon) == 0 {
		return 1
	}
	return 0
}
*/

func CheckMessageByUserID(userID int) error {
	var userMessage models.Message
	// result :=
	result:=database.DB.Where("user_id = ?", userID).Find(&userMessage)
	if result.Error !=nil {
		return result.Error
	}
	return nil
}


func GetTeamByTeamID(tid int) (*models.Team, error) {
	var team models.Team
	result := database.DB.Preload("Users").Where("id = ?", tid).Find(&team)
	if result.Error != nil {
		return nil, result.Error
	}
	return &team, nil
}

func UpdateUserNumber(tid int) error {
	team, err := GetTeamByTeamID(tid)
	if err != nil {
		return err
	}
	j := 0
	for range team.Users {
		j++
	}
	result := database.DB.Model(&models.Team{ID: tid}).Updates(models.Team{Number: j})
	return result.Error

}

func DeleteRelevantTeamInfo(uid int) error {
	// var userTeam []models.Team
	var userTeam models.User
	result := database.DB.Where("user_id = ?", uid).Find(&userTeam)
	if result.Error != nil {
		return result.Error
	}
	var team models.Team
	var user models.User
	database.DB.Take(&user, uid)
	// for i := 0; i < len(userTeam); i++ {
	database.DB.Take(&team, userTeam.TeamID)
	database.DB.Model(&team).Association("Users").Delete(&user)
	UpdateUserNumber(userTeam.TeamID)
	// }
	return nil
}

func GetAllUserID() ([]int, error) {
	var user []models.User
	result := database.DB.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	var ids []int
	for _, uids := range user {
		ids = append(ids, int(uids.UserID))
	}
	return ids, nil
}

func CreateMessage(notice string)error{
	num,err:=GetUserIDSlice()
	if err !=nil{
		return err
	}
	for _,i:=range num{
		mess :=models.Message{
			UserID: i,
			Information: notice,
			Time: time.Now().Format("2006-01-02 15:04"),
		}
		database.DB.Create(&mess)
	}
	return nil
}

func GetUserByUid(uid int)(*models.User,error){
	var user models.User
	result :=database.DB.Where("user_id = ?",uid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUseInforByUid(uid int)(*models.Userinfo,error){
	var user models.Userinfo
	result :=database.DB.Where("user_id = ?",uid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUserName(uname string)(*models.User,error){
	var user models.User
	result :=database.DB.Where("name = ?",uname).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func DeleteTeam(tid int,uname string)error{
	var team models.Team
	database.DB.Preload("Users").Take(&team,tid)
	CreateMessage(uname+"被删除，该所属队伍被解散")
	database.DB.Model(&team).Association("Users").Delete(&team.Users)
	result:=database.DB.Delete(&team)
	return result.Error
}

func GetUserIDSlice() ([]int, error) {
	var user []models.User
	result := database.DB.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	num :=[]int{}
	for _,j:=range user{
		num=append(num,j.UserID)
	}
	return num, nil
}

func CheckUserExistByUID(uid int) error {
	result := database.DB.Where("user_id = ?", uid).First(&models.User{})
	return result.Error
}