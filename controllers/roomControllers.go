package controllers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetRoom(c *gin.Context) {
	id := c.Param("id")

	var result *gorm.DB

	var room []models.Room

	if id == "" {
		result = initializers.DB.Preload("Devices.AttributeValues").Find(&room)
	} else {
		result = initializers.DB.Preload("Devices.AttributeValues").First(&room, "id = ?", id)
	}

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to fetch room",
		})

		return
	}

	if id == "" {
		c.JSON(http.StatusOK, room)
	} else {
		c.JSON(http.StatusOK, room[0])
	}
}

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

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created new room",
		"data":    room,
	})
}

func UpdateRoom(c *gin.Context) {
	// Get the vars off req body
	var body struct {
		ID   uuid.UUID
		Name string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}
	var room models.Room = models.Room{ID: body.ID}

	result := initializers.DB.Model(&room).
		Clauses(clause.Returning{}).
		Where("id = ?", body.ID).
		Updates(models.Room{Name: body.Name})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update room",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated the room",
		"data":    room,
	})
}

func DeleteRoom(c *gin.Context) {
	// Get the id off req body
	var body struct {
		ID string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	editResult := initializers.DB.Model(&models.Device{}).Where("room_id = ?", body.ID).
		Updates(map[string]interface{}{"Configured": false, "RoomID": nil})

	deleteResult := initializers.DB.Where("id = ?", body.ID).Delete(&models.Room{})

	if editResult.Error != nil || deleteResult.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete room",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted room",
	})
}
