package main

import (
	"Aura-Server/controllers"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()

	r.GET("/ping", controllers.Ping)
	
	r.Run(":3000")
}