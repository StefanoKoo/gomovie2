package main

import (
	database "gomovie2/database"
	"gomovie2/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	// user := models.Users{UserID: "user1", Email: "poyo062@naver.com", Password: "test", Name: "test"}
	// config.DB.Create(&user)
	r := gin.Default()
	routes.Setup(r)

	r.Run()
}
