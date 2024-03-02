package handlers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type deviceSetup struct {
	DeviceId   string   `json:"deviceId"`
	Attributes []string `json:"attributes"`
}

func SetupDevice(c mqtt.Client, m mqtt.Message) {
	var setupData deviceSetup

	err := json.Unmarshal(m.Payload(), &setupData)
	if err != nil {
		panic(err)

	}

	// Check if device exists in database
	var device models.Device
	initializers.DB.First(&device, "device_id = ?", setupData.DeviceId)

	if device.DeviceId != "" {
		// Update online state
		initializers.DB.Model(&device).
			Updates(models.Device{Online: true, LastOnline: time.Now()})

	} else {
		// Create new record in database

		var attributeValues []models.AttributeValue

		for _, newAttributeName := range setupData.Attributes {
			attributeValues = append(attributeValues, models.AttributeValue{DeviceID: device.ID, AttributeName: newAttributeName})
		}

		newDevice := models.Device{DeviceId: setupData.DeviceId, AttributeValues: attributeValues}

		resultDevices := initializers.DB.Create(&newDevice)

		if resultDevices.Error != nil {
			panic(resultDevices.Error)
		}

		fmt.Println("Device id:", setupData.DeviceId)
		fmt.Println("Attributes:", string(strings.Join(setupData.Attributes, ", ")))
	}
}
