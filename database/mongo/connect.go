package mongo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var RequestLog *mongo.Client

func Connect() {
	credential := options.Credential{
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PW"),
	}

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI")).SetAuth(credential)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	RequestLog = client
}
