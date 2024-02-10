package main

import (
	"Aura-Server/controllers"
	"Aura-Server/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	// Uncoment when there is need to sync database
	// initializers.SyncDatabase()
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.GET("/ping", controllers.Ping)
	r.GET("/syncdatabase", controllers.SyncDatabase)
	r.POST("/signup", controllers.Signup)

	r.Run(":3000")
}
