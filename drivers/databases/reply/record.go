package reply

import (
	"fgd/core/reply"
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
	AuthorID uint
	Author   user.User `gorm:"ForeignKey:AuthorID"`
	Image    *string
	Content  string

	LikedBy   []*user.User `gorm:"many2many:liked_reply"`
	UnlikedBy []*user.User `gorm:"many2many:unliked_reply"`

	ReplyReports []*user.User `gorm:"many2many:reply_reports"`
}

func (rec *Reply) toDomain() reply.Domain {
	return reply.Domain{
		ID: int(rec.ID),
		Author: reply.DomainAuthor{
			ID:           int(rec.Author.ID),
			Username:     rec.Author.Username,
			ProfileImage: rec.Author.ProfileImage,
		},
		Image:     rec.Image,
		Content:   rec.Content,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt.Time,
	}
}

func fromDomain(data *reply.Domain) *Reply {
	return &Reply{
		Image:   data.Image,
		Content: data.Content,
	}
}
