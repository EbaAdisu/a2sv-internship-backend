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

	// Get collections from the database
	taskCollection := client.Database("taskdb").Collection("tasks")
	userCollection := client.Database("taskdb").Collection("users")

	// Initialize services
	taskService := data.NewTaskService(taskCollection)
	userService := data.NewUserService(userCollection)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskService)
	userController := controllers.NewUserController(userService)

	// Setup and run Gin router
	r := router.SetupRouter(taskController, userController)
	r.Run(":8080")
}
