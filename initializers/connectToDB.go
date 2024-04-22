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
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	dsn := os.Getenv("DB_CREDENTIALS")

	opts := options.Client().
		ApplyURI(dsn).
		SetServerAPIOptions(serverAPI)

	var err error
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	Database = Client.Database("AVA-DB")

	var result bson.M
	if err := Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
