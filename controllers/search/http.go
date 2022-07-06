package search

import (
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/search/request"
	"fgd/controllers/search/response"
	"fgd/core/search"
	"fgd/core/thread"
	"fgd/core/topic"
	"fgd/core/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type SearchController struct {
	searchUsecase search.Usecase
	threadUsecase thread.Usecase
	topicUsecase  topic.Usecase
	userUsecase   user.Usecase
}

func InitSearchController(sc search.Usecase, tc thread.Usecase, pc topic.Usecase, uc user.Usecase) *SearchController {
	return &SearchController{
		searchUsecase: sc,
		threadUsecase: tc,
		topicUsecase:  nil,
		userUsecase:   uc,
	}
}

func (cr *SearchController) Search(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' query parameter")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'offset' query parameter")
	}
	keyword := c.QueryParam("keyword")
	if keyword == "" {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'keyword' query parameter")
	}
	scope := c.QueryParam("scope")

	if scope != "" {
		if scope == "user" {
			domains, err := cr.userUsecase.GetUsersByKeyword(keyword, limit, offset)
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, response.FromUserDomains(&domains))
		} else if scope == "topic" {
			domains, err := cr.topicUsecase.GetTopicsByKeyword(keyword, limit, offset)
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, response.FromTopicDomains(&domains))
		} else if scope == "thread" {
			domains, err := cr.threadUsecase.GetThreadByKeyword(keyword, limit, offset)
			if err != nil {
				return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
			}

			return controllers.SuccessResponse(c, http.StatusOK, response.FromThreadDomains(&domains))
		}
	}

	return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing or wrong required 'scope' query param")
}

func (cr *SearchController) StoreSearchHistory(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	keyword := c.QueryParam("keyword")
	if keyword == "" {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'keyword' query parameter")
	}

	query := request.Query{}
	query.UserID = uint(claims.UserID)
	query.Keyword = keyword

	err := cr.searchUsecase.StoreSearchKeyword(uint(claims.UserID), query.ToDomain())
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *SearchController) GetSearchHistory(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' query parameter")
	}

	keyword := c.QueryParam("keyword")

	domains, err := cr.searchUsecase.GetSearchHistory(uint(claims.UserID), keyword, limit)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromSearchDomains(&domains))
}
