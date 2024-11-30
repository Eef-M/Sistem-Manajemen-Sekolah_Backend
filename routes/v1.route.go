package routes

import (
	"sistem-manajemen-sekolah/controllers"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {
	route := app

	route.GET("/testing", controllers.Testing)
}
