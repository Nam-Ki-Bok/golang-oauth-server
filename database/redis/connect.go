package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

var DB *redis.Client
var Nil interface{}
var ctx = context.Background()

func Connect() {

	db := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_CONNECTION_INFO"),
	})

	if db.Ping(ctx).String() != "ping: PONG" {
		log.Fatal("Redis connection failed")
	}

	DB = db
	Nil = redis.Nil
	log.Println("Redis connected successfully")
}
