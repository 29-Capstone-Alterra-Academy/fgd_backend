package thread

import (
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/thread/request"
	"fgd/controllers/thread/response"
	"fgd/core/thread"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ThreadController struct {
	threadUsecase thread.Usecase
}

func InitThreadController(tc thread.Usecase) *ThreadController {
	return &ThreadController{
		threadUsecase: tc,
	}
}

func (cr *ThreadController) Create(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	thread := request.Thread{}
	err = c.Bind(&thread)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	newThread, err := cr.threadUsecase.CreateThread(thread.ToDomain(), claims.UserID, topicId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&newThread))
}

func (cr *ThreadController) Update(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	thread := request.Thread{}
	err = c.Bind(&thread)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	updatedThread, err := cr.threadUsecase.UpdateThread(thread.ToDomain(), claims.UserID, threadId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(&updatedThread))
}

func (cr *ThreadController) Delete(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.DeleteThread(claims.UserID, threadId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ThreadController) Like(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.Like(claims.UserID, threadId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ThreadController) UndoLike(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.UndoLike(claims.UserID, threadId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ThreadController) Unlike(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.Unlike(claims.UserID, threadId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ThreadController) UndoUnlike(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.UndoUnlike(claims.UserID, threadId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}
