package controllers

import (
	"Aura-Server/handlers"
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

func GetDevice(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		cursor, err := initializers.Database.Collection("devices").Find(context.TODO(), bson.D{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
		}

		var devices []models.Device = []models.Device{}

		if err = cursor.All(context.TODO(), &devices); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse data",
			})
			return
		}
		for _, room := range devices {
			res, _ := bson.MarshalExtJSON(room, false, false)
			fmt.Println(string(res))
		}
		c.JSON(http.StatusOK, devices)

	} else {
		var device models.Device

		objectId, _ := primitive.ObjectIDFromHex(id)
		filter := bson.D{{Key: "_id", Value: objectId}}

		err := initializers.Database.Collection("devices").FindOne(context.TODO(), filter).Decode(&device)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Device cannot be found",
			})
			return
		}

		c.JSON(http.StatusOK, device)
	}
}

func UpdateDevice(c *gin.Context) {
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

	_, err := initializers.Database.Collection("devices").
		UpdateOne(context.TODO(), filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update device",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated the device",
	})
}

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

	objectId, _ := primitive.ObjectIDFromHex(body.ID)
	filter := bson.D{{Key: "_id", Value: objectId}}

	_, err := initializers.Database.Collection("devices").DeleteOne(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete device",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted device"})
}

func ConfigureDevice(c *gin.Context) {
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

	_, err := initializers.Database.Collection("devices").
		UpdateOne(context.TODO(), filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to setup device",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully setup device",
	})
}

func SetDevice(c *gin.Context) {
	// Get the vars off req body
	var body models.DeviceAttributesToSet

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	err := handlers.SetAttributes(body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update device",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully set device",
	})
}
