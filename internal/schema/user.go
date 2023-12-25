package schema

import "auth/internal/model"

type UserCreate struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type GetUserData struct {
	model.TimestampMixin
	model.DeleteMixin
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}
