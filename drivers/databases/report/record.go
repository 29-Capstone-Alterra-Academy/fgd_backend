package report

import (
	"fgd/core/report"
	"fgd/drivers/databases/reply"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"
	"time"
)

type UserReport struct {
	ID         uint `gorm:"primarykey"`
	SuspectID  uint
	Suspect    user.User `gorm:"ForeignKey:SuspectID"`
	ReporterID uint
	Reporter   user.User `gorm:"ForeignKey:ReporterID"`
	Reason     ReportReason
	Reviewed   bool
	CreatedAt  time.Time
}

func (r *UserReport) toDomain() report.Domain {
	return report.Domain{
		ReporterID:           r.Reporter.ID,
		ReporterName:         r.Reporter.Username,
		ReporterProfileImage: r.Reporter.ProfileImage,
		ReasonID:             r.Reason.ID,
		ReasonDetail:         r.Reason.Detail,
		SuspectID:            &r.Suspect.ID,
		SuspectUsername:      &r.Suspect.Username,
		SuspectProfileImage:  r.Suspect.ProfileImage,
	}
}

func userFromDomain(data *report.Domain) *UserReport {
	return &UserReport{
		SuspectID:  *data.SuspectID,
		ReporterID: data.ReporterID,
		Reason: ReportReason{
			ID: data.ReasonID,
		},
		Reviewed: *data.Reviewed,
	}
}

type TopicReport struct {
	ID         uint `gorm:"primarykey"`
	TopicID    uint
	Topic      topic.Topic
	ReporterID uint
	Reporter   user.User `gorm:"ForeignKey:ReporterID"`
	ReasonID   uint
	Reason     ReportReason
	CreatedAt  time.Time
}

func (r *TopicReport) toDomain() report.Domain {
	return report.Domain{
		ReporterID:           r.Reporter.ID,
		ReporterName:         r.Reporter.Username,
		ReporterProfileImage: r.Reporter.ProfileImage,
		ReasonID:             r.Reason.ID,
		ReasonDetail:         r.Reason.Detail,
		TopicID:              &r.Topic.ID,
		TopicName:            &r.Topic.Name,
		TopicProfileImage:    r.Topic.ProfileImage,
	}
}

func topicFromDomain(data *report.Domain) *TopicReport {
	return &TopicReport{
		TopicID:    *data.TopicID,
		ReporterID: data.ReporterID,
		Reason: ReportReason{
			ID: data.ReasonID,
		},
	}
}

type ThreadReport struct {
	ID         uint `gorm:"primarykey"`
	TopicID    uint
	Topic      topic.Topic
	ThreadID   uint
	Thread     thread.Thread
	ReporterID uint
	Reporter   user.User `gorm:"ForeignKey:ReporterID"`
	Reason     ReportReason
	Reviewed   bool
	CreatedAt  time.Time
}

func (r *ThreadReport) toDomain() report.Domain {
	return report.Domain{
		ReporterID:           r.Reporter.ID,
		ReporterName:         r.Reporter.Username,
		ReporterProfileImage: r.Reporter.ProfileImage,
		ReasonID:             r.Reason.ID,
		ReasonDetail:         r.Reason.Detail,
		TopicID:              &r.Topic.ID,
		TopicName:            &r.Topic.Name,
		TopicProfileImage:    r.Topic.ProfileImage,
		ThreadID:             &r.Thread.ID,
		ThreadTitle:          &r.Thread.Title,
		ThreadContent:        r.Thread.Content,
		ThreadImage1:         r.Thread.Image1,
		ThreadImage2:         r.Thread.Image2,
		ThreadImage3:         r.Thread.Image3,
		ThreadImage4:         r.Thread.Image4,
		ThreadImage5:         r.Thread.Image5,
	}
}

func threadFromDomain(data *report.Domain) *ThreadReport {
	return &ThreadReport{
		ThreadID:   *data.ThreadID,
		ReporterID: data.ReporterID,
		Reason: ReportReason{
			ID: data.ReasonID,
		},
		Reviewed: *data.Reviewed,
	}
}

type ReplyReport struct {
	ID         uint `gorm:"primaryKey"`
	TopicID    uint
	Topic      topic.Topic
	ReplyID    uint
	Reply      reply.Reply
	ReporterID uint
	Reporter   user.User `gorm:"ForeignKey:ReporterID"`
	Reason     ReportReason
	Reviewed   bool
	CreatedAt  time.Time
}

func (r *ReplyReport) toDomain() report.Domain {
	return report.Domain{
		ReporterID:           r.Reporter.ID,
		ReporterName:         r.Reporter.Username,
		ReporterProfileImage: r.Reporter.ProfileImage,
		ReasonID:             r.Reason.ID,
		ReasonDetail:         r.Reason.Detail,
		TopicID:              &r.Topic.ID,
		TopicName:            &r.Topic.Name,
		TopicProfileImage:    r.Topic.ProfileImage,
		ReplyID:              &r.Reply.ID,
		ReplyContent:         &r.Reply.Content,
		ReplyImage:           r.Reply.Image,
	}
}

func replyFromDomain(data *report.Domain) *ReplyReport {
	return &ReplyReport{
		ReplyID:    *data.ReplyID,
		ReporterID: data.ReporterID,
		Reason: ReportReason{
			ID: data.ReasonID,
		},
		Reviewed: *data.Reviewed,
	}
}

type ReportReason struct {
	ID     uint `gorm:"primaryKey"`
	Detail string
}

func (r *ReportReason) toDomain() report.Domain {
	return report.Domain{
		ReasonID:     r.ID,
		ReasonDetail: r.Detail,
	}
}

func reasonFromDomain(data *report.Domain) *ReportReason {
	return &ReportReason{
		ID:     data.ReasonID,
		Detail: data.ReasonDetail,
	}
}
