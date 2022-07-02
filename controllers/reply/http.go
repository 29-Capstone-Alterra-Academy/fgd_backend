package reply

import (
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/reply/request"
	"fgd/controllers/reply/response"
	"fgd/core/reply"
	"fgd/helper/storage"
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

	var fileName string

	content := c.FormValue("content")
	image, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	newReply := request.Reply{}

	if err != http.ErrMissingFile {
		fileName, uploadErr := storage.StoreFile(image)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		newReply.Image = fileName
	} else {
		newReply.Image = fileName
	}

	newReply.Content = content

	replyDomain, err := cr.replyUsecase.CreateReplyReply(newReply.ToDomain(), claims.UserID, replyId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&replyDomain))
}

func (cr *ReplyController) CreateForThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	var fileName string

	content := c.FormValue("content")
	image, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	newReply := request.Reply{}

	if err != http.ErrMissingFile {
		fileName, uploadErr := storage.StoreFile(image)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		newReply.Image = fileName
	} else {
		newReply.Image = fileName
	}

	newReply.Content = content

	replyDomain, err := cr.replyUsecase.CreateReplyThread(newReply.ToDomain(), claims.UserID, threadId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&replyDomain))
}

func (cr *ReplyController) UpdateReply(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	var fileName string

	content := c.FormValue("content")
	image, err := c.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	updatedReply := request.Reply{}

	if err != http.ErrMissingFile {
		fileName, uploadErr := storage.StoreFile(image)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		updatedReply.Image = fileName
	} else {
		updatedReply.Image = fileName
	}

	updatedReply.Content = content

	replyDomain, err := cr.replyUsecase.EditReply(updatedReply.ToDomain(), claims.UserID, replyId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&replyDomain))
}

func (cr *ReplyController) DeleteReply(c echo.Context) error {
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

func (cr *ReplyController) LikeReply(c echo.Context) error {
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

func (cr *ReplyController) UndoLikeReply(c echo.Context) error {
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

func (cr *ReplyController) UnlikeReply(c echo.Context) error {
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

func (cr *ReplyController) UndoUnlikeReply(c echo.Context) error {
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
