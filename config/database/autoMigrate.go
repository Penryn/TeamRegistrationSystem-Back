package database

import (
	"TeamRegistrationSystem-Back/app/models"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Team{},
		&models.User{},
		&models.Userinfo{},
		&models.Message{},
	)
	//db.Model(&models.Team{}).Association("Users")

	return err
}
