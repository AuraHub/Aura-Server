package controllers

import (
	"Aura-Server/initializers"

	"github.com/gin-gonic/gin"
)

func SyncDatabase(c *gin.Context){
	initializers.SyncDatabase()

	c.AbortWithStatus(200)
}