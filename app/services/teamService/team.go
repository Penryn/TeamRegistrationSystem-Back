package teamService

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/config/database"
)

func CreateTeam(team models.Team)error{
	result:=database.DB.Create(&team)
	return result.Error
}

func GetUserByUserID(userid int) (*models.User, error) {
	var user models.User
	result := database.DB.Where("user_id = ?", userid).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func GetClassListByTeamID(teamID int) ([]models.Team, error) {
	result := database.DB.Where("team_id=?", teamID).Find(&models.Team{})
	if result.Error != nil {
		return nil, result.Error
	}
	var teamList []models.Team
	result = database.DB.Omit("avatar").Where("team_id=?", teamID).Find(&teamList)
	if result.Error != nil {
		return nil, result.Error
	}
	return teamList, nil
}

func GetClassListByTeamName(teamname string) ([]models.Team, error) {
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

func UserJoinTeam(uid,tid int)error{
	var user models.User
	database.DB.Take(&user,uid)
	var team models.Team
	database.DB.Take(&team,tid)
	err:=database.DB.Model(&team).Association("Users").Append(&user)
	return err
}
