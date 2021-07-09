package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/eranamarante/go-react-expense-tracker/server/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBInstance() *mongo.Client {
	config := helper.GetConfiguration()

	client, err := mongo.NewClient(options.Client().ApplyURI(config.ConnectionString))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	return client
}

var Client *mongo.Client = DBInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	config := helper.GetConfiguration()

	var collection *mongo.Collection = client.Database(config.DatabaseName).Collection(collectionName)

	return collection
}
