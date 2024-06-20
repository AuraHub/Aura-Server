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

	updates := bson.M{}
	for _, trigger := range body.Triggers {
		var actionsToAdd []interface{}
		for _, action := range trigger.Actions {
			actionToAdd := bson.M{
				"device_id": action.DeviceId,
				"action":    action.Action,
				"attribute": action.Attribute,
			}
			if action.Value != "" {
				actionToAdd["value"] = action.Value
			}
			actionsToAdd = append(actionsToAdd, actionToAdd)
		}
		// Use $push with $each to add new actions to the existing array
		updates["triggers."+trigger.Trigger+".actions"] = actionsToAdd
	}

	update := bson.D{{Key: "$set", Value: updates}}

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
