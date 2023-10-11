package database

import (
	"eleva/src/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MockConnection() {
	client, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("could not connect to database")
	}

	client.AutoMigrate(&models.User{}, &models.VerificationCode{})

	Client = client
}
