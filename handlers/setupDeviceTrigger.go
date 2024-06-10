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

type deviceTriggerSetup struct {
	DeviceTriggerId string   `json:"deviceTriggerId"`
	Triggers        []string `json:"triggers"`
}

func SetupDeviceTrigger(c mqtt.Client, m mqtt.Message) {
	// Convert data to JSON
	var setupData deviceTriggerSetup

	err := json.Unmarshal(m.Payload(), &setupData)
	if err != nil {
		panic(err)
	}

	// Check if device exists in database
	var deviceTrigger models.DeviceTrigger

	filter := bson.D{{Key: "device_trigger_id", Value: setupData.DeviceTriggerId}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "online", Value: true},
				{Key: "last_online", Value: time.Now()},
			},
		},
	}

	noDeviceTriggerInDB := initializers.Database.Collection("deviceTriggers").FindOneAndUpdate(context.TODO(), filter, update).Decode(&deviceTrigger)

	// Check if exists
	if noDeviceTriggerInDB != nil {
		// If not exists -> Define new device
		newDeviceTrigger := models.DeviceTrigger{
			DeviceTriggerId: setupData.DeviceTriggerId, RoomID: nil, Online: true, Configured: false, LastOnline: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Triggers: make(map[string]models.Trigger),
		}

		// Create list of attributes to connect
		for _, newAttributeName := range setupData.Triggers {
			newDeviceTrigger.Triggers[newAttributeName] = models.Trigger{
				Actions: []models.Action{},
			}
		}

		// Create new record in database
		_, err := initializers.Database.Collection("deviceTriggers").InsertOne(context.TODO(), newDeviceTrigger)

		if err != nil {
			panic(err)
		}
	}
}
