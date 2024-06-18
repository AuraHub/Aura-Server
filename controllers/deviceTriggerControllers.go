package controllers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTrigger(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		cursor, err := initializers.Database.Collection("deviceTriggers").Find(context.TODO(), bson.D{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
		}

		var deviceTriggers []models.DeviceTrigger

		if err = cursor.All(context.TODO(), &deviceTriggers); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse data",
			})
			return
		}
		for _, room := range deviceTriggers {
			res, _ := bson.MarshalExtJSON(room, false, false)
			fmt.Println(string(res))
		}
		c.JSON(http.StatusOK, deviceTriggers)

	} else {
		var deviceTrigger models.DeviceTrigger

		objectId, _ := primitive.ObjectIDFromHex(id)
		filter := bson.D{{Key: "_id", Value: objectId}}

		err := initializers.Database.Collection("deviceTriggers").FindOne(context.TODO(), filter).Decode(&deviceTrigger)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Device trigger cannot be found",
			})
			return
		}

		c.JSON(http.StatusOK, deviceTrigger)
	}
}

func UpdateTrigger(c *gin.Context) {
	// Get the vars off req body
	var body struct {
		ID     string
		RoomID *primitive.ObjectID
		Name   string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	objectId, _ := primitive.ObjectIDFromHex(body.ID)
	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: body.Name},
				{Key: "room_id", Value: body.RoomID},
				{Key: "updated_at", Value: time.Now()},
			},
		},
	}

	_, err := initializers.Database.Collection("deviceTriggers").
		UpdateOne(context.TODO(), filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update device trigger",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated the device trigger",
	})
}

func DeleteTrigger(c *gin.Context) {
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

	objectId, _ := primitive.ObjectIDFromHex(body.ID)
	filter := bson.D{{Key: "_id", Value: objectId}}

	_, err := initializers.Database.Collection("deviceTriggers").DeleteOne(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete device trigger",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted device trigger"})
}

func ConfigureTrigger(c *gin.Context) {
	// Get the vars off req body
	var body struct {
		ID     string
		RoomID *primitive.ObjectID
		Name   string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	objectId, _ := primitive.ObjectIDFromHex(body.ID)
	filter := bson.D{{Key: "_id", Value: objectId}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "name", Value: body.Name},
				{Key: "room_id", Value: body.RoomID},
				{Key: "configured", Value: true},
				{Key: "configured_at", Value: time.Now()},
			},
		},
	}

	_, err := initializers.Database.Collection("deviceTriggers").
		UpdateOne(context.TODO(), filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to setup device trigger",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully setup device trigger",
	})
}
