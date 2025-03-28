package controllers

import (
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService *data.TaskService
}

func NewTaskController(taskService *data.TaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.TaskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error retrieving tasks"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	task, err := tc.TaskService.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	createdTask, err := tc.TaskService.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating task"})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}
	updatedTask, err := tc.TaskService.UpdateTask(id, task)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := tc.TaskService.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
