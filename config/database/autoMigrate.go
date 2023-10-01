package database

import (
	"TeamRegistrationSystem-Back/app/models"

	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Userinfo{},
	)

	return err
}
