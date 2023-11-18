package repository

import (
	"context"
	"gorm.io/gorm"
	"service/internal/model"
)

type User interface {
	GetAll(ctx context.Context) ([]model.User, error)
	Create(ctx context.Context, user model.User) (model.User, error)
}

type UserDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (r *UserDB) GetAll(ctx context.Context) (users []model.User, err error) {
	err = r.db.WithContext(ctx).
		Model(&model.User{}).
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserDB) Create(ctx context.Context, user model.User) (model.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
