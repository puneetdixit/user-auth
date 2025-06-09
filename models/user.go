package models

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique"`
	Email      string `gorm:"unique"`
	Password   string
	IsVerified bool
}