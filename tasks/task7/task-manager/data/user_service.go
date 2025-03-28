package data

import (
	"context"
	"errors"
	"task_manager/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	collection *mongo.Collection
}

// NewUserService creates a new instance of UserService.
func NewUserService(collection *mongo.Collection) *UserService {
	return &UserService{collection: collection}
}

// RegisterUser creates a new user with a hashed password.
// If the collection is empty, the created user is an admin.
func (s *UserService) RegisterUser(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check if username exists
	var existing models.User
	err := s.collection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existing)
	if err == nil {
		return models.User{}, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	user.Password = string(hashedPassword)

	// Set role: if no users exist, assign admin role
	count, err := s.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return models.User{}, err
	}
	if count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	result, err := s.collection.InsertOne(ctx, user)
	if err != nil {
		return models.User{}, err
	}
	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

// LoginUser verifies the user credentials.
func (s *UserService) LoginUser(username, password string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user models.User
	err := s.collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("invalid credentials")
	}
	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, errors.New("invalid credentials")
	}
	return user, nil
}
