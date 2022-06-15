package topic

import (
	"os/user"

	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Name         string `gorm:"unique"`
	ProfileImage *string
	Description  string
	Rules        *string

	Moderators   []*user.User `gorm:"many2many:topic_moderator"`
	SubscribedBy []*user.User `gorm:"many2many:subscribed_topic"`
}
