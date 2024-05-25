package handlers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"context"
	"encoding/json"
	"fmt"

	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func SetAttributes(attributesToSet models.DeviceAttributesToSet) error {
	// Send data to device
	jsonData, _ := json.Marshal(attributesToSet.Attributes)

	initializers.PahoConnection.Publish(
		attributesToSet.DeviceId,
		0,
		false,
		jsonData,
	)

	fmt.Println(attributesToSet.DeviceId)
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
