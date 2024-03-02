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
	result := initializers.DB.Find(&devices)

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
	initializers.DB.First(&device, "device_id = ?", pingData.DeviceId).
		Updates(models.Device{Online: true, LastOnline: time.Now()})

	if len(devices) != 0 {
		// Remove device from array
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
