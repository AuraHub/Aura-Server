package initializers

import (
	"context"
	"fmt"

	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
)

func ConnectToDB() {
	dsn := os.Getenv("DB_CREDENTIALS")
	dbName := os.Getenv("DB_NAME")

	opts := options.Client().
		ApplyURI(dsn)

	var err error
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	Database = Client.Database(dbName)

	var result bson.M
	if err := Client.Database(dbName).RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
