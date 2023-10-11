package messageService

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/config/database"
)

func getAllUserInfo() ([]models.Userinfo, error) {
	var infoList []models.Userinfo
	result := database.DB.Find(&infoList)
	if result.Error != nil {
		return nil, result.Error
	}
	return infoList, nil
}

func getAllTeamInfo() ([]models.Team, error) {
	var infoList []models.Team
	result := database.DB.Find(&infoList)
	if result.Error != nil {
		return nil, result.Error
	}
	return infoList, nil
}
