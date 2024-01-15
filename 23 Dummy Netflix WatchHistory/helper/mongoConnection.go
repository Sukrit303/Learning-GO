package helper

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const url = ""
const dbName = "netflix"
const colName = "watchHistory"

var collection *mongo.Collection

func init() {

	// Client option to setup connection info
	clientOptions := options.Client().ApplyURI(url)

	// Connection to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	// Creating Collection
	collection = client.Database(dbName).Collection(colName)

	// Collection instance created
	fmt.Println("MongoDB collection created")
}
