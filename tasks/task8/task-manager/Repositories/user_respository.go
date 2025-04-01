package Repositories

import (
	"context"
	"task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Register(ctx context.Context, user Domain.User) (Domain.User, error)
	FindByUsername(ctx context.Context, username string) (Domain.User, error)
	Count(ctx context.Context) (int64, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database, collectionName string) UserRepository {
	return &userRepository{
		collection: db.Collection(collectionName),
	}
}

func (r *userRepository) Register(ctx context.Context, user Domain.User) (Domain.User, error) {
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return Domain.User{}, err
	}
	// Convert inserted ID to string.
	id, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		user.ID = id.Hex()
	}
	return user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (Domain.User, error) {
	var user Domain.User
	filter := bson.M{"username": username}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	return user, err
}

func (r *userRepository) Count(ctx context.Context) (int64, error) {
	return r.collection.CountDocuments(ctx, bson.D{})
}
