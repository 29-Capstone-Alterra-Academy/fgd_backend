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

	Moderating      []*topic.Topic `gorm:"many2many:topic_moderator"`
	SubscribedTopic []*topic.Topic `gorm:"many2many:subscribed_topic"`

	LikedThread   []*thread.Thread `gorm:"many2many:liked_thread"`
	UnlikedThread []*thread.Thread `gorm:"many2many:unliked_thread"`
	SavedThread   []*thread.Thread `gorm:"many2many:saved_thread"`

	Notifications []notification.Notification
}

type UserRole struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	Type   string
}
