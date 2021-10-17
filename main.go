package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	uri := os.Getenv("MONGODB_URI")
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
			msg := fmt.Sprintf("Job was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		defer cancel()

		c.JSON(http.StatusOK, result)
	})

	router.GET("/jobs", func(c *gin.Context) {
		jobs := []Job{}

		collection := client.Database("gomongo").Collection("job")

		cursor, err := collection.Find(context.TODO(), bson.M{})

		if err != nil {
			msg := fmt.Sprintf("no jobs found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		for cursor.Next(context.TODO()) {
			var job Job
			cursor.Decode(&job)
			jobs = append(jobs, job)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": jobs,
		})
		return
	})

	router.GET("/job/:jobId", func(c *gin.Context) {
		jobId := c.Param("jobId")

		job := Job{}
		collection := client.Database("gomongo").Collection("job")

		objId, _ := primitive.ObjectIDFromHex(jobId)
		err := collection.FindOne(context.TODO(), bson.M{"_id": objId}).Decode(&job)

		if err != nil {
			msg := fmt.Sprintf("current job not found")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": job,
		})
		return
	})

	router.GET("/jobs/:jobtype", func(c *gin.Context) {
		jobs := []Job{}

		jobtype := c.Param("jobtype")
		collection := client.Database("gomongo").Collection("job")

		// filter := bson.D{{"jobtype", jobtype}}
		cursor, err := collection.Find(context.TODO(), bson.D{{"jobtype", jobtype}})

		if err != nil {
			msg := fmt.Sprintf("no jobs found under this section")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		for cursor.Next(context.TODO()) {
			var job Job
			cursor.Decode(&job)
			jobs = append(jobs, job)
		}

		c.JSON(http.StatusOK, gin.H{
			"data": jobs,
		})

	})

	router.Run(":" + port)
}
