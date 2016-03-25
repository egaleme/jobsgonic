package main

import (
	"os"

	"github.com/egaleme/jobsgonic/app/common"
	"github.com/egaleme/jobsgonic/app/cors"
	"github.com/egaleme/jobsgonic/app/router"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	gin.SetMode(gin.ReleaseMode)
	common.Startup()
	app := gin.Default()
	app.Use(cors.Cors(cors.Options{}))
	router.InitRoutes(app)
	app.Static("/assets", "./public/assets")
	app.StaticFile("/", "./public")
	app.Run(":" + port)

}
