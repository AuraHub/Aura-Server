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
	// Define AttributesTypes
	AttributesTypes := map[string]string{"OnOff": "bool", "Brightness": "value"}

	// Convert data to JSON
	var setupData deviceSetup

	err := json.Unmarshal(m.Payload(), &setupData)
	if err != nil {
		panic(err)
	}

	// Check if device exists in database
	var device models.Device

	filter := bson.D{{Key: "device_id", Value: setupData.DeviceId}}
	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "online", Value: true},
				{Key: "last_online", Value: time.Now()},
			},
		},
	}
	noDeviceInDB := initializers.Database.Collection("devices").FindOneAndUpdate(context.TODO(), filter, update).Decode(&device)

	// Check if exists
	if noDeviceInDB == nil {
		// If exists in database -> update online state
		var output []models.AttributeToSet
		for name, attribute := range device.Attributes {
			output = append(output, models.AttributeToSet{
				Name:  name,
				Value: attribute.Value,
			})
		}

		attributes := models.DeviceAttributesToSet{
			DeviceId:   device.DeviceId,
			Attributes: output,
		}
		SendAttributes(attributes)

	} else {
		// If not exists -> Define new device
		newDevice := models.Device{
			DeviceId: setupData.DeviceId, RoomID: nil, Online: true, Configured: false, LastOnline: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Attributes: make(map[string]models.Attribute),
		}

		// Create list of attributes to connect
		for _, newAttributeName := range setupData.Attributes {
			newDevice.Attributes[newAttributeName] = models.Attribute{
				AttributeType: AttributesTypes[newAttributeName],
			}
		}

		// Create new record in database
		_, err := initializers.Database.Collection("devices").InsertOne(context.TODO(), newDevice)

		if err != nil {
			panic(err)
		}
	}
}
