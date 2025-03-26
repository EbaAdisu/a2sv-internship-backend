package data

import (
	"errors"
	"sync"
	"task_manager/models"
)

type TaskService struct {
	tasks map[int]models.Task
	mu    sync.Mutex
	nextID int
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[int]models.Task),
		nextID: 1,
	}
}

func (s *TaskService) GetAllTasks() []models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	tasks := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *TaskService) GetTaskByID(id int) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) CreateTask(task models.Task) models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	task.ID = s.nextID
	s.tasks[s.nextID] = task
	s.nextID++
	return task
}

func (s *TaskService) UpdateTask(id int, task models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}
	task.ID = id
	s.tasks[id] = task
	return task, nil
}

func (s *TaskService) DeleteTask(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, exists := s.tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}