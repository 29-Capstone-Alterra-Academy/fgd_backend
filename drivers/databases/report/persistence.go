package report

import (
	"fgd/core/report"
	"fgd/drivers/databases/reply"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type persistenceReportRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceReportRepository) GetTopicReplyReports(topicId int, limit int, offset int) ([]report.Domain, error) {
	completeReports := []ReplyReportComplete{}
	reports := []ReplyReport{}
	domains := []report.Domain{}

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Where("topic_id = ? AND reviewed = ?", topicId, false).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ReplyReportComplete{}

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

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Where("topic_id = ? AND reviewed = ?", topicId, false).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ThreadReportComplete{}

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

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Where("reviewed = ?", true).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ReplyReportComplete{}

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

	fetchRes := rp.Conn.Preload("Reason").Limit(limit).Offset(offset).Where("reviewed = ?", true).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ThreadReportComplete{}

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

func (rp *persistenceReportRepository) ApproveReplyReport(reporterId, replyId uint) error {
	report := ReplyReport{}
	fetchRes := rp.Conn.Where("user_id = ? AND reply_id = ?", reporterId, replyId).Take(&report)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	reporter := user.User{}
	fetchUserErr := rp.Conn.Take(&reporter, report.UserID).Error
	if fetchUserErr != nil {
		return fetchUserErr
	}

	assocErr := rp.Conn.Model(&reply.Reply{Model: gorm.Model{ID: report.ReplyID}}).Association("ReplyReports").Delete(&reporter)
	if assocErr != nil {
		return assocErr
	}

	delErr := rp.Conn.Delete(&reply.Reply{}, report.ReplyID).Error
	if delErr != nil {
		return delErr
	}

	cleanupErr := rp.Conn.Delete(&report).Error
	if cleanupErr != nil {
		return cleanupErr
	}

	return nil
}

func (rp *persistenceReportRepository) ApproveThreadReport(reporterId, threadId uint) error {
	report := ThreadReport{}
	fetchRes := rp.Conn.Where("user_id = ? AND thread_id = ?", reporterId, threadId).Take(&report)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	reporter := user.User{}
	fetchUserErr := rp.Conn.Take(&reporter, report.UserID).Error
	if fetchUserErr != nil {
		return fetchUserErr
	}

	assocErr := rp.Conn.Model(&thread.Thread{Model: gorm.Model{ID: report.ThreadID}}).Association("ThreadReports").Delete(&reporter)
	if assocErr != nil {
		return assocErr
	}

	delErr := rp.Conn.Delete(&thread.Thread{}, report.ThreadID).Error
	if delErr != nil {
		return delErr
	}

	cleanupErr := rp.Conn.Delete(&report).Error
	if cleanupErr != nil {
		return cleanupErr
	}

	return nil
}

func (rp *persistenceReportRepository) ApproveTopicReport(reporterId, topicId uint) error {
	report := TopicReport{}
	fetchRes := rp.Conn.Where("user_id = ? AND topic_id = ?", reporterId, topicId).Take(&report)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	reporter := user.User{}
	fetchUserErr := rp.Conn.Take(&reporter, report.UserID).Error
	if fetchUserErr != nil {
		return fetchUserErr
	}

	assocErr := rp.Conn.Model(&topic.Topic{Model: gorm.Model{ID: report.TopicID}}).Association("TopicReports").Delete(&reporter)
	if assocErr != nil {
		return assocErr
	}

	cleanupErr := rp.Conn.Delete(&report).Error
	if cleanupErr != nil {
		return cleanupErr
	}

	delReplyErr := rp.Conn.Where("topic_id = ?", report.TopicID).Delete(&reply.Reply{}).Error
	if delReplyErr != nil {
		return delReplyErr
	}

	delThreadErr := rp.Conn.Where("topic_id = ?", report.TopicID).Delete(&thread.Thread{}).Error
	if delThreadErr != nil {
		return delThreadErr
	}

	delTopicErr := rp.Conn.Delete(&topic.Topic{}, report.TopicID).Error
	if delTopicErr != nil {
		return delTopicErr
	}

	return nil
}

func (rp *persistenceReportRepository) ApproveUserReport(reporterId, suspectId uint) error {
	report := UserReport{}
	fetchRes := rp.Conn.Where("user_id = ? AND suspect_id = ?", reporterId, suspectId).Take(&report)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	reporter := user.User{}
	fetchUserErr := rp.Conn.Take(&reporter, report.UserID).Error
	if fetchUserErr != nil {
		return fetchUserErr
	}

	assocErr := rp.Conn.Model(&user.User{Model: gorm.Model{ID: report.SuspectID}}).Association("UserReports").Delete(&reporter)
	if assocErr != nil {
		return assocErr
	}

	delErr := rp.Conn.Delete(&user.User{}, report.SuspectID).Error
	if delErr != nil {
		return delErr
	}

	cleanupErr := rp.Conn.Delete(&report).Error
	if cleanupErr != nil {
		return cleanupErr
	}

	return nil
}

func (rp *persistenceReportRepository) DeleteReason(reasonId uint) error {
	return rp.Conn.Delete(&ReportReason{}, reasonId).Error
}

func (rp *persistenceReportRepository) ForwardReplyReport(reporterId, replyId uint) error {
	report := ReplyReport{}

	fetchRes := rp.Conn.Where("user_id = ? AND reply_id = ?", reporterId, replyId).Take(&report)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	report.Reviewed = true

	return rp.Conn.Save(&report).Error
}

func (rp *persistenceReportRepository) ForwardThreadReport(reporterId, threadId uint) error {
	report := ThreadReport{}

	fetchRes := rp.Conn.Where("user_id = ? AND thread_id = ?", reporterId, threadId).Take(&report)
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

func (rp *persistenceReportRepository) RemoveReplyReport(reporterId, replyId uint) error {
	report := ReplyReport{}
	fetchErr := rp.Conn.Where("user_id = ? AND reply_id = ?", reporterId, replyId).Take(&report).Error
	if fetchErr != nil {
		return fetchErr
	}

	assocErr := rp.Conn.Model(&reply.Reply{}).Association("ReplyReports").Delete(&report)
	if assocErr != nil {
		return assocErr
	}

	cleanupErr := rp.Conn.Delete(&report).Error
	if cleanupErr != nil {
		return cleanupErr
	}

	return nil
}

func (rp *persistenceReportRepository) RemoveThreadReport(reporterId, threadId uint) error {
	report := ThreadReport{}
	fetchErr := rp.Conn.Where("user_id = ? AND thread_id = ?", reporterId, threadId).Take(&report).Error
	if fetchErr != nil {
		return fetchErr
	}

	assocErr := rp.Conn.Model(&thread.Thread{}).Association("ThreadReports").Delete(&report)
	if assocErr != nil {
		return assocErr
	}

	cleanupErr := rp.Conn.Delete(&report).Error
	if cleanupErr != nil {
		return cleanupErr
	}

	return nil
}

func (rp *persistenceReportRepository) RemoveTopicReport(reporterId, topicId uint) error {
	report := TopicReport{}
	fetchErr := rp.Conn.Where("user_id = ? AND topic_id = ?", reporterId, topicId).Take(&report).Error
	if fetchErr != nil {
		return fetchErr
	}

	assocErr := rp.Conn.Model(&topic.Topic{}).Association("TopicReports").Delete(&report)
	if assocErr != nil {
		return assocErr
	}

	cleanupErr := rp.Conn.Delete(&report).Error
	if cleanupErr != nil {
		return cleanupErr
	}

	return nil
}

func (rp *persistenceReportRepository) RemoveUserReport(reporterId, suspectId uint) error {
	report := UserReport{}
	fetchErr := rp.Conn.Where("user_id = ? AND suspect_id = ?", reporterId, suspectId).Take(&report).Error
	if fetchErr != nil {
		return fetchErr
	}

	assocErr := rp.Conn.Model(&user.User{}).Association("UserReports").Delete(&report)
	if assocErr != nil {
		return assocErr
	}

	cleanupErr := rp.Conn.Delete(&report).Error
	if cleanupErr != nil {
		return cleanupErr
	}

	return nil
}

func (rp *persistenceReportRepository) ReportReply(reporterId, replyId, reasonId uint) (report.Domain, error) {
	completeReport := ReplyReportComplete{}
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
		TopicID:   thread.Topic.ID,
		ReplyID:   replyId,
		UserID:    reporterId,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	err := rp.Conn.Save(&report).Error
	if err != nil {
		return completeReport.toDomain(), err
	}

	takeErr := rp.Conn.Preload(clause.Associations).Take(&report).Error
	if takeErr != nil {
		return completeReport.toDomain(), takeErr
	}

	fetchReplyErr := rp.Conn.Take(&reply).Error
	if fetchReplyErr != nil {
		return completeReport.toDomain(), fetchReplyErr
	}

	fetchUserErr := rp.Conn.Take(&reporter).Error
	if fetchReplyErr != nil {
		return completeReport.toDomain(), fetchUserErr
	}

	completeReport.Reporter = reporter
	completeReport.Reply = reply
	completeReport.Topic = thread.Topic
	completeReport.Reason = reason
	completeReport.Reviewed = report.Reviewed
	completeReport.CreatedAt = report.CreatedAt

	return completeReport.toDomain(), nil
}

func (rp *persistenceReportRepository) ReportThread(reporterId, threadId, reasonId uint) (report.Domain, error) {
	completeReport := ThreadReportComplete{}
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
		TopicID:   thread.Topic.ID,
		ThreadID:  thread.ID,
		UserID:    reporter.ID,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	err := rp.Conn.Save(&report).Error
	if err != nil {
		return completeReport.toDomain(), err
	}

	takeErr := rp.Conn.Preload(clause.Associations).Take(&report).Error
	if takeErr != nil {
		return completeReport.toDomain(), takeErr
	}

	fetchThreadErr := rp.Conn.Take(&thread).Error
	if fetchThreadErr != nil {
		return completeReport.toDomain(), fetchThreadErr
	}

	fetchUserErr := rp.Conn.Take(&reporter).Error
	if fetchThreadErr != nil {
		return completeReport.toDomain(), fetchUserErr
	}

	completeReport.Reporter = reporter
	completeReport.Thread = thread
	completeReport.Topic = thread.Topic
	completeReport.Reason = reason
	completeReport.Reviewed = report.Reviewed
	completeReport.CreatedAt = report.CreatedAt

	return completeReport.toDomain(), nil
}

func (rp *persistenceReportRepository) ReportTopic(reporterId, topicId, reasonId uint) (report.Domain, error) {
	completeReport := TopicReportComplete{}
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

	err := rp.Conn.Save(&report).Error
	if err != nil {
		return completeReport.toDomain(), err
	}

	fetchTopicErr := rp.Conn.Take(&topic).Error
	if fetchTopicErr != nil {
		return completeReport.toDomain(), fetchTopicErr
	}

	fetchUserErr := rp.Conn.Take(&reporter).Error
	if fetchTopicErr != nil {
		return completeReport.toDomain(), fetchUserErr
	}

	completeReport.Reporter = reporter
	completeReport.Topic = topic
	completeReport.Reason = reason
	completeReport.CreatedAt = report.CreatedAt

	return completeReport.toDomain(), nil
}

func (rp *persistenceReportRepository) ReportUser(reporterId, suspectId, reasonId uint) (report.Domain, error) {
	completeReport := UserReportComplete{}
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

	err := rp.Conn.Save(&report).Error
	if err != nil {
		return completeReport.toDomain(), err
	}

	fetchSuspectErr := rp.Conn.Take(&suspect).Error
	if fetchSuspectErr != nil {
		return completeReport.toDomain(), fetchSuspectErr
	}

	fetchUserErr := rp.Conn.Take(&reporter).Error
	if fetchUserErr != nil {
		return completeReport.toDomain(), fetchUserErr
	}

	completeReport.Reporter = reporter
	completeReport.Suspect = suspect
	completeReport.Reason = reason
	completeReport.CreatedAt = report.CreatedAt

	return completeReport.toDomain(), nil
}

func InitPersistenceReportRepository(c *gorm.DB) report.Repository {
	return &persistenceReportRepository{
		Conn: c,
	}
}
