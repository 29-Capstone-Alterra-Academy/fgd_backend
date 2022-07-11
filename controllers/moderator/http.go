package moderator

import (
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/moderator/response"
	"fgd/core/moderator"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ModeratorController struct {
	moderatorUsecase moderator.Usecase
}

func InitModeratorController(mc moderator.Usecase) *ModeratorController {
	return &ModeratorController{
		moderatorUsecase: mc,
	}
}

func (cr *ModeratorController) RequestPromotion(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'topicId' path param")
	}

	domain, err := cr.moderatorUsecase.ApplyPromotion(uint(claims.UserID), uint(topicId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(domain))
}

func (cr *ModeratorController) ActionPromotion(c echo.Context) error {
	promotionId, err := strconv.Atoi(c.QueryParam("promotionId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'promotionId' query param")
	}

	action := c.QueryParam("action")

	if action == "approve" {
		err := cr.moderatorUsecase.ApprovePromotion(uint(promotionId))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.SuccessResponse(c, http.StatusOK, nil)
	} else if action == "reject" {
		err := cr.moderatorUsecase.RejectPromotion(uint(promotionId))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.SuccessResponse(c, http.StatusOK, nil)
	} else {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'action' query param")
	}
}

func (cr *ModeratorController) GetPromotionRequests(c echo.Context) error {
	domains, err := cr.moderatorUsecase.GetPromotionRequest()
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}
	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&domains))
}

func (cr *ModeratorController) StepDown(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	topicId, err := strconv.Atoi(c.QueryParam("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'topicId' query param")
	}

	err = cr.moderatorUsecase.StepDown(uint(claims.UserID), uint(topicId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ModeratorController) RemoveModerator(c echo.Context) error {
	userId, err := strconv.Atoi(c.QueryParam("userId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'userId' query param")
	}

	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'topicId' query param")
	}

	err = cr.moderatorUsecase.StepDown(uint(userId), uint(topicId))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}
