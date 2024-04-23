package handlers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"context"
	"encoding/json"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
)

type deviceSetup struct {
	DeviceId   string   `json:"deviceId"`
	Attributes []string `json:"attributes"`
}

func SetupDevice(c mqtt.Client, m mqtt.Message) {

	// Convert data to JSON
	var setupData deviceSetup

	err := json.Unmarshal(m.Payload(), &setupData)
	if err != nil {
		panic(err)
	}

	// Check if device exists in database
	var device models.Device

	filter := bson.D{{Key: "device_id", Value: setupData.DeviceId}}
	initializers.Database.Collection("devices").FindOne(context.TODO(), filter).Decode(&device)

	if device.DeviceId != "" {
		// Update online state
		update := bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{Key: "online", Value: true},
					{Key: "last_online", Value: time.Now()},
				},
			},
		}

		initializers.Database.Collection("devices").UpdateOne(context.TODO(), filter, update)

	} else {
		// Define new device
		newDevice := models.Device{
			DeviceId: setupData.DeviceId, RoomID: nil, Online: true, LastOnline: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Attributes: make(map[string]models.Attribute),
		}

		// Create list of attributes to connect
		for _, newAttributeName := range setupData.Attributes {
			newDevice.Attributes[newAttributeName] = models.Attribute{
				UpdatedAt: time.Now(),
			}
		}

		// Create new record in database
		_, err := initializers.Database.Collection("devices").InsertOne(context.TODO(), newDevice)

		if err != nil {
			panic(err)
		}
	}
}
