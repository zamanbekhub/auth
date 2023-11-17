package repository

import "gorm.io/gorm"

type Repository struct {
	User User
}

func NewRepositories(
	db *gorm.DB,
) *Repository {
	return &Repository{
		User: NewUserDB(db),
	}
}
