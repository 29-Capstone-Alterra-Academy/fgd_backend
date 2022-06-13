package notification

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID   uint
	TopicID  *uint
	ThreadID *uint
	ReplyID  *uint
	IsRead   bool
	Type     NotificationType
}

type NotificationType struct {
	ID   uint `gorm:"primaryKey"`
	Kind string
}
