package handlers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"encoding/json"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type devicePing struct {
	DeviceId string `json:"deviceId"`
}

var devices []models.Device

func PingDevices() {
	result := initializers.DB.Find(&devices).Where("online = ?", true)

	if len(devices) != 0 {
		if result.Error != nil {
			panic(result.Error)
		}

		// Send ping to devices
		initializers.PahoConnection.Publish("ping", 0, false, "")

		// Wait 5 seconds
		time.Sleep(5 * time.Second)

		// Remove devices which didn't respond
		initializers.DB.Model(&devices).
			Updates(map[string]interface{}{"Online": false})

		devices = nil
	}
}

func ReturnedPing(c mqtt.Client, m mqtt.Message) {
	// Convert data to JSON
	var pingData devicePing

	err := json.Unmarshal(m.Payload(), &pingData)
	if err != nil {
		panic(err)
	}

	// Update status to online
	var device models.Device
	initializers.DB.Model(&device).Where("device_id = ?", pingData.DeviceId).
		Updates(models.Device{Online: true, LastOnline: time.Now()})

	// Remove device from array
	if len(devices) != 0 {
		var index int

		for n, x := range devices {
			if x.DeviceId == device.DeviceId {
				index = n
				break
			}
		}

		devices = append(devices[:index], devices[index+1:]...)
	}
}
