package controllers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRoom(c *gin.Context) {
	// Get the credentials off req body
	var body struct {
		Name string
	}

	// Get user from middleware
	user, _ := c.Get("user")

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// Craete new room
	room := models.Room{
		Name:      body.Name,
		CreatedBy: user.(models.User).ID,
	}
	result := initializers.DB.Create(&room)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create room",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
