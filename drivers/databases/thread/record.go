package thread

import (
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
	Content  string
	Image1   *string
	Image2   *string
	Image3   *string
	Image4   *string
	Image5   *string

	LikedBy   []*user.User `gorm:"many2many:liked_thread"`
	UnlikedBy []*user.User `gorm:"many2many:unliked_thread"`
	SavedBy   []*Thread    `gorm:"many2many:saved_thread"`
}
