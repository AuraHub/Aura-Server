package main

import (
	"Aura-Server/controllers"
	"Aura-Server/handlers"
	"Aura-Server/initializers"
	"Aura-Server/middleware"
	"time"

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

	// Room routes
	r.POST("/room", middleware.RequireAuth, controllers.NewRoom)

	// Device routes
	r.POST("/device", middleware.RequireAuth, controllers.ConfigureDevice)
	r.DELETE("/device", middleware.RequireAuth, controllers.DeleteDevice)

	// Subscribes to MQTT topics
	initializers.PahoConnection.Subscribe("setup", 0, handlers.SetupDevice)
	initializers.PahoConnection.Subscribe("returnPing", 0, handlers.ReturnedPing)

	ticker := time.NewTicker(25 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

				handlers.PingDevices()

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	r.Run(":3000")
}
