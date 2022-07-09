package report

import (
	"fgd/app/config"
	"fgd/helper/format"
)

type reportUsecase struct {
	config           config.Config
	reportRepository Repository
}

func (uc *reportUsecase) GetTopicReplyReports(topicId, limit, offset int) ([]Domain, error) {
	reports, err := uc.reportRepository.GetTopicReplyReports(topicId, limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, report := range reports {
		format.FormatImageLink(
			uc.config,
			report.ReporterProfileImage,
			report.TopicProfileImage,
			report.ReplyImage,
		)
	}

	return reports, nil
}

func (uc *reportUsecase) GetTopicThreadReports(topicId, limit, offset int) ([]Domain, error) {
	reports, err := uc.reportRepository.GetTopicThreadReports(topicId, limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, report := range reports {
		format.FormatImageLink(
			uc.config,
			report.ReporterProfileImage,
			report.TopicProfileImage,
			report.ThreadImage1,
			report.ThreadImage2,
			report.ThreadImage3,
			report.ThreadImage4,
			report.ThreadImage5,
		)
	}

	return reports, nil
}

func (uc *reportUsecase) GetReplyReports(limit, offset int) ([]Domain, error) {
	reports, err := uc.reportRepository.GetReplyReports(limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, report := range reports {
		format.FormatImageLink(
			uc.config,
			report.ReporterProfileImage,
			report.TopicProfileImage,
			report.ReplyImage,
		)
	}

	return reports, nil
}

func (uc *reportUsecase) GetThreadReports(limit, offset int) ([]Domain, error) {
	reports, err := uc.reportRepository.GetThreadReports(limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, report := range reports {
		format.FormatImageLink(
			uc.config,
			report.ReporterProfileImage,
			report.TopicProfileImage,
			report.ThreadImage1,
			report.ThreadImage2,
			report.ThreadImage3,
			report.ThreadImage4,
			report.ThreadImage5,
		)
	}

	return reports, nil
}

func (uc *reportUsecase) GetTopicReports(limit, offset int) ([]Domain, error) {
	reports, err := uc.reportRepository.GetTopicReports(limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, report := range reports {
		format.FormatImageLink(
			uc.config,
			report.ReporterProfileImage,
			report.TopicProfileImage,
		)
	}

	return reports, nil
}

func (uc *reportUsecase) GetUserReports(limit, offset int) ([]Domain, error) {
	reports, err := uc.reportRepository.GetUserReports(limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, report := range reports {
		format.FormatImageLink(
			uc.config,
			report.ReporterProfileImage,
			report.SuspectProfileImage,
		)
	}

	return reports, nil
}

func (uc *reportUsecase) AddReason(data *Domain) error {
	return uc.reportRepository.AddReason(data)
}

func (uc *reportUsecase) ApproveReplyReport(reporterId, replyId uint) error {
	return uc.reportRepository.ApproveReplyReport(reporterId, replyId)
}

func (uc *reportUsecase) ApproveThreadReport(reporterId, threadId uint) error {
	return uc.reportRepository.ApproveThreadReport(reporterId, threadId)
}

func (uc *reportUsecase) ApproveTopicReport(reporterId, topicId uint) error {
	return uc.reportRepository.ApproveTopicReport(reporterId, topicId)
}

func (uc *reportUsecase) ApproveUserReport(reporterId, suspectId uint) error {
	return uc.reportRepository.ApproveUserReport(reporterId, suspectId)
}

func (uc *reportUsecase) DeleteReason(reasonId uint) error {
	return uc.reportRepository.DeleteReason(reasonId)
}

func (uc *reportUsecase) ForwardReplyReport(reporterId, replyId uint) error {
	return uc.reportRepository.ForwardReplyReport(reporterId, replyId)
}

func (uc *reportUsecase) ForwardThreadReport(reporterId, threadId uint) error {
	return uc.reportRepository.ForwardThreadReport(reporterId, threadId)
}

func (uc *reportUsecase) IgnoreReplyReport(reporterId, replyId uint) error {
	return uc.reportRepository.RemoveReplyReport(reporterId, replyId)
}

func (uc *reportUsecase) IgnoreThreadReport(reporterId, threadId uint) error {
	return uc.reportRepository.RemoveThreadReport(reporterId, threadId)
}

func (uc *reportUsecase) GetReasons() ([]Domain, error) {
	return uc.reportRepository.GetReasons()
}

func (uc *reportUsecase) RemoveReplyReport(reporterId, replyId uint) error {
	return uc.reportRepository.RemoveReplyReport(reporterId, replyId)
}

func (uc *reportUsecase) RemoveThreadReport(reporterId, threadId uint) error {
	return uc.reportRepository.RemoveThreadReport(reporterId, threadId)
}

func (uc *reportUsecase) RemoveTopicReport(reporterId, topicId uint) error {
	return uc.reportRepository.RemoveTopicReport(reporterId, topicId)
}

func (uc *reportUsecase) RemoveUserReport(reporterId, suspectId uint) error {
	return uc.reportRepository.RemoveUserReport(reporterId, suspectId)
}

func (uc *reportUsecase) ReportReply(reporterId, replyId, reasonId uint) (Domain, error) {
	report, err := uc.reportRepository.ReportReply(reporterId, replyId, reasonId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(
		uc.config,
		report.ReporterProfileImage,
		report.TopicProfileImage,
		report.ReplyImage,
	)

	return report, nil
}

func (uc *reportUsecase) ReportThread(reporterId, threadId, reasonId uint) (Domain, error) {
	report, err := uc.reportRepository.ReportThread(reporterId, threadId, reasonId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(
		uc.config,
		report.ReporterProfileImage,
		report.TopicProfileImage,
		report.ThreadImage1,
		report.ThreadImage2,
		report.ThreadImage3,
		report.ThreadImage4,
		report.ThreadImage5,
	)

	return report, nil
}

func (uc *reportUsecase) ReportTopic(reporterId, topicId, reasonId uint) (Domain, error) {
	report, err := uc.reportRepository.ReportTopic(reporterId, topicId, reasonId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(
		uc.config,
		report.ReporterProfileImage,
		report.TopicProfileImage,
	)

	return report, nil
}

func (uc *reportUsecase) ReportUser(reporterId, suspectId, reasonId uint) (Domain, error) {
	report, err := uc.reportRepository.ReportUser(reporterId, suspectId, reasonId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(
		uc.config,
		report.ReporterProfileImage,
		report.SuspectProfileImage,
	)

	return report, nil
}

func InitReportUsecase(rp Repository, conf config.Config) Usecase {
	return &reportUsecase{
		reportRepository: rp,
		config:           conf,
	}
}
