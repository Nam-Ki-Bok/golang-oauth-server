package redis

import "github.com/go-redis/redis/v7"

var DB *redis.Client
var NilErr interface{}

func Connect() {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	DB = db
	NilErr = redis.Nil
}
