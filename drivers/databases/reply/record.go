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
	Image    *ReplyImage `gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Content  string
}

type ReplyImage struct {
	gorm.Model
	ReplyID  uint
	ImageURL string
}
