package main

import (
	"gomovie2/config"
	"gomovie2/models"
)

func main() {
	config.Init()

	user := models.Users{UserID: "user1", Email: "poyo062@naver.com", Password: "test", Name: "test"}
	config.DB.Create(&user)

}
