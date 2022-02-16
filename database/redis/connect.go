package redis

import "github.com/go-redis/redis/v7"

var DB *redis.Client

func Connect() {
	db := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	DB = db
}