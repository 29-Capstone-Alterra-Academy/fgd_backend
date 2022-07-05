package report

import (
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/report/request"
	"fgd/controllers/report/response"
	"fgd/core/report"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReportController struct {
	reportUsecase report.Usecase
}

func InitReportController(rc report.Usecase) *ReportController {
	return &ReportController{
		reportUsecase: rc,
	}
}

func (cr *ReportController) ReportUser(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	suspectId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing userId param")
	}

	reasonId, err := strconv.Atoi(c.QueryParam("reasonId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing reasonId query param")
	}

	reportDomain, err := cr.reportUsecase.ReportUser(uint(claims.UserID), uint(suspectId), uint(reasonId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&reportDomain, "user"))
}

func (cr *ReportController) ReportTopic(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing topicId param")
	}

	reasonId, err := strconv.Atoi(c.QueryParam("reasonId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing reasonId query param")
	}

	reportDomain, err := cr.reportUsecase.ReportTopic(uint(claims.UserID), uint(topicId), uint(reasonId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&reportDomain, "topic"))
}

func (cr *ReportController) ReportThread(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing threadId param")
	}

	reasonId, err := strconv.Atoi(c.QueryParam("reasonId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing reasonId query param")
	}

	reportDomain, err := cr.reportUsecase.ReportThread(uint(claims.UserID), uint(threadId), uint(reasonId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&reportDomain, "thread"))
}

func (cr *ReportController) ReportReply(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing replyId param")
	}

	reasonId, err := strconv.Atoi(c.QueryParam("reasonId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing reasonId query param")
	}

	reportDomain, err := cr.reportUsecase.ReportReply(uint(claims.UserID), uint(replyId), uint(reasonId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&reportDomain, "reply"))
}

func (cr *ReportController) GetUserReports(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' variable on query param")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'offset' variable on query param")
	}

	reportDomains, err := cr.reportUsecase.GetUserReports(limit, offset)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&reportDomains, "user"))
}

func (cr *ReportController) GetTopicReports(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' variable on query param")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'offset' variable on query param")
	}

	reportDomains, err := cr.reportUsecase.GetTopicReports(limit, offset)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&reportDomains, "topic"))
}

func (cr *ReportController) GetThreadReports(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' variable on query param")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'offset' variable on query param")
	}

	reportDomains, err := cr.reportUsecase.GetThreadReports(limit, offset)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&reportDomains, "thread"))
}

func (cr *ReportController) GetReplyReports(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' variable on query param")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'offset' variable on query param")
	}

	reportDomains, err := cr.reportUsecase.GetReplyReports(limit, offset)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&reportDomains, "reply"))
}

func (cr *ReportController) ApproveUserReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.ApproveUserReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) ApproveTopicReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.ApproveTopicReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) ApproveThreadReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.ApproveThreadReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) ApproveReplyReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.ApproveReplyReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) ForwardThreadReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.ForwardThreadReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) ForwardReplyReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.ForwardReplyReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) IgnoreThreadReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.IgnoreThreadReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) IgnoreReplyReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.IgnoreReplyReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) RemoveUserReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.RemoveUserReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) RemoveTopicReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.RemoveTopicReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) RemoveThreadReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.RemoveThreadReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) RemoveReplyReport(c echo.Context) error {
	reportId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'id' path parameter")
	}

	err = cr.reportUsecase.RemoveReplyReport(uint(reportId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) AddReason(c echo.Context) error {
	newReason := request.Reason{}

	err := c.Bind(&newReason)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	err = cr.reportUsecase.AddReason(newReason.ToDomain())
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) GetReasons(c echo.Context) error {
	reasonDomains, err := cr.reportUsecase.GetReasons()
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&reasonDomains, "reason"))
}

func (cr *ReportController) DeleteReason(c echo.Context) error {
	reasonId, err := strconv.Atoi(c.Param("reasonId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'reasonId' path parameter")
	}

	err = cr.reportUsecase.DeleteReason(uint(reasonId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReportController) GetTopicScopeReports(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'topicId' in path param")
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' in path param")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'offset' in path param")
	}

	scope := c.QueryParam("scope")
	if scope == "" {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'scope' in query param")
	}

	if scope == "thread" {
		threadDomains, err := cr.reportUsecase.GetTopicThreadReports(topicId, limit, offset)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&threadDomains, "thread"))
	} else if scope == "reply" {
		replyDomains, err := cr.reportUsecase.GetTopicReplyReports(topicId, limit, offset)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&replyDomains, "reply"))
	}
	return controllers.FailureResponse(c, http.StatusBadRequest, "")
}
