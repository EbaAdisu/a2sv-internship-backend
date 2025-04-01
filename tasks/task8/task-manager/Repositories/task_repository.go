package Repositories

import (
	"context"
	"errors"
	"task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	Create(ctx context.Context, task Domain.Task) error
	GetAll(ctx context.Context) ([]Domain.Task, error)
	GetByID(ctx context.Context, id string) (Domain.Task, error)
	Update(ctx context.Context, task Domain.Task) error
	Delete(ctx context.Context, id string) error
}

type taskRepository struct {
	collection *mongo.Collection
}

// NewTaskRepository creates a new task repository.
func NewTaskRepository(db *mongo.Database, collectionName string) TaskRepository {
	return &taskRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *taskRepository) Create(ctx context.Context, task Domain.Task) error {
	_, err := r.collection.InsertOne(ctx, task)
	return err
}

func (r *taskRepository) GetAll(ctx context.Context) ([]Domain.Task, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	var tasks []Domain.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetByID(ctx context.Context, id string) (Domain.Task, error) {
	var task Domain.Task
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return task, err
	}

	filter := bson.M{"_id": objID}
	err = r.collection.FindOne(ctx, filter).Decode(&task)
	return task, err
}

func (r *taskRepository) Update(ctx context.Context, task Domain.Task) error {
	objID, err := primitive.ObjectIDFromHex(task.ID)
	if err != nil {
		return err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"due_date":    task.DueDate,
			"status":      task.Status,
		},
	}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (r *taskRepository) Delete(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": objID}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
