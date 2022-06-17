package reply

import (
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model
	ThreadID uint
	Thread   thread.Thread
	ParentID *uint
	Parent   *Reply
	UserID   uint
	User     user.User
	Image    *string
	Content  string

	LikedBy   []*user.User `gorm:"many2many:liked_reply"`
	UnlikedBy []*user.User `gorm:"many2many:unliked_reply"`
}
