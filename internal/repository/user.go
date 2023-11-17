package repository

import "gorm.io/gorm"

type User interface {
}

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}
