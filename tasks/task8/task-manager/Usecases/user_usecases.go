package Usecases

import (
	"context"
	"task-manager/Domain"
	"task-manager/Infrastructure"
	"task-manager/Repositories"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, user Domain.User) (Domain.User, error)
	LoginUser(ctx context.Context, username, password string) (Domain.User, error)
}

type userUsecase struct {
	userRepo Repositories.UserRepository
}

func NewUserUsecase(repo Repositories.UserRepository, timeout ...interface{}) UserUsecase {
	// Timeout not used here; you might add one if needed.
	return &userUsecase{
		userRepo: repo,
	}
}

func (uu *userUsecase) RegisterUser(ctx context.Context, user Domain.User) (Domain.User, error) {
	// Hash the user password using Infrastructure service.
	hashed, err := Infrastructure.HashPassword(user.Password)
	if err != nil {
		return Domain.User{}, err
	}
	user.Password = hashed

	// Set role if first user.
	count, err := uu.userRepo.Count(ctx)
	if err == nil && count == 0 {
		user.Role = "admin"
	} else {
		user.Role = "user"
	}

	return uu.userRepo.Register(ctx, user)
}

func (uu *userUsecase) LoginUser(ctx context.Context, username, password string) (Domain.User, error) {
	user, err := uu.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return Domain.User{}, err
	}
	if !Infrastructure.CheckPassword(user.Password, password) {
		return Domain.User{}, err
	}
	return user, nil
}
