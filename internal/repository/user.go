package repository

import (
	"auth/internal/model"
	"context"
	"errors"
	"gorm.io/gorm"
)

type User interface {
	Get(ctx context.Context, params GetUserParams) (user model.User, err error)
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

func (r *UserDB) Create(ctx context.Context, user model.User) (model.User, error) {
	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserDB) Get(
	ctx context.Context,
	params GetUserParams,
) (user model.User, err error) {
	query := r.db.Model(&model.User{})

	if params.UserID != nil {
		query = query.Where(`user_id = ?`, *params.UserID)
	}
	if params.UserName != nil {
		query = query.Where(`username = ?`, *params.UserName)
	}

	err = query.First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.User{}, err
		}

		return model.User{}, err
	}

	return user, nil
}

type GetUserParams struct {
	UserID   *uint
	UserName *string
}
