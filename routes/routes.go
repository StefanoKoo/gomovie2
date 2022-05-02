package routes

import (
	"gomovie2/controllers"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.POST("/api/v1/register", controllers.Register)
}
