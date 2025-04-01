package routers

import (
	"task-manager/Delivery/controllers"
	"task-manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the routes and middleware.
func SetupRouter(ctrl *controllers.Controller) *gin.Engine {
	r := gin.Default()

	// Public routes for auth (registration and login)
	r.POST("/register", ctrl.Register)
	r.POST("/login", ctrl.Login)

	// Group for protected routes using JWT middleware from Infrastructure
	protected := r.Group("/")
	protected.Use(Infrastructure.AuthMiddleware())
	{
		protected.GET("/tasks", ctrl.GetAllTasks)

		// Get a single task by ID.
		protected.GET("/tasks/:id", ctrl.GetTaskByID)

		// Admin-only routes (using an additional middleware)
		admin := protected.Group("/")
		admin.Use(Infrastructure.AdminMiddleware())
		{
			admin.POST("/tasks", ctrl.CreateTask)
			admin.PUT("/tasks/:id", ctrl.UpdateTask)
			admin.DELETE("/tasks/:id", ctrl.DeleteTask)
		}
	}

	return r
}
