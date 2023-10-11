package messageService

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/config/database"
)

// func getAllUserInfoHandler(w http.ResponseWriter,r* http.Request){

// }
func getAllUserInfo(userID []int) ([]models.Userinfo, error) {
	// for i := 1; i <= len(userID) ; i++ {
	//过滤出包含在userID中的记录
	result := database.DB.Where("user_id IN?", userID).Find(&models.Userinfo{})
	if result.Error != nil {
		return nil, result.Error
	}
	var infoList []models.Userinfo
	result = database.DB.Where("user_id IN?", userID).Find(&infoList)
	if result.Error != nil {
		return nil, result.Error
	}
	// }
	return infoList, nil
}

func geyAllTeamInfo()

/*
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
*/
