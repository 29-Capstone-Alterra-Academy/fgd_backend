package moderator

import (
	"fgd/core/moderator"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"
	"time"
)

type ModeratorRequest struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	User      user.User
	TopicID   uint
	Topic     topic.Topic
	CreatedAt time.Time
}

func (r *ModeratorRequest) toDomain() moderator.Domain {
	return moderator.Domain{
		ID:               r.ID,
		UserID:           r.UserID,
		Username:         r.User.Username,
		UserProfileImage: r.User.ProfileImage,
		TopicID:          r.TopicID,
		TopicName:        r.Topic.Name,
		CreatedAt:        r.CreatedAt,
	}
}
