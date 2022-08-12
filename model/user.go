package model

import "gorm.io/gorm"

type User struct {
	UserID int    `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
}

type UserRepository interface {
	GetByUserID(db *gorm.DB, userID int) (*User, error)
	AddOrUpdate(db *gorm.DB, user *User) (*User, error)
	DeleteUser(db *gorm.DB, userID int) error
}
