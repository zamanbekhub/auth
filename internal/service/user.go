package service

import (
	"auth/internal/model"
	"auth/internal/repository"
	"auth/internal/schema"
	"auth/utils/token"
	"context"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	GetByID(ctx context.Context, userID uint) (model.User, error)
	Login(ctx context.Context, data schema.LoginInput) (string, error)
	Register(ctx context.Context, data schema.UserCreate) (model.User, error)
}

type Userauth struct {
	userRepo repository.User
}

func NewUserauth(userRepo repository.User) *Userauth {
	return &Userauth{
		userRepo: userRepo,
	}
}

func (s *Userauth) GetByID(ctx context.Context, userID uint) (model.User, error) {
	user, err := s.userRepo.Get(ctx, repository.GetUserParams{UserID: &userID})
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (s *Userauth) Login(ctx context.Context, data schema.LoginInput) (string, error) {
	user, err := s.userRepo.Get(ctx, repository.GetUserParams{UserName: &data.Username})
	if err != nil {
		return "", err
	}

	err = verifyPassword(data.Password, user.Password)
	if err != nil {
		return "nil", err
	}

	tokenVal, err := token.GenerateToken(user.UserID)
	if err != nil {
		return "nil", err
	}

	return tokenVal, nil
}

func (s *Userauth) Register(ctx context.Context, data schema.UserCreate) (model.User, error) {
	hashedPassword, err := getPasswordHash(data.Password)
	if err != nil {
		return model.User{}, err
	}

	user, err := s.userRepo.Create(ctx, model.User{
		Username: data.Username,
		Password: hashedPassword,
	})
	if err != nil {
		return user, err
	}

	//metric.SuccessUserCreated.Inc()
	return user, nil
}

func getPasswordHash(password string) (string, error) {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
