package main

import (
	"Aura-Server/controllers"
	"Aura-Server/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	r := gin.Default()

	r.GET("/ping", controllers.Ping)
	
	r.Run(":3000")
}