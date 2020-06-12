package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func initAPI(ctx context.Context, mongoClient *mongo.Client) *gin.Engine {
	router := gin.Default()

	router.GET("/api", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Welcome to a Go API",
		})
	})

	router.POST("/api/name", func(context *gin.Context) {
		db := mongoClient.Database(os.Getenv("HEROKU_DB"))
		collection := db.Collection("golang")

		var person Person
		context.BindJSON(&person)

		result, err := collection.InsertOne(ctx, bson.D{
			{Key: "name", Value: person.Name},
		})

		if err != nil {
			log.Fatal(err)
		}

		context.JSON(http.StatusOK, gin.H{
			"result": result,
		})
	})

	return router
}
