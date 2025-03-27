package main

import (
	"context"
	"log"
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Configure MongoDB connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verify connection
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	// Get tasks collection from the database
	collection := client.Database("taskdb").Collection("tasks")

	// Initialize TaskService with MongoDB collection
	taskService := data.NewTaskService(collection)

	// Initialize TaskController
	taskController := controllers.NewTaskController(taskService)

	// Setup and run Gin router
	r := router.SetupRouter(taskController)
	r.Run(":8080")
}
