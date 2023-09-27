package userService

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/config/database"
)


func CheckUserExistByPhone(pHone string) error {
	result := database.DB.Where("phone = ?", pHone).First(&models.User{})
	return result.Error
}

func CheckUserExistByEmail(eMail string) error {
	result := database.DB.Where("email = ?", eMail).First(&models.User{})
	return result.Error
}

func CheckUserExistByName(naMe string) error {
	result := database.DB.Where("name = ?", naMe).First(&models.User{})
	return result.Error
}

func CheckUserExistByAccount(account string) error {
	result := database.DB.Where("phone = ? or email =? or name=?", account,account,account).First(&models.User{})
	return result.Error
}

func GetUserByAccount(aCCount string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("phone = ? or email =? or name=?", aCCount,aCCount,aCCount).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func Register(user models.User) error {
	result := database.DB.Create(&user)
	return result.Error
}
