package service

import "service/internal/repository"

type Services struct {
	User User
}

func NewServices(repos *repository.Repository) *Services {
	return &Services{
		User: NewUserService(repos.User),
	}
}
