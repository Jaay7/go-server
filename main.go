package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := "mongodb://localhost:27017/"
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI(uri))

	router := gin.New()
	router.Use(gin.Logger())

	router.POST("/createjob", func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var job Job

		if err := c.BindJSON(&job); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := client.Database("gomongo").Collection("job")

		result, err := collection.InsertOne(ctx, job)

		if err != nil {
			msg := fmt.Sprintf("Food item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	})

	router.Run(":" + port)
}
