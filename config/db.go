package config

import (
	"gomovie2/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error

	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect DB")
	}

	DB.AutoMigrate(&models.Movies{}, &models.Users{})
}
