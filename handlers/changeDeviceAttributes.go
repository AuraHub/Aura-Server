package handlers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func ChangeAttributes(attributesToSet models.DeviceAttributesToSet) error {
	SendAttributes(attributesToSet)

	filter := bson.D{{Key: "device_id", Value: attributesToSet.DeviceId}}

	changes := bson.M{}

	for _, change := range attributesToSet.Attributes {
		changes["attributes."+change.Name+".value"] = change.Value
		changes["attributes."+change.Name+".updated_at"] = time.Now()
	}

	update := bson.D{{Key: "$set", Value: changes}}

	_, err := initializers.Database.Collection("devices").
		UpdateOne(context.TODO(), filter, update)
	return err
}

func SendAttributes(attributesToSet models.DeviceAttributesToSet) {
	// Send data to device
	for _, attribute := range attributesToSet.Attributes {
		initializers.PahoConnection.Publish(
			attributesToSet.DeviceId+"|"+attribute.Name,
			0,
			false,
			attribute.Value,
		)
	}
}
