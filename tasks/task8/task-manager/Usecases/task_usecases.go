package Usecases

import (
	"context"
	"time"

	"task-manager/Domain"
	"task-manager/Repositories"
)

type TaskUsecase interface {
	CreateTask(ctx context.Context, task Domain.Task) (Domain.Task, error)
	GetAllTasks(ctx context.Context) ([]Domain.Task, error)
	GetTaskByID(ctx context.Context, id string) (Domain.Task, error)
	UpdateTask(ctx context.Context, task Domain.Task) (Domain.Task, error)
	DeleteTask(ctx context.Context, id string) error
}

type taskUsecase struct {
	taskRepo       Repositories.TaskRepository
	contextTimeout time.Duration
}

// NewTaskUsecase creates a new TaskUsecase.
func NewTaskUsecase(repo Repositories.TaskRepository, timeout time.Duration) TaskUsecase {
	return &taskUsecase{
		taskRepo:       repo,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) CreateTask(ctx context.Context, task Domain.Task) (Domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()
	err := tu.taskRepo.Create(c, task)
	return task, err
}

func (tu *taskUsecase) GetAllTasks(ctx context.Context) ([]Domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()
	return tu.taskRepo.GetAll(c)
}
func (tu *taskUsecase) GetTaskByID(ctx context.Context, id string) (Domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()
	return tu.taskRepo.GetByID(c, id)
}

func (tu *taskUsecase) UpdateTask(ctx context.Context, task Domain.Task) (Domain.Task, error) {
	c, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()
	err := tu.taskRepo.Update(c, task)
	return task, err
}

func (tu *taskUsecase) DeleteTask(ctx context.Context, id string) error {
	c, cancel := context.WithTimeout(ctx, tu.contextTimeout)
	defer cancel()
	return tu.taskRepo.Delete(c, id)
}
