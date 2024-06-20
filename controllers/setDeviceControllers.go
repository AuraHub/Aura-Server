package controllers

import (
	"Aura-Server/handlers"
	"Aura-Server/initializers"
	"Aura-Server/models"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func SetDevice(c *gin.Context) {
	// Get the vars off req body
	var body models.DeviceAttributesToSet

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	err := handlers.ChangeAttributes(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to set device",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully set device",
	})
}

func SetDeviceTrigger(c *gin.Context) {
	// Get the vars off req body
	var body models.DeviceTriggersToSet

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	filter := bson.D{{Key: "device_id", Value: body.DeviceId}}

	changes := bson.M{}

	for _, trigger := range body.Triggers {
		var actionsToAdd []models.Action
		for _, change := range trigger.Actions {
			actionsToAdd = append(actionsToAdd, models.Action{
				DeviceId: change.DeviceId,
				Action:   change.Action,
				Value:    change.Value,
			})
		}
		changes["triggers."+trigger.Trigger+".actions"] = actionsToAdd
	}

	update := bson.D{{Key: "$set", Value: changes}}

	_, err := initializers.Database.Collection("deviceTriggers").
		UpdateOne(context.TODO(), filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update device trigger",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully set device trigger",
	})
}
