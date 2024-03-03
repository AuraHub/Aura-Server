package initializers

import (
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var PahoConnection mqtt.Client

func ConnectPaho() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(os.Getenv("BROKER_IP"))
	opts.SetClientID("AuraHub")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	PahoConnection = mqtt.NewClient(opts)
	if token := PahoConnection.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
