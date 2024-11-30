package main

import (
	"sistem-manajemen-sekolah/initializers"
	"sistem-manajemen-sekolah/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
	initializers.SyncDatabase()
}

func main() {
	app := gin.Default()

	routes.InitRoute(app)
	app.Run()
}
