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
}

func main() {
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	r.GET("/ping", controllers.Ping)

	// User routes
	r.GET("/user", middleware.RequireAuth, controllers.GetUser)
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/logout", controllers.Logout)

	// Room routes
	r.GET("/room", middleware.RequireAuth, controllers.GetRoom)
	r.GET("/room/:id", middleware.RequireAuth, controllers.GetRoom)
	r.POST("/room", middleware.RequireAuth, controllers.NewRoom)
	r.PUT("/room", middleware.RequireAuth, controllers.UpdateRoom)
	r.DELETE("/room", middleware.RequireAuth, controllers.DeleteRoom)

	// Device routes
	r.POST("/set/device", middleware.RequireAuth, controllers.SetDevice)
	r.POST("/set/trigger", middleware.RequireAuth, controllers.SetDeviceTrigger)

	r.GET("/device", middleware.RequireAuth, controllers.GetDevice)
	r.GET("/device/:id", middleware.RequireAuth, controllers.GetDevice)
	r.POST("/device", middleware.RequireAuth, controllers.ConfigureDevice)
	r.PUT("/device", middleware.RequireAuth, controllers.UpdateDevice)
	r.DELETE("/device", middleware.RequireAuth, controllers.DeleteDevice)

	// Device trigger routes
	r.GET("/trigger", middleware.RequireAuth, controllers.GetTrigger)
	r.GET("/trigger/:id", middleware.RequireAuth, controllers.GetTrigger)
	r.POST("/trigger", middleware.RequireAuth, controllers.ConfigureTrigger)
	r.PUT("/trigger", middleware.RequireAuth, controllers.UpdateTrigger)
	r.DELETE("/trigger", middleware.RequireAuth, controllers.DeleteTrigger)

	// Subscribes to MQTT topics
	initializers.PahoConnection.Subscribe("setupDevice", 0, handlers.SetupDevice)
	initializers.PahoConnection.Subscribe("setupDeviceTrigger", 0, handlers.SetupDeviceTrigger)
	initializers.PahoConnection.Subscribe("returnPing", 0, handlers.ReturnedPing)

	ticker := time.NewTicker(25 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:

				handlers.Ping()

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
