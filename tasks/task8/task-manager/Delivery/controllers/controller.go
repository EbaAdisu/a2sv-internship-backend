package controllers

import (
	"net/http"
	"time"

	"task-manager/Domain"
	"task-manager/Infrastructure"
	"task-manager/Usecases"

	"github.com/gin-gonic/gin"
)

// Controller holds the use cases.
type Controller struct {
	TaskUsecase Usecases.TaskUsecase
	UserUsecase Usecases.UserUsecase
}

// NewController creates a new controller instance.
func NewController(taskUsecase Usecases.TaskUsecase, userUsecase Usecases.UserUsecase) *Controller {
	return &Controller{
		TaskUsecase: taskUsecase,
		UserUsecase: userUsecase,
	}
}

// Register handles POST /register.
func (ctrl *Controller) Register(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	createdUser, err := ctrl.UserUsecase.RegisterUser(c, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	createdUser.Password = ""
	c.JSON(http.StatusCreated, createdUser)
}

// Login handles POST /login and returns a JWT token.
func (ctrl *Controller) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	user, err := ctrl.UserUsecase.LoginUser(c, input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}
	tokenString, err := Infrastructure.GenerateToken(user, time.Now().Add(24*time.Hour))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Token generation failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// GetAllTasks handles GET /tasks (accessible to authenticated users).
func (ctrl *Controller) GetAllTasks(c *gin.Context) {
	tasks, err := ctrl.TaskUsecase.GetAllTasks(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
func (ctrl *Controller) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := ctrl.TaskUsecase.GetTaskByID(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask handles POST /tasks (admin only).
func (ctrl *Controller) CreateTask(c *gin.Context) {
	var task Domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	createdTask, err := ctrl.TaskUsecase.CreateTask(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating task"})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

// UpdateTask handles PUT /tasks/:id (admin only).
func (ctrl *Controller) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task Domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	// Ensure the task to update has the correct ID.
	task.ID = id
	updatedTask, err := ctrl.TaskUsecase.UpdateTask(c, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating task"})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask handles DELETE /tasks/:id (admin only).
func (ctrl *Controller) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := ctrl.TaskUsecase.DeleteTask(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting task"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
