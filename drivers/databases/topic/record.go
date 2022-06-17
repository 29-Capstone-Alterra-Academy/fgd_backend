package topic

import (
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
}
