package models

type User struct {
	BaseModel

	Username   string `gorm:"unique"`
	Email      string `gorm:"unique"`
	Password   string
	IsVerified bool
}
