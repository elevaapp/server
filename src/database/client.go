package database

import (
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

	Client = client
}
