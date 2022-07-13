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
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing 'userId' param")
	}

	reasonId, err := strconv.Atoi(c.QueryParam("reasonId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing 'reasonId' query param")
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

	return controllers.SuccessResponse(c, http.StatusCreated, nil)
}

func (cr *ReportController) GetReasons(c echo.Context) error {
	reasonDomains, err := cr.reportUsecase.GetReasons()
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomainsReason(&reasonDomains))
}

func (cr *ReportController) DeleteReason(c echo.Context) error {
	reasonId, err := strconv.Atoi(c.QueryParam("reasonId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'reasonId' path parameter")
	}

	err = cr.reportUsecase.DeleteReason(uint(reasonId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success deleting report reason",
	})
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

func (cr *ReportController) ActionTopicScopeReport(c echo.Context) error {
	scope := c.QueryParam("scope")
	if scope == "" {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'scope' in query param")
	}

	reporterId, err := strconv.Atoi(c.QueryParam("reporterId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'reporterId' query param")
	}

	action := c.QueryParam("action")
	if action == "" {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'action' in query param")
	}

	if scope == "thread" {
		threadId, err := strconv.Atoi(c.QueryParam("threadId"))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing 'threadId' query param")
		}
		if action == "forward" {
			err = cr.reportUsecase.ForwardThreadReport(uint(reporterId), uint(threadId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success forwarding thread report to admin",
			})

		} else if action == "ignore" {
			err = cr.reportUsecase.IgnoreThreadReport(uint(reporterId), uint(threadId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success deleting thread report",
			})
		}

	} else if scope == "reply" {
		replyId, err := strconv.Atoi(c.QueryParam("replyId"))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing 'replyId' query param")
		}
		if action == "forward" {
			err = cr.reportUsecase.ForwardReplyReport(uint(reporterId), uint(replyId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success forwarding reply report to admin",
			})

		} else if action == "ignore" {
			err = cr.reportUsecase.IgnoreReplyReport(uint(reporterId), uint(replyId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success deleting reply report",
			})

		}
	}
	return controllers.FailureResponse(c, http.StatusBadRequest, "")
}

func (cr *ReportController) GetReports(c echo.Context) error {
	scope := c.QueryParam("scope")
	if scope == "" {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'scope' query param")
	}
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' in query param")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'offset' in query param")
	}

	if scope == "user" {
		reportDomains, err := cr.reportUsecase.GetUserReports(limit, offset)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&reportDomains, "user"))
	} else if scope == "topic" {
		reportDomains, err := cr.reportUsecase.GetTopicReports(limit, offset)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&reportDomains, "topic"))
	} else if scope == "thread" {
		reportDomains, err := cr.reportUsecase.GetThreadReports(limit, offset)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&reportDomains, "thread"))
	} else if scope == "reply" {
		reportDomains, err := cr.reportUsecase.GetReplyReports(limit, offset)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&reportDomains, "reply"))
	}

	return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing some requirement in request")
}

func (cr *ReportController) ActionReport(c echo.Context) error {
	scope := c.QueryParam("scope")
	if scope == "" {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'scope' query param")
	}

	action := c.QueryParam("action")
	if action == "" {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'action' query param")
	}

	reporterId, err := strconv.Atoi(c.QueryParam("reporterId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'reporterId' query param")
	}

	if scope == "user" {
		suspectId, err := strconv.Atoi(c.QueryParam("suspectId"))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing 'suspectId' query param")
		}
		if action == "approve" {
			err = cr.reportUsecase.ApproveUserReport(uint(reporterId), uint(suspectId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success approving user report",
			})

		} else if action == "remove" {
			err = cr.reportUsecase.RemoveUserReport(uint(reporterId), uint(suspectId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success removing user report",
			})

		}
	} else if scope == "topic" {
		topicId, err := strconv.Atoi(c.QueryParam("topicId"))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing 'topicId' query param")
		}
		if action == "approve" {
			err = cr.reportUsecase.ApproveTopicReport(uint(reporterId), uint(topicId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success approving topic report",
			})

		} else if action == "remove" {
			err = cr.reportUsecase.RemoveTopicReport(uint(reporterId), uint(topicId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success removing topic report",
			})

		}
	} else if scope == "thread" {
		threadId, err := strconv.Atoi(c.QueryParam("threadId"))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing 'threadId' query param")
		}
		if action == "approve" {
			err = cr.reportUsecase.ApproveThreadReport(uint(reporterId), uint(threadId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success approving thread report",
			})

		} else if action == "remove" {
			err = cr.reportUsecase.RemoveThreadReport(uint(reporterId), uint(threadId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success removing thread report",
			})

		}
	} else if scope == "reply" {
		replyId, err := strconv.Atoi(c.QueryParam("replyId"))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing 'replyId' query param")
		}
		if action == "approve" {
			err = cr.reportUsecase.ApproveReplyReport(uint(reporterId), uint(replyId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success approving reply report",
			})

		} else if action == "remove" {
			err = cr.reportUsecase.RemoveReplyReport(uint(reporterId), uint(replyId))
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
				"message": "Success removing reply report",
			})

		}
	}

	return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing some requirement in request")
}
