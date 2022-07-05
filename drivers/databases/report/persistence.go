package report

import (
	"fgd/core/report"
	"fgd/drivers/databases/reply"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"
	"time"

	"gorm.io/gorm"
)

type persistenceReportRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceReportRepository) GetTopicReplyReports(topicId int, limit int, offset int) ([]report.Domain, error) {
	completeReports := []ReplyReportComplete{}
	reports := []ReplyReport{}
	domains := []report.Domain{}

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ReplyReportComplete{}
		completeReport.ID = report.ID

		_ = rp.Conn.Preload("Thread").Find(&completeReport.Reply, report.ReplyID)
		_ = rp.Conn.Preload("Topic").Find(&completeReport.Reply.Thread)
		if completeReport.Reply.Thread.Topic.ID != uint(topicId) {
			continue
		}
		_ = rp.Conn.Find(&completeReport.Reporter, report.UserID)

		completeReport.Topic = completeReport.Reply.Thread.Topic
		completeReport.Reason = report.Reason
		completeReport.Reviewed = report.Reviewed
		completeReport.CreatedAt = report.CreatedAt
		completeReports = append(completeReports, completeReport)
	}

	for _, report := range completeReports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetTopicThreadReports(topicId int, limit int, offset int) ([]report.Domain, error) {
	completeReports := []ThreadReportComplete{}
	reports := []ThreadReport{}
	domains := []report.Domain{}

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Where("topic_id = ?", topicId).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ThreadReportComplete{}
		completeReport.ID = report.ID

		_ = rp.Conn.Preload("Topic").Find(&completeReport.Thread, report.ThreadID)
		_ = rp.Conn.Find(&completeReport.Reporter, report.UserID)

		completeReport.Topic = completeReport.Thread.Topic
		completeReport.Reason = report.Reason
		completeReport.Reviewed = report.Reviewed
		completeReport.CreatedAt = report.CreatedAt
		completeReports = append(completeReports, completeReport)
	}

	for _, report := range completeReports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetReplyReports(limit, offset int) ([]report.Domain, error) {
	completeReports := []ReplyReportComplete{}
	reports := []ReplyReport{}
	domains := []report.Domain{}

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ReplyReportComplete{}
		completeReport.ID = report.ID

		_ = rp.Conn.Preload("Thread").Find(&completeReport.Reply, report.ReplyID)
		_ = rp.Conn.Preload("Topic").Find(&completeReport.Reply.Thread)
		_ = rp.Conn.Find(&completeReport.Reporter, report.UserID)

		completeReport.Topic = completeReport.Reply.Thread.Topic
		completeReport.Reason = report.Reason
		completeReport.Reviewed = report.Reviewed
		completeReport.CreatedAt = report.CreatedAt
		completeReports = append(completeReports, completeReport)
	}

	for _, report := range completeReports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetThreadReports(limit, offset int) ([]report.Domain, error) {
	completeReports := []ThreadReportComplete{}
	reports := []ThreadReport{}
	domains := []report.Domain{}

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ThreadReportComplete{}
		completeReport.ID = report.ID

		_ = rp.Conn.Preload("Topic").Find(&completeReport.Thread, report.ThreadID)
		_ = rp.Conn.Find(&completeReport.Reporter, report.UserID)

		completeReport.Topic = completeReport.Thread.Topic
		completeReport.Reason = report.Reason
		completeReport.Reviewed = report.Reviewed
		completeReport.CreatedAt = report.CreatedAt
		completeReports = append(completeReports, completeReport)
	}

	for _, report := range completeReports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetTopicReports(limit, offset int) ([]report.Domain, error) {
	completeReports := []TopicReportComplete{}
	reports := []TopicReport{}
	domains := []report.Domain{}

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := TopicReportComplete{}
		completeReport.ID = report.ID

		_ = rp.Conn.Find(&completeReport.Topic, report.TopicID)
		_ = rp.Conn.Find(&completeReport.Reporter, report.UserID)

		completeReport.Reason = report.Reason
		completeReport.CreatedAt = report.CreatedAt
		completeReports = append(completeReports, completeReport)
	}

	for _, report := range completeReports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetUserReports(limit, offset int) ([]report.Domain, error) {
	completeReports := []UserReportComplete{}
	reports := []UserReport{}
	domains := []report.Domain{}

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := UserReportComplete{}
		completeReport.ID = report.ID

		_ = rp.Conn.Find(&completeReport.Suspect, report.SuspectID)
		_ = rp.Conn.Find(&completeReport.Reporter, report.UserID)

		completeReport.Reason = report.Reason
		completeReport.CreatedAt = report.CreatedAt
		completeReports = append(completeReports, completeReport)
	}

	for _, report := range completeReports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) AddReason(data *report.Domain) error {
	reason := ReportReason{}

	newReason := reasonFromDomain(data)

	reason.Detail = newReason.Detail

	return rp.Conn.Create(&reason).Error
}

