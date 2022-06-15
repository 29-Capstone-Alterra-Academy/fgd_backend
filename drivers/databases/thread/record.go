package thread

import (
	"fgd/drivers/databases/topic"
	"os/user"

	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	TopicID  uint
	Topic    topic.Topic
	AuthorID uint
	Author   user.User    `gorm:"ForeignKey:AuthorID"`
	Title    string
	Content  string
	Image    *ThreadImage `gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

  LikedBy []*user.User `gorm:"many2many:liked_thread"`
  UnlikedBy []*user.User `gorm:"many2many:unliked_thread"`
}

type ThreadImage struct {
	gorm.Model
	ThreadID  uint
	ImageURL1 string
	ImageURL2 string
	ImageURL3 string
	ImageURL4 string
	ImageURL5 string
}
