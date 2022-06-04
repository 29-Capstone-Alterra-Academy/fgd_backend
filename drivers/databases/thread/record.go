package thread

import (
	"fgd/drivers/databases/topic"
	"os/user"

	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	TopicID  uint `json:"topic_id"`
	Topic    topic.Topic
	AuthorID uint         `json:"user_id"`
	Author   user.User    `gorm:"ForeignKey:AuthorID"`
	Title    string       `json:"title"`
	Content  string       `json:"content"`
	Image    *ThreadImage `json:"image" gorm:"contraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ThreadImage struct {
	gorm.Model
	ThreadID  uint   `json:"thread_id"`
	ImageURL1 string `json:"image1"`
	ImageURL2 string `json:"image2"`
	ImageURL3 string `json:"image3"`
	ImageURL4 string `json:"image4"`
	ImageURL5 string `json:"image5"`
}
