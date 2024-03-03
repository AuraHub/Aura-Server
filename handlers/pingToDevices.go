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

func PingDevices() {
	var devices []models.Device

	// Send ping to devices
	initializers.PahoConnection.Publish("ping", 0, false, "")

	// Wait 3 seconds
	time.Sleep(3 * time.Second)

	// Calculate the timestamp 3 seconds ago
	fiveSecondsAgo := time.Now().Add(-3 * time.Second)

	// Remove devices which didn't respond
	initializers.DB.Model(&devices).Where("online = ? AND last_online <= ?", true, fiveSecondsAgo).
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
	initializers.DB.Model(&device).Where("device_id = ?", pingData.DeviceId).
		Updates(models.Device{Online: true, LastOnline: time.Now()})
}
