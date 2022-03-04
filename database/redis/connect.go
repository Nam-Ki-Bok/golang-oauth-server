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
		log.Fatal("Redis connect error")
	}

	DB = db
	Nil = redis.Nil
}
