package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	UserID   string   `json:"id"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Name     string   `json:"name"`
	Likes    []Movies `gorm:"foreignKey:Title" json:"likes"`
}
