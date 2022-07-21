package topic

import (
	"fgd/core/topic"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name         string `gorm:"unique"`
	ProfileImage *string
	Description  string
	Rules        *string

	ModeratedBy  []*user.User `gorm:"many2many:topic_moderator"`
	SubscribedBy []*user.User `gorm:"many2many:subscribed_topic"`

	TopicReports []*user.User `gorm:"many2many:topic_reports;constraint:OnDelete:CASCADE;"`
}

func (r *Topic) toDomain() topic.Domain {
	topic := topic.Domain{
		ID:           int(r.ID),
		Name:         r.Name,
		ProfileImage: r.ProfileImage,
		Description:  r.Description,
		Rules:        r.Rules,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}

	if r.DeletedAt.Valid {
		topic.DeletedAt = &r.DeletedAt.Time
	}

	return topic
}

func fromDomain(topicDomain topic.Domain) *Topic {
	return &Topic{
		Model: gorm.Model{
			ID: uint(topicDomain.ID),
		},
		Name:         topicDomain.Name,
		ProfileImage: topicDomain.ProfileImage,
		Description:  topicDomain.Description,
		Rules:        topicDomain.Rules,
	}
}
