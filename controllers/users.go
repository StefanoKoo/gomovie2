package controllers

import (
	"gomovie2/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	email := c.Param("email")

	if utils.IsEmailValid(email) == true {
		c.JSON(http.StatusOK, "Done")
	}
}
