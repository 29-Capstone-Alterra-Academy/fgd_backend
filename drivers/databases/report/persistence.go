package report

import (
	"fgd/core/report"
	"fgd/drivers/databases/reply"
	"fgd/drivers/databases/thread"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type persistenceReportRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceReportRepository) GetTopicReplyReports(topicId int, limit int, offset int) ([]report.Domain, error) {
	reports := []ReplyReport{}

	fetchRes := rp.Conn.Preload("Reporter").Preload("Topic").Preload("Reply").Preload("Reason").Limit(limit).Offset(offset).Where("topic_id = ?", topicId).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	domains := []report.Domain{}

	for _, report := range reports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetTopicThreadReports(topicId int, limit int, offset int) ([]report.Domain, error) {
	reports := []ThreadReport{}

	fetchRes := rp.Conn.Preload("Reporter").Preload("Topic").Preload("Thread").Preload("Reason").Limit(limit).Offset(offset).Where("topic_id = ?", topicId).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	domains := []report.Domain{}

	for _, report := range reports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetReplyReports(limit, offset int) ([]report.Domain, error) {
	reports := []ReplyReport{}

	fetchRes := rp.Conn.Preload("Reporter").Preload("Topic").Preload("Reply").Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	domains := []report.Domain{}

	for _, report := range reports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetThreadReports(limit, offset int) ([]report.Domain, error) {
	reports := []ThreadReport{}

	fetchRes := rp.Conn.Preload("Reporter").Preload("Topic").Preload("Thread").Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	domains := []report.Domain{}

	for _, report := range reports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetTopicReports(limit, offset int) ([]report.Domain, error) {
	reports := []TopicReport{}

	fetchRes := rp.Conn.Preload("Reporter").Preload("Topic").Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	domains := []report.Domain{}

	for _, report := range reports {
		domains = append(domains, report.toDomain())
	}

	return domains, nil
}

func (rp *persistenceReportRepository) GetUserReports(limit, offset int) ([]report.Domain, error) {
	reports := []UserReport{}

	fetchRes := rp.Conn.Preload("Reporter").Preload("Suspect").Preload("Reason").Limit(limit).Offset(offset).Find(&reports)
	if fetchRes.Error != nil {
		return []report.Domain{}, fetchRes.Error
	}

	domains := []report.Domain{}

	for _, report := range reports {
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
	fetchRes := rp.Conn.Preload("Reply").Take(&report, replyReportId)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	return rp.Conn.Delete(report.Reply).Error
}

func (rp *persistenceReportRepository) ApproveThreadReport(threadReportId uint) error {
	report := ThreadReport{}
	fetchRes := rp.Conn.Preload("Thread").Take(&report, threadReportId)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	return rp.Conn.Delete(report.Thread).Error
}

func (rp *persistenceReportRepository) ApproveTopicReport(topicReportId uint) error {
	report := TopicReport{}
	fetchRes := rp.Conn.Preload("Topic").Take(&report, topicReportId)
	if fetchRes.Error != nil {
		return fetchRes.Error
	}

	return rp.Conn.Delete(report.Topic).Error
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
		Reporter: user.User{Model: gorm.Model{ID: reporterId}},
		Reply:    reply,
		Topic:    thread.Topic,
		Reason:   ReportReason{ID: reasonId},
	}

	res := rp.Conn.Create(&report)

	return report.toDomain(), res.Error
}

func (rp *persistenceReportRepository) ReportThread(reporterId, threadId, reasonId uint) (report.Domain, error) {
	thread := thread.Thread{Model: gorm.Model{ID: threadId}}
	fetchRes := rp.Conn.Preload("Topic").Take(&thread)
	if fetchRes.Error != nil {
		return report.Domain{}, fetchRes.Error
	}

	report := ThreadReport{
		Reporter: user.User{Model: gorm.Model{ID: reporterId}},
		Thread:   thread,
		Topic:    thread.Topic,
		Reason:   ReportReason{ID: reasonId},
	}

	res := rp.Conn.Create(&report)

	return report.toDomain(), res.Error
}

func (rp *persistenceReportRepository) ReportTopic(reporterId, topicId, reasonId uint) (report.Domain, error) {
	report := TopicReport{
		Reporter: user.User{Model: gorm.Model{ID: reporterId}},
		Topic:    topic.Topic{Model: gorm.Model{ID: topicId}},
		Reason:   ReportReason{ID: reasonId},
	}

	res := rp.Conn.Create(&report)

	return report.toDomain(), res.Error
}

func (rp *persistenceReportRepository) ReportUser(reporterId, suspectId, reasonId uint) (report.Domain, error) {
	report := UserReport{
		Reporter: user.User{Model: gorm.Model{ID: reporterId}},
		Suspect:  user.User{Model: gorm.Model{ID: suspectId}},
		Reason:   ReportReason{ID: reasonId},
	}

	res := rp.Conn.Create(&report)

	return report.toDomain(), res.Error
}

func InitPersistenceReportRepository(c *gorm.DB) report.Repository {
	return &persistenceReportRepository{
		Conn: c,
	}
}
