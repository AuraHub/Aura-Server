package controllers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteDevice(c *gin.Context) {
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

	result := initializers.DB.Where("id = ?", body.ID).Delete(&models.Device{})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete device",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted device"})
}
