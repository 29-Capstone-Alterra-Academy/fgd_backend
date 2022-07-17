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
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	completeReports := []ReplyReportComplete{}
	reports := []ReplyReport{}
	domains := []report.Domain{}

	fetchResErr := tx.Preload("Reason").Limit(limit).Offset(offset).Where("topic_id = ? AND reviewed = ?", topicId, false).Find(&reports).Error
	if fetchResErr != nil {
		return domains, fetchResErr
	}

	for _, report := range reports {
		completeReport := ReplyReportComplete{}

		tx.Preload("Thread").Find(&completeReport.Reply, report.ReplyID)
		tx.Preload("Topic").Find(&completeReport.Reply.Thread)
		if completeReport.Reply.Thread.Topic.ID != uint(topicId) {
			continue
		}
		tx.Find(&completeReport.Reporter, report.UserID)

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
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	completeReports := []ThreadReportComplete{}
	reports := []ThreadReport{}
	domains := []report.Domain{}

	fetchResErr := tx.Preload("Reason").Limit(limit).Offset(offset).Where("topic_id = ? AND reviewed = ?", topicId, false).Find(&reports).Error
	if fetchResErr != nil {
		return domains, fetchResErr
	}

	for _, report := range reports {
		completeReport := ThreadReportComplete{}

		tx.Preload("Topic").Find(&completeReport.Thread, report.ThreadID)
		tx.Find(&completeReport.Reporter, report.UserID)

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
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	completeReports := []ReplyReportComplete{}
	reports := []ReplyReport{}
	domains := []report.Domain{}

	fetchRes := tx.Preload("Reason").Limit(limit).Offset(offset).Where("reviewed = ?", true).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ReplyReportComplete{}

		tx.Preload("Thread").Find(&completeReport.Reply, report.ReplyID)
		tx.Preload("Topic").Find(&completeReport.Reply.Thread)
		tx.Find(&completeReport.Reporter, report.UserID)

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
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	completeReports := []ThreadReportComplete{}
	reports := []ThreadReport{}
	domains := []report.Domain{}

	fetchRes := tx.Preload("Reason").Limit(limit).Offset(offset).Where("reviewed = ?", true).Find(&reports)
	if fetchRes.Error != nil {
		return domains, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := ThreadReportComplete{}

		tx.Preload("Topic").Find(&completeReport.Thread, report.ThreadID)
		tx.Find(&completeReport.Reporter, report.UserID)

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
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	completeReports := []TopicReportComplete{}
	reports := []TopicReport{}
	domains := []report.Domain{}

	fetchRes := tx.Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := TopicReportComplete{}

		tx.Find(&completeReport.Topic, report.TopicID)
		tx.Find(&completeReport.Reporter, report.UserID)

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
	tx := rp.Conn.Session(&gorm.Session{SkipDefaultTransaction: true})

	completeReports := []UserReportComplete{}
	reports := []UserReport{}
	domains := []report.Domain{}

	fetchRes := tx.Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	for _, report := range reports {
		completeReport := UserReportComplete{}

		tx.Find(&completeReport.Suspect, report.SuspectID)
		tx.Find(&completeReport.Reporter, report.UserID)

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
	tx := rp.Conn.Begin()

	reason := ReportReason{}

	newReason := reasonFromDomain(data)

	reason.Detail = newReason.Detail

	err := tx.Create(&reason).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) ApproveReplyReport(reporterId, replyId uint) error {
	tx := rp.Conn.Begin()

	report := ReplyReport{}
	fetchResErr := tx.Where("user_id = ? AND reply_id = ?", reporterId, replyId).Take(&report).Error
	if fetchResErr != nil {
		return fetchResErr
	}

	reporter := user.User{}
	fetchUserErr := tx.Take(&reporter, report.UserID).Error
	if fetchUserErr != nil {
		return fetchUserErr
	}

	tx.SavePoint("checkpoint")

	assocErr := tx.Model(&reply.Reply{Model: gorm.Model{ID: report.ReplyID}}).Association("ReplyReports").Delete(&reporter)
	if assocErr != nil {
		tx.RollbackTo("checkpoint")
		return assocErr
	}

	delErr := tx.Delete(&reply.Reply{}, report.ReplyID).Error
	if delErr != nil {
		tx.RollbackTo("checkpoint")
		return delErr
	}

	cleanupErr := tx.Delete(&report).Error
	if cleanupErr != nil {
		tx.RollbackTo("checkpoint")
		return cleanupErr
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) ApproveThreadReport(reporterId, threadId uint) error {
	tx := rp.Conn.Begin()

	report := ThreadReport{}
	fetchResErr := tx.Where("user_id = ? AND thread_id = ?", reporterId, threadId).Take(&report).Error
	if fetchResErr != nil {
		return fetchResErr
	}

	tx.SavePoint("checkpoint")

	reporter := user.User{}
	fetchUserErr := tx.Take(&reporter, report.UserID).Error
	if fetchUserErr != nil {
		tx.RollbackTo("checkpoint")
		return fetchUserErr
	}

	assocErr := tx.Model(&thread.Thread{Model: gorm.Model{ID: report.ThreadID}}).Association("ThreadReports").Delete(&reporter)
	if assocErr != nil {
		tx.RollbackTo("checkpoint")
		return assocErr
	}

	delErr := tx.Delete(&thread.Thread{}, report.ThreadID).Error
	if delErr != nil {
		tx.RollbackTo("checkpoint")
		return delErr
	}

	cleanupErr := tx.Delete(&report).Error
	if cleanupErr != nil {
		tx.RollbackTo("checkpoint")
		return cleanupErr
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) ApproveTopicReport(reporterId, topicId uint) error {
	tx := rp.Conn.Begin()

	report := TopicReport{}
	fetchResErr := tx.Where("user_id = ? AND topic_id = ?", reporterId, topicId).Take(&report).Error
	if fetchResErr != nil {
		return fetchResErr
	}

	reporter := user.User{}
	fetchUserErr := tx.Take(&reporter, report.UserID).Error
	if fetchUserErr != nil {
		return fetchUserErr
	}

	tx.SavePoint("checkpoint")

	assocErr := tx.Model(&topic.Topic{Model: gorm.Model{ID: report.TopicID}}).Association("TopicReports").Delete(&reporter)
	if assocErr != nil {
		tx.RollbackTo("checkpoint")
		return assocErr
	}

	cleanupErr := tx.Delete(&report).Error
	if cleanupErr != nil {
		tx.RollbackTo("checkpoint")
		return cleanupErr
	}

	delReplyErr := tx.Where("topic_id = ?", report.TopicID).Delete(&reply.Reply{}).Error
	if delReplyErr != nil {
		tx.RollbackTo("checkpoint")
		return delReplyErr
	}

	delThreadErr := tx.Where("topic_id = ?", report.TopicID).Delete(&thread.Thread{}).Error
	if delThreadErr != nil {
		tx.RollbackTo("checkpoint")
		return delThreadErr
	}

	delTopicErr := tx.Delete(&topic.Topic{}, report.TopicID).Error
	if delTopicErr != nil {
		tx.RollbackTo("checkpoint")
		return delTopicErr
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) ApproveUserReport(reporterId, suspectId uint) error {
	tx := rp.Conn.Begin()

	report := UserReport{}
	fetchResErr := tx.Where("user_id = ? AND suspect_id = ?", reporterId, suspectId).Take(&report).Error
	if fetchResErr != nil {
		return fetchResErr
	}

	reporter := user.User{}
	fetchUserErr := tx.Take(&reporter, report.UserID).Error
	if fetchUserErr != nil {
		return fetchUserErr
	}

	tx.SavePoint("checkpoint")

	assocErr := tx.Model(&user.User{Model: gorm.Model{ID: report.SuspectID}}).Association("UserReports").Delete(&reporter)
	if assocErr != nil {
		tx.RollbackTo("checkpoint")
		return assocErr
	}

	delErr := tx.Delete(&user.User{}, report.SuspectID).Error
	if delErr != nil {
		tx.RollbackTo("checkpoint")
		return delErr
	}

	cleanupErr := tx.Delete(&report).Error
	if cleanupErr != nil {
		tx.RollbackTo("checkpoint")
		return cleanupErr
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) DeleteReason(reasonId uint) error {
	tx := rp.Conn.Begin()
	err := tx.Delete(&ReportReason{}, reasonId).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) ForwardReplyReport(reporterId, replyId uint) error {
	tx := rp.Conn.Begin()

	report := ReplyReport{}

	fetchResErr := tx.Where("user_id = ? AND reply_id = ?", reporterId, replyId).Take(&report).Error
	if fetchResErr != nil {
		return fetchResErr
	}

	report.Reviewed = true

	err := tx.Save(&report).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) ForwardThreadReport(reporterId, threadId uint) error {
	tx := rp.Conn.Begin()

	report := ThreadReport{}

	fetchResErr := tx.Where("user_id = ? AND thread_id = ?", reporterId, threadId).Take(&report).Error
	if fetchResErr != nil {
		return fetchResErr
	}

	report.Reviewed = true

	err := tx.Save(&report).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
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
	tx := rp.Conn.Begin()

	report := ReplyReport{}
	fetchErr := tx.Where("user_id = ? AND reply_id = ?", reporterId, replyId).Take(&report).Error
	if fetchErr != nil {
		return fetchErr
	}

	tx.SavePoint("checkpoint")

	assocErr := tx.Model(&reply.Reply{}).Association("ReplyReports").Delete(&report)
	if assocErr != nil {
		tx.RollbackTo("checkpoint")
		return assocErr
	}

	cleanupErr := tx.Delete(&report).Error
	if cleanupErr != nil {
		tx.RollbackTo("checkpoint")
		return cleanupErr
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) RemoveThreadReport(reporterId, threadId uint) error {
	tx := rp.Conn.Begin()

	report := ThreadReport{}
	fetchErr := tx.Where("user_id = ? AND thread_id = ?", reporterId, threadId).Take(&report).Error
	if fetchErr != nil {
		return fetchErr
	}

	tx.SavePoint("checkpoint")

	assocErr := tx.Model(&thread.Thread{}).Association("ThreadReports").Delete(&report)
	if assocErr != nil {
		tx.RollbackTo("checkpoint")
		return assocErr
	}

	cleanupErr := tx.Delete(&report).Error
	if cleanupErr != nil {
		tx.RollbackTo("checkpoint")
		return cleanupErr
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) RemoveTopicReport(reporterId, topicId uint) error {
	tx := rp.Conn.Begin()

	report := TopicReport{}
	fetchErr := tx.Where("user_id = ? AND topic_id = ?", reporterId, topicId).Take(&report).Error
	if fetchErr != nil {
		return fetchErr
	}

	tx.SavePoint("checkpoint")

	assocErr := tx.Model(&topic.Topic{}).Association("TopicReports").Delete(&report)
	if assocErr != nil {
		tx.RollbackTo("checkpoint")
		return assocErr
	}

	cleanupErr := tx.Delete(&report).Error
	if cleanupErr != nil {
		tx.RollbackTo("checkpoint")
		return cleanupErr
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) RemoveUserReport(reporterId, suspectId uint) error {
	tx := rp.Conn.Begin()

	report := UserReport{}
	fetchErr := tx.Where("user_id = ? AND suspect_id = ?", reporterId, suspectId).Take(&report).Error
	if fetchErr != nil {
		return fetchErr
	}

	tx.SavePoint("checkpoint")

	assocErr := tx.Model(&user.User{}).Association("UserReports").Delete(&report)
	if assocErr != nil {
		tx.RollbackTo("checkpoint")
		return assocErr
	}

	cleanupErr := tx.Delete(&report).Error
	if cleanupErr != nil {
		tx.RollbackTo("checkpoint")
		return cleanupErr
	}

	return tx.Commit().Error
}

func (rp *persistenceReportRepository) ReportReply(reporterId, replyId, reasonId uint) (report.Domain, error) {
	tx := rp.Conn.Begin()

	completeReport := ReplyReportComplete{}
	reply := reply.Reply{Model: gorm.Model{ID: replyId}}
	reporter := user.User{Model: gorm.Model{ID: reporterId}}
	reason := ReportReason{ID: reasonId}

	tx.SavePoint("checkpoint")

	appendErr := tx.Model(&reply).Association("ReplyReports").Append(&reporter)
	if appendErr != nil {
		tx.RollbackTo("checkpoint")
		return report.Domain{}, appendErr
	}

	fetchReplyRes := tx.Preload("Thread").Take(&reply)
	if fetchReplyRes.Error != nil {
		tx.RollbackTo("checkpoint")
		return report.Domain{}, fetchReplyRes.Error
	}

	thread := reply.Thread
	fetchThreadErr := tx.Preload("Topic").Take(&thread).Error
	if fetchThreadErr != nil {
		tx.RollbackTo("checkpoint")
		return report.Domain{}, fetchThreadErr
	}

	report := ReplyReport{
		TopicID:   thread.Topic.ID,
		ReplyID:   replyId,
		UserID:    reporterId,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	err := tx.Save(&report).Error
	if err != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), err
	}

	fetchReplyErr := tx.Take(&reply).Error
	if fetchReplyErr != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), fetchReplyErr
	}

	fetchUserErr := tx.Take(&reporter).Error
	if fetchReplyErr != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), fetchUserErr
	}

	completeReport.Reporter = reporter
	completeReport.Reply = reply
	completeReport.Topic = thread.Topic
	completeReport.Reason = reason
	completeReport.Reviewed = report.Reviewed
	completeReport.CreatedAt = report.CreatedAt

	return completeReport.toDomain(), tx.Commit().Error
}

func (rp *persistenceReportRepository) ReportThread(reporterId, threadId, reasonId uint) (report.Domain, error) {
	tx := rp.Conn.Begin()

	completeReport := ThreadReportComplete{}
	thread := thread.Thread{Model: gorm.Model{ID: threadId}}
	reporter := user.User{Model: gorm.Model{ID: reporterId}}
	reason := ReportReason{ID: reasonId}

	tx.SavePoint("checkpoint")

	appendErr := tx.Model(&thread).Association("ThreadReports").Append(&reporter)
	if appendErr != nil {
		tx.RollbackTo("checkpoint")
		return report.Domain{}, appendErr
	}

	fetchErr := tx.Preload("Topic").Take(&thread).Error
	if fetchErr != nil {
		tx.RollbackTo("checkpoint")
		return report.Domain{}, fetchErr
	}

	report := ThreadReport{
		TopicID:   thread.Topic.ID,
		ThreadID:  thread.ID,
		UserID:    reporter.ID,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	err := tx.Save(&report).Error
	if err != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), err
	}

	fetchThreadErr := tx.Take(&thread).Error
	if fetchThreadErr != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), fetchThreadErr
	}

	fetchUserErr := tx.Take(&reporter).Error
	if fetchThreadErr != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), fetchUserErr
	}

	completeReport.Reporter = reporter
	completeReport.Thread = thread
	completeReport.Topic = thread.Topic
	completeReport.Reason = reason
	completeReport.Reviewed = report.Reviewed
	completeReport.CreatedAt = report.CreatedAt

	return completeReport.toDomain(), tx.Commit().Error
}

