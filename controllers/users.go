package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID string `json: id`
	PW string `json: pw`
}

func Register(c *gin.Context) {
	user := new(User)

	if err := c.ShouldBind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, user)
}

// func AddUser(c *gin.Context) {
// 	email := c.Param("email")

// 	if utils.IsEmailValid(email) == true {
// 		c.JSON(http.StatusOK, "Done")
// 	}
// }
