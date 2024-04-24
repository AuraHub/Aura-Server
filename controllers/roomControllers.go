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

func GetRoom(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		cursor, err := initializers.Database.Collection("rooms").Find(context.TODO(), bson.D{})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to read body",
			})
			return
		}

		var rooms []models.Room = []models.Room{}

		if err = cursor.All(context.TODO(), &rooms); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse data",
			})
		}
		for _, room := range rooms {
			res, _ := bson.MarshalExtJSON(room, false, false)
			fmt.Println(string(res))
		}
		c.JSON(http.StatusOK, rooms)

	} else {
		var room models.Room

		objectId, _ := primitive.ObjectIDFromHex(id)
		filter := bson.D{{Key: "_id", Value: objectId}}

		err := initializers.Database.Collection("rooms").FindOne(context.TODO(), filter).Decode(&room)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Room cannot be found",
			})
			return
		}

		c.JSON(http.StatusOK, room)

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
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, err := initializers.Database.Collection("rooms").InsertOne(context.TODO(), room)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create room",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully created new room",
		"data":    result,
	})
}

func UpdateRoom(c *gin.Context) {
	// Get the vars off req body
	var body struct {
		ID   string
		Name string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	id, _ := primitive.ObjectIDFromHex(body.ID)
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: body.Name}}}}

	_, err := initializers.Database.Collection("rooms").
		UpdateOne(context.TODO(), filter, update)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update room",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully updated the room",
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

	objectId, _ := primitive.ObjectIDFromHex(body.ID)
	filter := bson.D{{Key: "_id", Value: objectId}}

	_, err := initializers.Database.Collection("rooms").DeleteOne(context.TODO(), filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete room",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully deleted room",
	})
}
