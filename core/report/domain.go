package report

import "time"

type Domain struct {
	ID                   uint
	ReporterID           uint
	ReporterName         string
	ReporterProfileImage *string

	ReasonID     uint
	ReasonDetail string

	SuspectID           *uint
	SuspectUsername     *string
	SuspectProfileImage *string

	TopicID           *uint
	TopicName         *string
	TopicProfileImage *string

	ThreadID      *uint
	ThreadTitle   *string
	ThreadContent *string
	ThreadImage1  *string
	ThreadImage2  *string
	ThreadImage3  *string
	ThreadImage4  *string
	ThreadImage5  *string

	ReplyID      *uint
	ReplyContent *string
	ReplyImage   *string

	Reviewed  *bool
	CreatedAt time.Time
}

type Usecase interface {
	ReportUser(reporterId, suspectId, reasonId uint) (Domain, error)
	GetUserReports(limit, offset int) ([]Domain, error)
	ApproveUserReport(userReportId uint) error
	RemoveUserReport(userReportId uint) error

	ReportThread(reporterId, threadId, reasonId uint) (Domain, error)
	ForwardThreadReport(threadReportId uint) error
	GetTopicThreadReports(topicId, limit, offset int) ([]Domain, error)
	GetThreadReports(limit, offset int) ([]Domain, error)
	IgnoreThreadReport(threadReportId uint) error
	ApproveThreadReport(threadReportId uint) error
	RemoveThreadReport(threadReportId uint) error

	ReportReply(reporterId, replyId, reasonId uint) (Domain, error)
	ForwardReplyReport(replyReportId uint) error
	GetTopicReplyReports(topicId, limit, offset int) ([]Domain, error)
	GetReplyReports(limit, offset int) ([]Domain, error)
	IgnoreReplyReport(replyReportId uint) error
	ApproveReplyReport(replyReportId uint) error
	RemoveReplyReport(replyReportId uint) error

	ReportTopic(reporterId, topicId, reasonId uint) (Domain, error)
	GetTopicReports(limit, offset int) ([]Domain, error)
	ApproveTopicReport(topicReportId uint) error
	RemoveTopicReport(topicReportId uint) error

	AddReason(data *Domain) error
	GetReasons() ([]Domain, error)
	DeleteReason(reasonId uint) error
}

type Repository interface {
	ReportUser(reporterId, suspectId, reasonId uint) (Domain, error)
	GetUserReports(limit, offset int) ([]Domain, error)
	ApproveUserReport(userReportId uint) error
	RemoveUserReport(userReportId uint) error

	ReportThread(reporterId, threadId, reasonId uint) (Domain, error)
	ForwardThreadReport(threadReportId uint) error
	GetTopicThreadReports(topicId, limit, offset int) ([]Domain, error)
	GetThreadReports(limit, offset int) ([]Domain, error)
	ApproveThreadReport(threadReportId uint) error
	RemoveThreadReport(threadReportId uint) error

	ReportReply(reporterId, replyId, reasonId uint) (Domain, error)
	ForwardReplyReport(replyReportId uint) error
	ApproveReplyReport(replyReportId uint) error
	GetTopicReplyReports(topicId, limit, offset int) ([]Domain, error)
	GetReplyReports(limit, offset int) ([]Domain, error)
	RemoveReplyReport(replyReportId uint) error

	ReportTopic(reporterId, topicId, reasonId uint) (Domain, error)
	GetTopicReports(limit, offset int) ([]Domain, error)
	ApproveTopicReport(topicReportId uint) error
	RemoveTopicReport(topicReportId uint) error

	AddReason(data *Domain) error
	GetReasons() ([]Domain, error)
	DeleteReason(reasonId uint) error
}
