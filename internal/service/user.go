package service

import (
	"context"
	"service/internal/model"
	"service/internal/repository"
)

type User interface {
	GetAll(ctx context.Context) ([]model.User, error)
}

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u UserService) GetAll(ctx context.Context) ([]model.User, error) {
	return []model.User{
		{
			UserID:   1,
			FullName: "A",
		},
		{
			UserID:   2,
			FullName: "B",
		},
	}, nil
}
