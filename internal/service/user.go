package service

import (
	"context"
	"service/internal/model"
	"service/internal/repository"
	"service/internal/schema"
)

type User interface {
	GetAll(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, data schema.UserCreate) (model.User, error)
}

type UserService struct {
	userRepo repository.User
}

func NewUserService(userRepo repository.User) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) GetAll(ctx context.Context) ([]model.User, error) {
	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserService) Create(ctx context.Context, data schema.UserCreate) (model.User, error) {
	user, err := s.userRepo.Create(ctx, model.User{
		FullName: data.FullName,
	})
	if err != nil {
		return user, err
	}

	return user, nil
}
