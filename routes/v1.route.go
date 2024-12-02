package routes

import (
	"sistem-manajemen-sekolah/controllers"
	"sistem-manajemen-sekolah/controllers/auth"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.GET("/testing", controllers.Testing)
	route.POST("/login", auth.Login)
}
