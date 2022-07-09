package report

import "time"

type Domain struct {
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
	ApproveUserReport(reporterId, suspectId uint) error
	RemoveUserReport(reporterId, suspectId uint) error

	ReportThread(reporterId, threadId, reasonId uint) (Domain, error)
	ForwardThreadReport(reporterId, threadId uint) error
	GetTopicThreadReports(topicId, limit, offset int) ([]Domain, error)
	GetThreadReports(limit, offset int) ([]Domain, error)
	IgnoreThreadReport(reporterId, threadId uint) error
	ApproveThreadReport(reporterId, threadId uint) error
	RemoveThreadReport(reporterId, threadId uint) error

	ReportReply(reporterId, replyId, reasonId uint) (Domain, error)
	ForwardReplyReport(reporterId, replyId uint) error
	GetTopicReplyReports(topicId, limit, offset int) ([]Domain, error)
	GetReplyReports(limit, offset int) ([]Domain, error)
	IgnoreReplyReport(reporterId, replyId uint) error
	ApproveReplyReport(reporterId, replyId uint) error
	RemoveReplyReport(reporterId, replyId uint) error

	ReportTopic(reporterId, topicId, reasonId uint) (Domain, error)
	GetTopicReports(limit, offset int) ([]Domain, error)
	ApproveTopicReport(reporterId, topicId uint) error
	RemoveTopicReport(reporterId, topicId uint) error

	AddReason(data *Domain) error
	GetReasons() ([]Domain, error)
	DeleteReason(reasonId uint) error
}

type Repository interface {
	ReportUser(reporterId, suspectId, reasonId uint) (Domain, error)
	GetUserReports(limit, offset int) ([]Domain, error)
	ApproveUserReport(reporterId, suspectId uint) error
	RemoveUserReport(reporterId, suspectId uint) error

	ReportThread(reporterId, threadId, reasonId uint) (Domain, error)
	ForwardThreadReport(reporterId, threadId uint) error
	GetTopicThreadReports(topicId, limit, offset int) ([]Domain, error)
	GetThreadReports(limit, offset int) ([]Domain, error)
	ApproveThreadReport(reporterId, threadId uint) error
	RemoveThreadReport(reporterId, threadId uint) error

	ReportReply(reporterId, replyId, reasonId uint) (Domain, error)
	ForwardReplyReport(reporterId, replyId uint) error
	ApproveReplyReport(reporterId, replyId uint) error
	GetTopicReplyReports(topicId, limit, offset int) ([]Domain, error)
	GetReplyReports(limit, offset int) ([]Domain, error)
	RemoveReplyReport(reporterId, replyId uint) error

	ReportTopic(reporterId, topicId, reasonId uint) (Domain, error)
	GetTopicReports(limit, offset int) ([]Domain, error)
	ApproveTopicReport(reporterId, replyId uint) error
	RemoveTopicReport(reporterId, replyId uint) error

	AddReason(data *Domain) error
	GetReasons() ([]Domain, error)
	DeleteReason(reasonId uint) error
}
