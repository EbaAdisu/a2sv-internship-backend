package main

import (
	"task_manager/controllers"
	"task_manager/data"
	"task_manager/router"
)

func main() {
    // Initialize task service
    taskService := data.NewTaskService()

    // Initialize task controller with the task service
    taskController := controllers.NewTaskController(taskService)
    
    // Pass task controller to router setup
    r := router.SetupRouter(taskController)
    r.Run(":8080")
}