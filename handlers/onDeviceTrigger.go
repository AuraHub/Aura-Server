package handlers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"context"
	"encoding/json"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.mongodb.org/mongo-driver/bson"
)

type triggerSent struct {
	DeviceId string `json:"deviceId"`
	Trigger  string `json:"trigger"`
}

func OnDeviceTrigger(c mqtt.Client, m mqtt.Message) {
	// Convert data to JSON
	var triggerData triggerSent
	err := json.Unmarshal(m.Payload(), &triggerData)
	if err != nil {
		log.Printf("Error unmarshalling message: %v", err)
		return
	}

	// Retrieve the device trigger from the database
	filter := bson.D{{Key: "device_id", Value: triggerData.DeviceId}}
	var deviceTrigger models.DeviceTrigger
	err = initializers.Database.Collection("deviceTriggers").FindOne(context.TODO(), filter).Decode(&deviceTrigger)
	if err != nil {
		log.Printf("Error finding device trigger: %v", err)
		return
	}

	// Execute actions associated with the trigger
	actions := deviceTrigger.Triggers[triggerData.Trigger].Actions
	for _, action := range actions {
		filter := bson.D{{Key: "device_id", Value: action.DeviceId}}
		switch action.Action {
		case "set":
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "attributes." + action.Attribute + ".value", Value: action.Value},
					{Key: "attributes." + action.Attribute + ".updated_at", Value: time.Now()},
				}},
			}
			_, err := initializers.Database.Collection("devices").UpdateOne(context.TODO(), filter, update)
			if err != nil {
				log.Printf("Error updating device: %v", err)
				continue
			}
			initializers.PahoConnection.Publish(action.DeviceId+"|"+action.Attribute, 1, true, action.Value)
		case "switch":
			var currentDevice models.Device
			err = initializers.Database.Collection("devices").FindOne(context.TODO(), filter).Decode(&currentDevice)
			if err != nil {
				log.Printf("Error retrieving device for switch action: %v", err)
				continue
			}

			newValue := "true"
			if currentDevice.Attributes[action.Attribute].Value == "true" {
				newValue = "false"
			}
			update := bson.D{
				{Key: "$set", Value: bson.D{
					{Key: "attributes." + action.Attribute + ".value", Value: newValue},
					{Key: "attributes." + action.Attribute + ".updated_at", Value: time.Now()},
				}},
			}
			_, err = initializers.Database.Collection("devices").UpdateOne(context.TODO(), filter, update)
			if err != nil {
				log.Printf("Error updating device switch action: %v", err)
				continue
			}
			initializers.PahoConnection.Publish(action.DeviceId+"|"+action.Attribute, 1, true, newValue)
		}
	}
}
