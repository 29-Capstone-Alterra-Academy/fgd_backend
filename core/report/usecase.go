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

func (uc *reportUsecase) ApproveReplyReport(replyReportId uint) error {
	return uc.reportRepository.ApproveReplyReport(replyReportId)
}

func (uc *reportUsecase) ApproveThreadReport(threadReportId uint) error {
	return uc.reportRepository.ApproveThreadReport(threadReportId)
}

func (uc *reportUsecase) ApproveTopicReport(topicReportId uint) error {
	return uc.reportRepository.ApproveTopicReport(topicReportId)
}

func (uc *reportUsecase) ApproveUserReport(userReportId uint) error {
	return uc.reportRepository.ApproveUserReport(userReportId)
}

func (uc *reportUsecase) DeleteReason(reasonId uint) error {
	return uc.reportRepository.DeleteReason(reasonId)
}

func (uc *reportUsecase) ForwardReplyReport(replyReportId uint) error {
	return uc.reportRepository.ForwardReplyReport(replyReportId)
}

func (uc *reportUsecase) ForwardThreadReport(threadReportId uint) error {
	return uc.reportRepository.ForwardThreadReport(threadReportId)
}

func (uc *reportUsecase) IgnoreReplyReport(replyReportId uint) error {
	return uc.reportRepository.RemoveReplyReport(replyReportId)
}

func (uc *reportUsecase) IgnoreThreadReport(threadReportId uint) error {
	return uc.reportRepository.RemoveThreadReport(threadReportId)
}

func (uc *reportUsecase) GetReasons() ([]Domain, error) {
	return uc.reportRepository.GetReasons()
}

func (uc *reportUsecase) RemoveReplyReport(replyReportId uint) error {
	return uc.reportRepository.RemoveReplyReport(replyReportId)
}

func (uc *reportUsecase) RemoveThreadReport(threadReportId uint) error {
	return uc.reportRepository.RemoveThreadReport(threadReportId)
}

func (uc *reportUsecase) RemoveTopicReport(topicReportId uint) error {
	return uc.reportRepository.RemoveTopicReport(topicReportId)
}

func (uc *reportUsecase) RemoveUserReport(userReportId uint) error {
	return uc.reportRepository.RemoveUserReport(userReportId)
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
