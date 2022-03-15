package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	k "gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"infradev-practice/Wade/OAuth2.0-server/kafka"
	"infradev-practice/Wade/OAuth2.0-server/utils"
	"net/http"
)

func Publish(c *gin.Context) {
	c.Next()

	msg, ok := c.Get("msg")
	if !ok {
		utils.ReturnError(http.StatusBadRequest, errors.New("message dose not exist"))
	}

	// Produce messages to topic (asynchronously)
	topic := "test-infra"
	err := kafka.Prod.Produce(&k.Message{
		TopicPartition: k.TopicPartition{Topic: &topic, Partition: k.PartitionAny},
		Value:          []byte(msg.(string)),
	}, nil)
	if err != nil {
		utils.ReturnError(http.StatusBadRequest, err)
	}

	// execute kafka publish
	fmt.Printf("kafka publish : %s\n", msg)

	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