func (rp *persistenceReportRepository) ApproveReplyReport(replyReportId uint) error {
	report := ReplyReport{}
	fetchRes := rp.Conn.Take(&report, replyReportId)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	return rp.Conn.Delete(&reply.Reply{}, report.ReplyID).Error
}

func (rp *persistenceReportRepository) ApproveThreadReport(threadReportId uint) error {
	report := ThreadReport{}
	fetchRes := rp.Conn.Take(&report, threadReportId)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	return rp.Conn.Delete(&thread.Thread{}, report.ThreadID).Error
}

func (rp *persistenceReportRepository) ApproveTopicReport(topicReportId uint) error {
	report := TopicReport{}
	fetchRes := rp.Conn.Take(&report, topicReportId)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	return rp.Conn.Delete(&topic.Topic{}, report.TopicID).Error
}

func (rp *persistenceReportRepository) ApproveUserReport(userReportId uint) error {
	report := UserReport{}
	fetchRes := rp.Conn.Preload("User").Take(&report, userReportId)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	return rp.Conn.Delete(report.SuspectID).Error
}

func (rp *persistenceReportRepository) DeleteReason(reasonId uint) error {
	return rp.Conn.Delete(&ReportReason{}, reasonId).Error
}

func (rp *persistenceReportRepository) ForwardReplyReport(replyReportId uint) error {
	report := ReplyReport{}

	fetchRes := rp.Conn.Take(&report, replyReportId)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	report.Reviewed = true

	return rp.Conn.Save(&report).Error
}

func (rp *persistenceReportRepository) ForwardThreadReport(threadReportId uint) error {
	report := ThreadReport{}

	fetchRes := rp.Conn.Take(&report, threadReportId)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	report.Reviewed = true

	return rp.Conn.Save(&report).Error
}