func (rp *persistenceReportRepository) ReportTopic(reporterId, topicId, reasonId uint) (report.Domain, error) {
	tx := rp.Conn.Begin()

	completeReport := TopicReportComplete{}
	topic := topic.Topic{Model: gorm.Model{ID: topicId}}
	reporter := user.User{Model: gorm.Model{ID: reporterId}}
	reason := ReportReason{ID: reasonId}

	tx.SavePoint("checkpoint")

	reasonErr := tx.Take(&reason).Error
	if reasonErr != nil {
		tx.RollbackTo("checkpoint")
		return report.Domain{}, reasonErr
	}

	appendErr := tx.Model(&topic).Association("TopicReports").Append(&reporter)
	if appendErr != nil {
		tx.RollbackTo("checkpoint")
		return report.Domain{}, appendErr
	}

	report := TopicReport{
		UserID:    reporter.ID,
		TopicID:   topic.ID,
		Reason:    reason,
		CreatedAt: time.Now(),
	}

	err := tx.Save(&report).Error
	if err != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), err
	}

	fetchTopicErr := tx.Take(&topic).Error
	if fetchTopicErr != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), fetchTopicErr
	}

	fetchUserErr := tx.Take(&reporter).Error
	if fetchTopicErr != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), fetchUserErr
	}

	completeReport.Reporter = reporter
	completeReport.Topic = topic
	completeReport.Reason = reason
	completeReport.CreatedAt = report.CreatedAt

	return completeReport.toDomain(), tx.Commit().Error
}

