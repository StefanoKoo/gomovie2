package models

type Users struct {
	Id       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}
