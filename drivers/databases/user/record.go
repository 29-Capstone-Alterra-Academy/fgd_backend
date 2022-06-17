package user

import (
	"fgd/drivers/databases/notification"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/topic"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role         UserRole
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	ProfileImage *string
	Gender       string
	BirthDate    time.Time

	Following []*User `gorm:"many2many:user_follow"`

	Notifications []notification.Notification
}

type UserRole struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Type   string
}