func (rp *persistenceReportRepository) ReportUser(reporterId, suspectId, reasonId uint) (report.Domain, error) {
	tx := rp.Conn.Begin()

	completeReport := UserReportComplete{}
	suspect := user.User{Model: gorm.Model{ID: suspectId}}
	reporter := user.User{Model: gorm.Model{ID: reporterId}}
	reason := ReportReason{ID: reasonId}

	tx.SavePoint("checkpoint")

	reasonErr := rp.Conn.Take(&reason).Error
	if reasonErr != nil {
		tx.RollbackTo("checkpoint")
		return report.Domain{}, reasonErr
	}

	appendErr := rp.Conn.Model(&suspect).Association("UserReports").Append(&reporter)
	if appendErr != nil {
		tx.RollbackTo("checkpoint")
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
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), err
	}

	fetchSuspectErr := rp.Conn.Take(&suspect).Error
	if fetchSuspectErr != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), fetchSuspectErr
	}

	fetchUserErr := rp.Conn.Take(&reporter).Error
	if fetchUserErr != nil {
		tx.RollbackTo("checkpoint")
		return completeReport.toDomain(), fetchUserErr
	}

	completeReport.Reporter = reporter
	completeReport.Suspect = suspect
	completeReport.Reason = reason
	completeReport.CreatedAt = report.CreatedAt

	return completeReport.toDomain(), tx.Commit().Error
}

func InitPersistenceReportRepository(c *gorm.DB) report.Repository {
	return &persistenceReportRepository{
		Conn: c,
	}
}
