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
	ID        uint
	Reporter  user.User
	Suspect   user.User
	Reason    ReportReason
	CreatedAt time.Time
}

type UserReport struct {
	ID        uint `gorm:"primarykey"`
	UserID    uint `gorm:"primaryKey"`
	SuspectID uint `gorm:"primaryKey"`
	ReasonID  *uint
	Reason    ReportReason `gorm:"foreignKey:ReasonID"`
	CreatedAt time.Time
}

func (r *UserReportComplete) toDomain() report.Domain {
	return report.Domain{
		ID:                   r.ID,
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
		SuspectID: *data.SuspectID,
		UserID:    data.ReporterID,
		ReasonID:  &data.ReasonID,
	}
}

type TopicReportComplete struct {
	ID        uint
	Topic     topic.Topic
	Reporter  user.User
	Reason    ReportReason
	CreatedAt time.Time
}

type TopicReport struct {
	ID        uint `gorm:"primaryKey"`
	TopicID   uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
	ReasonID  *uint
	Reason    ReportReason `gorm:"foreignKey:ReasonID"`
	CreatedAt time.Time
}

func (r *TopicReportComplete) toDomain() report.Domain {
	return report.Domain{
		ID:                   r.ID,
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
		TopicID:  *data.TopicID,
		UserID:   data.ReporterID,
		ReasonID: &data.ReasonID,
	}
}

type ThreadReportComplete struct {
	ID        uint
	Topic     topic.Topic
	Thread    thread.Thread
	Reporter  user.User
	Reason    ReportReason
	Reviewed  bool
	CreatedAt time.Time
}

type ThreadReport struct {
	ID        uint
	ThreadID  uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
	ReasonID  *uint
	Reason    ReportReason `gorm:"foreignKey:ReasonID"`
	Reviewed  bool         `gorm:"default:false"`
	CreatedAt time.Time
}

func (r *ThreadReportComplete) toDomain() report.Domain {
	return report.Domain{
		ID:                   r.ID,
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
		ThreadID: *data.ThreadID,
		UserID:   data.ReporterID,
		Reason: ReportReason{
			ID: data.ReasonID,
		},
		Reviewed: *data.Reviewed,
	}
}

type ReplyReportComplete struct {
	ID        uint
	Topic     topic.Topic
	Reply     reply.Reply
	Reporter  user.User
	Reason    ReportReason
	Reviewed  bool
	CreatedAt time.Time
}

type ReplyReport struct {
	ID        uint
	ReplyID   uint `gorm:"primaryKey"`
	UserID    uint `gorm:"primaryKey"`
	ReasonID  *uint
	Reason    ReportReason `gorm:"foreignKey:ReasonID"`
	Reviewed  bool         `gorm:"default:false"`
	CreatedAt time.Time
}

func (r *ReplyReportComplete) toDomain() report.Domain {
	return report.Domain{
		ID:                   r.ID,
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
