package models

type URL struct {
	BaseModel

	UserID          uint
	User 			User	`gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Url 			string 	`gorm:"url"`
	Timeout 		int		`gorm:"timeout"`
	NotifyOnTimeout bool	`gorm:"notify_on_timeout"` 
	NotifyEmail 	string 	`gorm:"notify_email"`
}