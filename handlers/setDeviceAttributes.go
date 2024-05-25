package handlers

import (
	"Aura-Server/initializers"
	"Aura-Server/models"
	"context"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetAttributes(attributesToSet models.DeviceAttributesToSet) error {
	objectId, _ := primitive.ObjectIDFromHex(attributesToSet.ID)
	filter := bson.D{{Key: "_id", Value: objectId}}

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
