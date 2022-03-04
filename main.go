package main

import (
	"context"
	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/database/mongo"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
	"infradev-practice/Wade/OAuth2.0-server/server"
	"log"
)

func init() {
	maria.Connect()
	redis.Connect()
	mongo.Connect()
}

var ctx = context.Background()

func main() {
	r := server.Setup()

	redis.DB.FlushAll(ctx)

	log.Fatal(r.Run(":1054"))
}
