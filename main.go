package main

import (
	"context"
	"github.com/joho/godotenv"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/database/mongo"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
	"infradev-practice/Wade/OAuth2.0-server/kafka"
	"infradev-practice/Wade/OAuth2.0-server/server"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	maria.Connect()
	redis.Connect()
	mongo.Connect()
	kafka.Connect()
}

var ctx = context.Background()

func main() {
	defer kafka.Prod.Close()

	r := server.Setup()

	log.Fatal(r.Run(":1054"))
}
