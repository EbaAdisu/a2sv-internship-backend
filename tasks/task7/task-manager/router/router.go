package router

import (
	"task_manager/controllers"
	"task_manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskCtrl *controllers.TaskController, userCtrl *controllers.UserController) *gin.Engine {
	r := gin.Default()

	// Public routes for user registration and login
	r.POST("/register", userCtrl.Register)
	r.POST("/login", userCtrl.Login)

	// Public task routes: anyone with valid token can view tasks.
	r.GET("/tasks", middleware.AuthMiddleware(), taskCtrl.GetAllTasks)
	r.GET("/tasks/:id", middleware.AuthMiddleware(), taskCtrl.GetTaskByID)

	// Protected routes: only admins
	adminRoutes := r.Group("/", middleware.AuthMiddleware(), middleware.AdminMiddleware())
	adminRoutes.POST("/tasks", taskCtrl.CreateTask)
	adminRoutes.PUT("/tasks/:id", taskCtrl.UpdateTask)
	adminRoutes.DELETE("/tasks/:id", taskCtrl.DeleteTask)

	return r
}
