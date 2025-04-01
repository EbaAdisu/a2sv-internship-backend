package main

import (
	"context"
	"log"
	"time"

	"task-manager/Delivery/controllers"
	"task-manager/Delivery/routers"
	"task-manager/Repositories"
	"task-manager/Usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // MongoDB connection.
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal(err)
    }
    db := client.Database("taskdb")

    // Repositories.
    taskRepo := Repositories.NewTaskRepository(db, "tasks")
    userRepo := Repositories.NewUserRepository(db, "users")

    // Usecases.
    taskUsecase := Usecases.NewTaskUsecase(taskRepo, 2*time.Second)
    userUsecase := Usecases.NewUserUsecase(userRepo, 2*time.Second)

    // Create the controller.
    ctrl := controllers.NewController(taskUsecase, userUsecase)

    // Setup and run the router.
    r := routers.SetupRouter(ctrl)
    r.Run(":8080")
}