func (rp *persistenceReportRepository) GetReasons() ([]report.Domain, error) {
	reasons := []ReportReason{}

	res := rp.Conn.Find(&reasons)
	if res.Error != nil {
		return []report.Domain{}, res.Error
	}

	domains := []report.Domain{}

	for _, reason := range reasons {
		domains = append(domains, reason.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) RemoveReplyReport(replyReportId uint) error {
	return rp.Conn.Delete(&ReplyReport{}, replyReportId).Error
}

func (rp *persistenceReportRepository) RemoveThreadReport(threadReportId uint) error {
	return rp.Conn.Delete(&ThreadReport{}, threadReportId).Error
}

func (rp *persistenceReportRepository) RemoveTopicReport(topicReportId uint) error {
	return rp.Conn.Delete(&TopicReport{}, topicReportId).Error
}

func (rp *persistenceReportRepository) RemoveUserReport(userReportId uint) error {
	return rp.Conn.Delete(&UserReport{}, userReportId).Error
}

func (rp *persistenceReportRepository) ReportReply(reporterId, replyId, reasonId uint) (report.Domain, error) {
	reply := reply.Reply{Model: gorm.Model{ID: replyId}}
	reporter := user.User{Model: gorm.Model{ID: reporterId}}
	reason := ReportReason{ID: reasonId}

	appendErr := rp.Conn.Model(&reply).Association("ReplyReports").Append(&reporter)
	if appendErr != nil {
		return report.Domain{}, appendErr
	}

	fetchReplyRes := rp.Conn.Preload("Thread").Take(&reply)
	if fetchReplyRes.Error != nil {
		return report.Domain{}, fetchReplyRes.Error
	}

	thread := reply.Thread
	fetchThreadRes := rp.Conn.Preload("Topic").Take(&thread)
	if fetchThreadRes.Error != nil {
		return report.Domain{}, fetchThreadRes.Error
	}

	report := ReplyReport{
		ID:        reporterId,
		ReplyID:   replyId,
		UserID:    reporterId,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	res := rp.Conn.Create(&report)

	completeReport := ReplyReportComplete{
		ID:        report.ID,
		Topic:     thread.Topic,
		Reply:     reply,
		Reporter:  reporter,
		Reason:    reason,
		Reviewed:  report.Reviewed,
		CreatedAt: report.CreatedAt,
	}

	return completeReport.toDomain(), res.Error
}

func (rp *persistenceReportRepository) ReportThread(reporterId, threadId, reasonId uint) (report.Domain, error) {
	thread := thread.Thread{Model: gorm.Model{ID: threadId}}
	reporter := user.User{Model: gorm.Model{ID: reporterId}}
	reason := ReportReason{ID: reasonId}

	appendErr := rp.Conn.Model(&thread).Association("ThreadReports").Append(&reporter)
	if appendErr != nil {
		return report.Domain{}, appendErr
	}

	fetchRes := rp.Conn.Preload("Topic").Take(&thread)
	if fetchRes.Error != nil {
		return report.Domain{}, fetchRes.Error
	}

	report := ThreadReport{
		ThreadID:  thread.ID,
		UserID:    reporter.ID,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	res := rp.Conn.Create(&report)

	completeReport := ThreadReportComplete{
		ID:        report.ID,
		Topic:     thread.Topic,
		Thread:    thread,
		Reporter:  reporter,
		Reason:    reason,
		Reviewed:  report.Reviewed,
		CreatedAt: report.CreatedAt,
	}

	return completeReport.toDomain(), res.Error
}

func (rp *persistenceReportRepository) ReportTopic(reporterId, topicId, reasonId uint) (report.Domain, error) {
	topic := topic.Topic{Model: gorm.Model{ID: topicId}}
	reporter := user.User{Model: gorm.Model{ID: reporterId}}
	reason := ReportReason{ID: reasonId}

	if reasonErr := rp.Conn.Take(&reason).Error; reasonErr != nil {
		return report.Domain{}, reasonErr
	}

	appendErr := rp.Conn.Model(&topic).Association("TopicReports").Append(&reporter)
	if appendErr != nil {
		return report.Domain{}, appendErr
	}

	report := TopicReport{
		UserID:    reporter.ID,
		TopicID:   topic.ID,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	res := rp.Conn.Create(&report)

	completeReport := TopicReportComplete{
		ID:        report.ID,
		Topic:     topic,
		Reporter:  reporter,
		Reason:    reason,
		CreatedAt: report.CreatedAt,
	}

	return completeReport.toDomain(), res.Error
}

func (rp *persistenceReportRepository) ReportUser(reporterId, suspectId, reasonId uint) (report.Domain, error) {
	suspect := user.User{Model: gorm.Model{ID: suspectId}}
	reporter := user.User{Model: gorm.Model{ID: reporterId}}
	reason := ReportReason{ID: reasonId}

	if reasonErr := rp.Conn.Take(&reason).Error; reasonErr != nil {
		return report.Domain{}, reasonErr
	}

	appendErr := rp.Conn.Model(&suspect).Association("UserReports").Append(&reporter)
	if appendErr != nil {
		return report.Domain{}, appendErr
	}

	report := UserReport{
		UserID:    reporter.ID,
		SuspectID: suspect.ID,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	res := rp.Conn.Create(&report)

	completeReport := UserReportComplete{
		ID:        reporterId,
		Reporter:  reporter,
		Suspect:   suspect,
		Reason:    reason,
		CreatedAt: report.CreatedAt,
	}

	return completeReport.toDomain(), res.Error
}

func InitPersistenceReportRepository(c *gorm.DB) report.Repository {
	return &persistenceReportRepository{
		Conn: c,
	}
}
