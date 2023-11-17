package model

type User struct {
	UserID   uint   `json:"user_id" gorm:"user_id"`
	FullName string `json:"full_name" gorm:"full_name"`
}
