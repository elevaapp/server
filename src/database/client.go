package database

import (
	"eleva/src/database/models"
	"eleva/src/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Client *gorm.DB

func Connect() {
	client, err := gorm.Open(postgres.Open(utils.GetEnv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic("could not connect to database")
	}

	client.AutoMigrate(&models.User{}, &models.VerificationCode{})

	Client = client
}
