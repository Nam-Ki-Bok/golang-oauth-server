package main

import (
	"github.com/gin-gonic/gin"
	"infradev-practice/Wade/OAuth2.0-server/token"

	"infradev-practice/Wade/OAuth2.0-server/database/maria"
	"infradev-practice/Wade/OAuth2.0-server/database/redis"
	"log"
)

func init() {
	maria.Connect()
	redis.Connect()
}

func main() {
	r := gin.Default()

	r.GET("/request/token", token.Request)
	r.POST("/generate/token", token.Generate)

	log.Fatal(r.Run(":9096"))
}
