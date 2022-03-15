package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"infradev-practice/Wade/OAuth2.0-server/database/mongo"
	"infradev-practice/Wade/OAuth2.0-server/server/api"
	"time"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func Setup() *gin.Engine {
	r := gin.New()

	r.Use(gin.LoggerWithFormatter(customLogFormatter))
	r.Use(gin.CustomRecovery(customRecover))

	api.InitAuth(r)

	api.InitPersonal(r)
	api.InitService(r)
	api.InitCost(r)
	api.InitStock(r)
	api.InitStatsPrimeCost(r)

	return r
}

func customLogFormatter(param gin.LogFormatterParams) string {

	today := time.Now().Format("2006-01-02")
	conn := mongo.RequestLog.Database("request_log").Collection(today)

	log := bson.D{
		{Key: "clientIP", Value: param.ClientIP},
		{Key: "timeStamp", Value: param.TimeStamp.Format("2006-01-02 15:04:05")},
		{Key: "statusCode", Value: param.StatusCode},
		{Key: "method", Value: param.Method},
		{Key: "path", Value: param.Path},
		{Key: "auth", Value: param.Request.Header.Get("Authorization")},
		{Key: "userAgent", Value: param.Request.UserAgent()},
		{Key: "error", Value: param.ErrorMessage},
	}
	result, err := conn.InsertOne(context.TODO(), log)
	if err != nil {
		panic(err)
	}

	objectID := result.InsertedID.(primitive.ObjectID).String() + "\n"
	return objectID
}

func customRecover(c *gin.Context, recovered interface{}) {
	r := recovered.(map[string]interface{})
	code := r["code"].(int)
	err := r["err"].(error)

	c.Error(err)
	c.AbortWithStatusJSON(code, gin.H{
		"status": "false",
		"error":  err.Error(),
	})
}
