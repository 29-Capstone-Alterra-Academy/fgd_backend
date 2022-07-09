package report

import (
	"fgd/core/report"
	"fgd/drivers/databases/reply"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"
	"time"
)

type UserReportComplete struct {
	Reporter  user.User
	Suspect   user.User
	Reason    ReportReason
	CreatedAt time.Time
}

type UserReport struct {
	UserID    uint `gorm:"primaryKey"`
	SuspectID uint `gorm:"primaryKey"`
	ReasonID  *uint
	Reason    ReportReason `gorm:"foreignKey:ReasonID"`
	CreatedAt time.Time
}

func (r *UserReportComplete) toDomain() report.Domain {
	return report.Domain{
		ReporterID:           r.Reporter.ID,
		ReporterName:         r.Reporter.Username,
		ReporterProfileImage: r.Reporter.ProfileImage,
		ReasonID:             r.Reason.ID,
		ReasonDetail:         r.Reason.Detail,
		SuspectID:            &r.Suspect.ID,
		SuspectUsername:      &r.Suspect.Username,
		SuspectProfileImage:  r.Suspect.ProfileImage,
		CreatedAt:            r.CreatedAt,
	}
}

func userFromDomain(data *report.Domain) *UserReport {
	return &UserReport{
		SuspectID: *data.SuspectID,
		UserID:    data.ReporterID,
		ReasonID:  &data.ReasonID,
	}
}

type TopicReportComplete struct {
	Topic     topic.Topic
	Reporter  user.User
	Reason    ReportReason
	CreatedAt time.Time
}

type TopicReport struct {
	TopicID   uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
	ReasonID  *uint
	Reason    ReportReason `gorm:"foreignKey:ReasonID"`
	CreatedAt time.Time
}

func (r *TopicReportComplete) toDomain() report.Domain {
	return report.Domain{
		ReporterID:           r.Reporter.ID,
		ReporterName:         r.Reporter.Username,
		ReporterProfileImage: r.Reporter.ProfileImage,
		ReasonID:             r.Reason.ID,
		ReasonDetail:         r.Reason.Detail,
		TopicID:              &r.Topic.ID,
		TopicName:            &r.Topic.Name,
		TopicProfileImage:    r.Topic.ProfileImage,
		CreatedAt:            r.CreatedAt,
	}
}

func topicFromDomain(data *report.Domain) *TopicReport {
	return &TopicReport{
		TopicID:  *data.TopicID,
		UserID:   data.ReporterID,
		ReasonID: &data.ReasonID,
	}
}

type ThreadReportComplete struct {
	Topic     topic.Topic
	Thread    thread.Thread
	Reporter  user.User
	Reason    ReportReason
	Reviewed  bool
	CreatedAt time.Time
}

type ThreadReport struct {
	TopicID   uint `gorm:"primaryKey"`
	ThreadID  uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
	ReasonID  *uint
	Reason    ReportReason `gorm:"foreignKey:ReasonID"`
	Reviewed  bool         `gorm:"default:false"`
	CreatedAt time.Time
}

func (r *ThreadReportComplete) toDomain() report.Domain {
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
		CreatedAt:            r.CreatedAt,
	}
}

func threadFromDomain(data *report.Domain) *ThreadReport {
	return &ThreadReport{
		TopicID:  *data.TopicID,
		ThreadID: *data.ThreadID,
		UserID:   data.ReporterID,
		Reason: ReportReason{
			ID: data.ReasonID,
		},
		Reviewed: *data.Reviewed,
	}
}

type ReplyReportComplete struct {
	Topic     topic.Topic
	Reply     reply.Reply
	Reporter  user.User
	Reason    ReportReason
	Reviewed  bool
	CreatedAt time.Time
}

type ReplyReport struct {
	TopicID   uint `gorm:"primaryKey"`
	ReplyID   uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
	ReasonID  *uint
	Reason    ReportReason `gorm:"foreignKey:ReasonID"`
	Reviewed  bool         `gorm:"default:false"`
	CreatedAt time.Time
}

func (r *ReplyReportComplete) toDomain() report.Domain {
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
		CreatedAt:            r.CreatedAt,
	}
}

func replyFromDomain(data *report.Domain) *ReplyReport {
	return &ReplyReport{
		TopicID: *data.TopicID,
		ReplyID: *data.ReplyID,
		UserID:  data.ReporterID,
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
