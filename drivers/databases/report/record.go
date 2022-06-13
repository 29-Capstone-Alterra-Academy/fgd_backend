package report

import (
	"fgd/drivers/databases/reply"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"
	"time"
)

type UserReport struct {
	ID         uint `gorm:"primarykey"`
	ReporterID uint
	Reporter   user.User `gorm:"ForeignKey:ReporterID"`
	SuspectID  uint
	Suspect    user.User `gorm:"ForeignKey:SuspectID"`
	Reason     *string
	Reviewed   bool
	CreatedAt  time.Time
}

type TopicReport struct {
	ID         uint `gorm:"primarykey"`
	TopicID    uint
	Topic      topic.Topic
	ReporterID uint
	Reporter   user.User `gorm:"ForeignKey:ReporterID"`
	Reason     string
	CreatedAt  time.Time
}

type ThreadReport struct {
	ID         uint `gorm:"primarykey"`
	TopicID    uint
	Topic      topic.Topic
	ThreadID   uint
	Thread     thread.Thread
	ReplyID    *uint
	Reply      *reply.Reply
	ReporterID uint
	Reporter   user.User `gorm:"ForeignKey:ReporterID"`
	Reason     string
	Reviewed   bool
	CreatedAt  time.Time
}
