package reply

import (
	"errors"
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/reply/request"
	"fgd/controllers/reply/response"
	"fgd/core/reply"
	"fgd/helper/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ReplyController struct {
	replyUsecase  reply.Usecase
	storageHelper *storage.StorageHelper
}

func InitReplyController(rc reply.Usecase, sh *storage.StorageHelper) *ReplyController {
	return &ReplyController{
		replyUsecase:  rc,
		storageHelper: sh,
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
		fileName, uploadErr := cr.storageHelper.StoreFile(image)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
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
		fileName, uploadErr := cr.storageHelper.StoreFile(image)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&replyDomain))
}

func (cr *ReplyController) GetReply(c echo.Context) error {
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' in query param")
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'offset' query parameter")
	}
	scope := c.QueryParam("scope")

	if scope == "thread" {
		threadId, err := strconv.Atoi(c.QueryParam("threadId"))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'threadId' in path param")
		}

		domains, err := cr.replyUsecase.GetReplyByThreadID(threadId, limit, offset)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
			}
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&domains))
	} else if scope == "reply" {
		replyId, err := strconv.Atoi(c.QueryParam("replyId"))
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'replyId' in path param")
		}
		domains, err := cr.replyUsecase.GetReplyByThreadID(replyId, limit, offset)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
			}
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&domains))
	} else {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'scope' in query param")
	}
}

func (cr *ReplyController) GetReplyChilds(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'replyId' in path param")
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'limit' in query param")
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required 'offset' query parameter")
	}

	domains, err := cr.replyUsecase.GetReplyChilds(replyId, limit, offset)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&domains))
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
		fileName, uploadErr := cr.storageHelper.StoreFile(image)
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success deleting reply",
	})
}

func (cr *ReplyController) LikeReply(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.replyUsecase.Like(claims.UserID, replyId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success liking reply",
	})
}

func (cr *ReplyController) UndoLikeReply(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.replyUsecase.UndoLike(claims.UserID, replyId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success undoing like from reply",
	})
}

func (cr *ReplyController) UnlikeReply(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.replyUsecase.Unlike(claims.UserID, replyId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success unliking reply",
	})
}

func (cr *ReplyController) UndoUnlikeReply(c echo.Context) error {
	replyId, err := strconv.Atoi(c.Param("replyId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.replyUsecase.UndoUnlike(claims.UserID, replyId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success deleting unlike from reply",
	})
}
