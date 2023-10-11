package adminService

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/config/database"
)

/*
	func isAdmin(UserID int)(int,error){
		var user models.User
		err := database.DB.First(&user, userID).Error
		if err != nil {
			return false, err
		}
		return user.Permission, nil
	}

}
*/
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

func deleteInfoByUserID(userID int) error {
	//?
	err := database.DB.Where("uid = ?", userID).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
