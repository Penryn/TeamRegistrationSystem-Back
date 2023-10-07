package teamService

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/config/database"
	"fmt"
	"time"
)

func CreateTeam(team models.Team)error{
	var user models.User
	database.DB.Take(&user,team.CaptainID)
	result:=database.DB.Create(&models.Team{
		TeamName:team.TeamName,
		CaptainID:team.CaptainID,
		Confirm:team.Confirm,
		Slogan:team.Slogan,
		TeamPassword: team.TeamPassword,
		Number: team.Number,
		Users: []models.User{user},
		
	})
	return result.Error
}

func GetTeamByTeamID(tid int) (*models.Team, error) {
	var team models.Team
	result := database.DB.Where("id = ?", tid).First(&team)
	if result.Error != nil {
		return nil, result.Error
	}
	return &team, nil
}

func ComPare(t1 string, t2 string) bool {
	return t1 == t2
}

func ComPaRe(t1 int, t2 int) bool {
	return t1 == t2
}

func GetUserByUserID(userid int) (*models.User, error) {
	var user models.User
	result := database.DB.Where("user_id = ?", userid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetTeamListByTeamID(teamID int) ([]models.Team, error) {
	result := database.DB.Where("id=?", teamID).Find(&models.Team{})
	if result.Error != nil {
		return nil, result.Error
	}
	var teamList []models.Team
	result = database.DB.Omit("avatar").Where("id=?", teamID).Find(&teamList)
	if result.Error != nil {
		return nil, result.Error
	}
	return teamList, nil
}

func GetTeamMoreListByTeamID(teamID int) ([]models.Team, error) {
	result := database.DB.Where("id=?", teamID).Find(&models.Team{})
	if result.Error != nil {
		return nil, result.Error
	}
	var teamList []models.Team
	result = database.DB.Preload("Users").Where("id=?", teamID).Find(&teamList)
	if result.Error != nil {
		return nil, result.Error
	}
	return teamList, nil
}

func GetTeamListByTeamName(teamname string) ([]models.Team, error) {
	result := database.DB.Where("team_name=?", teamname).Find(&models.Team{})
	if result.Error != nil {
		return nil, result.Error
	}
	var teamList []models.Team
	result = database.DB.Omit("avatar").Where("team_name=?", teamname).Find(&teamList)
	if result.Error != nil {
		return nil, result.Error
	}
	return teamList, nil
}
//
func UserJoinTeam(uid,tid int)error{
	var user models.User
	database.DB.Take(&user,uid)
	var team models.Team
	database.DB.Take(&team,tid)
	uSer,_:=GetUserByUserID(uid)
	CreateMessage(tid,uSer.Name+"加入队伍")
	err:=database.DB.Model(&team).Association("Users").Append(&user)
	UpdateUserNumber(tid)
	return err
}
//
func Updateteaminfo(team models.Team) error {
	result :=database.DB.Model(&team).Updates(models.Team{TeamName: team.TeamName,Slogan: team.Slogan})
	CreateMessage(team.ID,"队伍基本信息已更新")
	UpdateUserNumber(team.ID)
	return result.Error
}
//
func UpdataTeamAvatar(info models.Team)error{
	result :=database.DB.Model(&info).Update("avatar",info.Avatar)
	CreateMessage(info.ID,"队伍头像已更新")
	UpdateUserNumber(info.ID)
	return result.Error
}
//
func LeaveTeam(uid,tid int)error{
	var team models.Team
	database.DB.Take(&team,tid)
	var user models.User
	database.DB.Take(&user,uid)
	uSer,_:=GetUserByUserID(uid)
	CreateMessage(tid,uSer.Name+"离开队伍")
	err:=database.DB.Model(&team).Association("Users").Delete(&user)
	UpdateUserNumber(tid)
	return err
}

func DeleteTeam(tid int)error{
	var team models.Team
	database.DB.Preload("Users").Take(&team,tid)
	CreateMessage(tid,"你的队伍被解散")
	database.DB.Model(&team).Association("Users").Delete(&team.Users)
	result:=database.DB.Delete(&team)
	return result.Error
}

func GetTeamMoreByTeamID(tid int) (*models.Team, error) {
	var team models.Team
	result := database.DB.Preload("Users").Where("id = ?", tid).First(&team)
	if result.Error != nil {
		return nil, result.Error
	}
	return &team, nil
}


func UpdateUserNumber(tid int)error{
	team,err:=GetTeamMoreByTeamID(tid)
	if err!=nil{
		return err
	}
	//fmt.Println(team.Users)
	//num :=len(team.Users)
	j:=0
	for range team.Users{
		j++
	}
	result :=database.DB.Model(&models.Team{ID: tid}).Updates(models.Team{Number: j})
	return result.Error

}

func Updatasubmit(team models.Team)error{
	result :=database.DB.Model(&team).Update("confirm",1)
	CreateMessage(team.ID,"你的队伍确定报名")
	return result.Error
}

func Updatacancel(team models.Team)error{
	result :=database.DB.Model(&team).Update("confirm",0)
	CreateMessage(team.ID,"你的队伍取消报名")
	return result.Error
}

func GetUserIDSliceByTeamID(tid int) ([]int, error) {
	var user []models.User
	result := database.DB.Where("team_id = ?", tid).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	num :=[]int{}
	for _,j:=range user{
		fmt.Println(j.UserID)
		num=append(num,j.UserID)
	}
	return num, nil
}

func CreateMessage(tid int,notice string)error{
	num,err:=GetUserIDSliceByTeamID(tid)
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

func CheckUserExistByUID(uid int) error {
	result := database.DB.Where("user_id = ?", uid).First(&models.User{})
	return result.Error
}