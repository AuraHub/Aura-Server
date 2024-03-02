package controllers

import (
	"Aura-Server/handlers"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	handlers.PingDevices()

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
