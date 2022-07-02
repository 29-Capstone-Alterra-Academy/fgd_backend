package thread

import (
	"fgd/core/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	TopicID  uint
	Topic    topic.Topic
	AuthorID uint
	Author   user.User `gorm:"ForeignKey:AuthorID"`
	Title    string
	Content  *string
	Image1   *string
	Image2   *string
	Image3   *string
	Image4   *string
	Image5   *string

	LikedBy   []*user.User `gorm:"many2many:liked_thread"`
	UnlikedBy []*user.User `gorm:"many2many:unliked_thread"`

	SavedBy []*Thread `gorm:"many2many:saved_thread"`
}

func (r *Thread) toDomain() *thread.Domain {
	return &thread.Domain{
		ID: int(r.ID),
		Author: thread.DomainAuthor{
			ID:       int(r.Author.ID),
			Username: r.Author.Username,
		},
		Topic: thread.DomainTopic{
			ID:           int(r.Topic.ID),
			Name:         r.Topic.Name,
			ProfileImage: r.Topic.ProfileImage,
		},
		Image1:      r.Image1,
		Image2:      r.Image2,
		Image3:      r.Image3,
		Image4:      r.Image4,
		Image5:      r.Image5,
		Title:       r.Title,
		Content:     r.Content,
		LikeCount:   0,
		UnlikeCount: 0,
		ReplyCount:  0,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		DeletedAt:   r.DeletedAt.Time,
	}
}

func fromDomain(threadDomain thread.Domain) *Thread {
	return &Thread{
		Model: gorm.Model{
			ID: uint(threadDomain.ID),
		},
		TopicID: uint(threadDomain.Topic.ID),
		Topic: topic.Topic{
			Model: gorm.Model{ID: uint(threadDomain.Topic.ID)},
		},
		AuthorID: uint(threadDomain.Author.ID),
		Author: user.User{
			Model: gorm.Model{ID: uint(threadDomain.Author.ID)},
		},
		Title:   threadDomain.Title,
		Content: threadDomain.Content,
		Image1:  threadDomain.Image1,
		Image2:  threadDomain.Image2,
		Image3:  threadDomain.Image3,
		Image4:  threadDomain.Image4,
		Image5:  threadDomain.Image5,
	}
}
