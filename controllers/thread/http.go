package thread

import (
	"errors"
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/thread/request"
	"fgd/controllers/thread/response"
	"fgd/core/thread"
	"fgd/helper/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ThreadController struct {
	threadUsecase thread.Usecase
	storageHelper *storage.StorageHelper
}

func InitThreadController(tc thread.Usecase, sh *storage.StorageHelper) *ThreadController {
	return &ThreadController{
		threadUsecase: tc,
		storageHelper: sh,
	}
}

func (cr *ThreadController) CreateThread(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	var image1 string
	var image2 string
	var image3 string
	var image4 string
	var image5 string

	title := c.FormValue("title")
	content := c.FormValue("content")

	newThread := request.Thread{}

	newThread.Title = title
	newThread.Content = content

	imageFile1, err := c.FormFile("image_1")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image1, uploadErr := cr.storageHelper.StoreFile(imageFile1)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		newThread.Image1 = image1
	} else {
		newThread.Image1 = image1
	}

	imageFile2, err := c.FormFile("image_2")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image2, uploadErr := cr.storageHelper.StoreFile(imageFile2)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		newThread.Image2 = image2
	} else {
		newThread.Image2 = image2
	}

	imageFile3, err := c.FormFile("image_3")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image3, uploadErr := cr.storageHelper.StoreFile(imageFile3)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		newThread.Image3 = image3
	} else {
		newThread.Image3 = image3
	}

	imageFile4, err := c.FormFile("image_4")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image4, uploadErr := cr.storageHelper.StoreFile(imageFile4)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		newThread.Image4 = image4
	} else {
		newThread.Image4 = image4
	}

	imageFile5, err := c.FormFile("image_5")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image5, uploadErr := cr.storageHelper.StoreFile(imageFile5)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		newThread.Image5 = image5
	} else {
		newThread.Image5 = image5
	}

	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	threadDomain, err := cr.threadUsecase.CreateThread(newThread.ToDomain(), claims.UserID, topicId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(&threadDomain))
}

func (cr *ThreadController) UpdateThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	var image1 string
	var image2 string
	var image3 string
	var image4 string
	var image5 string

	title := c.FormValue("title")
	content := c.FormValue("content")

	updatedThread := request.Thread{}

	updatedThread.Title = title
	updatedThread.Content = content

	imageFile1, err := c.FormFile("image_1")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image1, uploadErr := cr.storageHelper.StoreFile(imageFile1)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		updatedThread.Image1 = image1
	} else {
		updatedThread.Image1 = image1
	}

	imageFile2, err := c.FormFile("image_2")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image2, uploadErr := cr.storageHelper.StoreFile(imageFile2)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		updatedThread.Image2 = image2
	} else {
		updatedThread.Image2 = image2
	}

	imageFile3, err := c.FormFile("image_3")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image3, uploadErr := cr.storageHelper.StoreFile(imageFile3)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		updatedThread.Image3 = image3
	} else {
		updatedThread.Image3 = image3
	}

	imageFile4, err := c.FormFile("image_4")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image4, uploadErr := cr.storageHelper.StoreFile(imageFile4)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		updatedThread.Image4 = image4
	} else {
		updatedThread.Image4 = image4
	}

	imageFile5, err := c.FormFile("image_5")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	if err != http.ErrMissingFile {
		image5, uploadErr := cr.storageHelper.StoreFile(imageFile5)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		updatedThread.Image5 = image5
	} else {
		updatedThread.Image5 = image5
	}

	err = c.Bind(&updatedThread)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	threadDomain, err := cr.threadUsecase.UpdateThread(updatedThread.ToDomain(), claims.UserID, threadId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(&threadDomain))
}

func (cr *ThreadController) DeleteThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.DeleteThread(claims.UserID, threadId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success deleting thread",
	})
}

func (cr *ThreadController) GetThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	threadDomain, err := cr.threadUsecase.GetThreadByID(threadId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(&threadDomain))
}

func (cr *ThreadController) GetThreads(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))

	userIdStr := c.QueryParam("userId")
	topicIdStr := c.QueryParam("topicId")

	if userIdStr != "" {
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
		}

		threadDomain, err := cr.threadUsecase.GetThreadByAuthorID(userId, limit, offset)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&threadDomain))
	} else if topicIdStr != "" {
		topicId, err := strconv.Atoi(topicIdStr)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
		}

		threadDomain, err := cr.threadUsecase.GetThreadByTopicID(topicId, limit, offset)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&threadDomain))
	}

	return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required query parameter")
}

func (cr *ThreadController) LikeThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.Like(claims.UserID, threadId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success liking thread",
	})
}

func (cr *ThreadController) UndoLikeThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.UndoLike(claims.UserID, threadId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success undoing like from thread",
	})
}

func (cr *ThreadController) UnlikeThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.Unlike(claims.UserID, threadId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success unliking thread",
	})
}

func (cr *ThreadController) UndoUnlikeThread(c echo.Context) error {
	threadId, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.threadUsecase.UndoUnlike(claims.UserID, threadId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success undoing unlike from thread",
	})
}
