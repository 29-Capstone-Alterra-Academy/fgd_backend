package user

import (
	"fgd/drivers/databases/notification"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role         UserRole  `json:"role"`
	Username     string    `json:"username" gorm:"unique"`
	Email        string    `json:"email" gorm:"unique"`
	Password     string    `json:"password"`
	ProfileImage *string   `json:"profile_image"`
	Gender       string    `json:"gender"`
	BirthDate    time.Time `json:"birth_date"`

	Notifications []notification.Notification `json:"notifications"`
}

type UserRole struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Type   string
}
