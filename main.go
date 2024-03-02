package main

import (
	"Aura-Server/controllers"
	"Aura-Server/handlers"
	"Aura-Server/initializers"
	"Aura-Server/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	initializers.ConnectPaho()

	// Uncomment when there is need to sync database
	// initializers.SyncDatabase()
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.GET("/ping", controllers.Ping)
	r.GET("/syncdatabase", controllers.SyncDatabase)

	// User routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/logout", controllers.Logout)

	// Subscribes to MQTT topics
	initializers.PahoConnection.Subscribe("setup", 0, handlers.SetupDevice)
	initializers.PahoConnection.Subscribe("returnPing", 0, handlers.ReturnedPing)

	r.Run(":3000")
}
