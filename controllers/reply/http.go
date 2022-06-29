package reply

import (
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/reply/request"
	"fgd/controllers/reply/response"
	"fgd/core/reply"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ReplyController struct {
	replyUsecase reply.Usecase
}

func InitReplyController(rc reply.Usecase) *ReplyController {
	return &ReplyController{
		replyUsecase: rc,
	}
}

func (cr *ReplyController) CreateForReply(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	reply := request.Reply{}
	err = c.Bind(&reply)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	newReply, err := cr.replyUsecase.CreateReplyReply(reply.ToDomain(), claims.UserID, replyId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&newReply))
}

func (cr *ReplyController) CreateForThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	reply := request.Reply{}
	err = c.Bind(&reply)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	newReply, err := cr.replyUsecase.CreateReplyThread(reply.ToDomain(), claims.UserID, threadId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&newReply))
}

func (cr *ReplyController) Edit(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	reply := request.Reply{}
	err = c.Bind(&reply)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	updatedReply, err := cr.replyUsecase.EditReply(reply.ToDomain(), claims.UserID, replyId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&updatedReply))
}

func (cr *ReplyController) Delete(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.replyUsecase.DeleteReply(claims.UserID, replyId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReplyController) Like(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.replyUsecase.Like(claims.UserID, replyId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReplyController) UndoLike(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.replyUsecase.UndoLike(claims.UserID, replyId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReplyController) Unlike(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.replyUsecase.Unlike(claims.UserID, replyId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *ReplyController) UndoUnlike(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.replyUsecase.UndoUnlike(claims.UserID, replyId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}
