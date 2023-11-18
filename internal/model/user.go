package model

type User struct {
	UserID   uint   `gorm:"primaryKey;user_id" json:"user_id"`
	FullName string `gorm:"full_name" json:"full_name"`
}
