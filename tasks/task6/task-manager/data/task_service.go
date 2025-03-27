package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	collection *mongo.Collection
}

func NewTaskService(collection *mongo.Collection) *TaskService {
	return &TaskService{
		collection: collection,
	}
}

func (s *TaskService) GetAllTasks() ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *TaskService) GetTaskByID(id string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid task ID")
	}

	var task models.Task
	err = s.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return models.Task{}, errors.New("task not found")
	}
	return task, nil
}

func (s *TaskService) CreateTask(task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := s.collection.InsertOne(ctx, task)
	if err != nil {
		return models.Task{}, err
	}
	task.ID = result.InsertedID.(primitive.ObjectID)
	return task, nil
}

func (s *TaskService) UpdateTask(id string, task models.Task) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.New("invalid task ID")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"title":       task.Title,
		"description": task.Description,
		"due_date":    task.DueDate,
		"status":      task.Status,
	}}

	result, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil || result.MatchedCount == 0 {
		return models.Task{}, errors.New("task not found")
	}

	task.ID = objID
	return task, nil
}

func (s *TaskService) DeleteTask(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid task ID")
	}

	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil || result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
