package messageService

import (
	"TeamRegistrationSystem-Back/app/models"
	"TeamRegistrationSystem-Back/config/database"
)

func GetClassList(userID int) ([]models.Message, error) {
	result := database.DB.Where("user_id=?", userID).Find(&models.Message{})
	if result.Error != nil {
		return nil, result.Error
	}
	var messageList []models.Message
	result = database.DB.Where("user_id=?", userID).Find(&messageList)
	if result.Error != nil {
		return nil, result.Error
	}
	return messageList, nil
}
func CheckUserExistByUID(uid int) error {
	result := database.DB.Where("user_id = ?", uid).First(&models.User{})
	return result.Error
}