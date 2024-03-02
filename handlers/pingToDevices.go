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

}

func ReturnedPing(c mqtt.Client, m mqtt.Message) {
	// Convert data to JSON
	var pingData devicePing

	err := json.Unmarshal(m.Payload(), &pingData)
	if err != nil {
		panic(err)
	}

	var device models.Device
	initializers.DB.First(&device, "device_id = ?", pingData.DeviceId).
		Updates(models.Device{Online: true, LastOnline: time.Now()})
}
