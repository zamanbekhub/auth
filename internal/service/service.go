package service

import (
	"auth/internal/repository"
)

type Services struct {
	User User
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		User: NewUserauth(repos.User),
	}
}
