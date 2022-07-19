package topic

import (
	"errors"
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/topic/request"
	"fgd/controllers/topic/response"
	"fgd/core/auth"
	"fgd/core/topic"
	"fgd/core/user"
	"fgd/helper/storage"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TopicController struct {
	authUsecase   auth.Usecase
	topicUsecase  topic.Usecase
	userUsecase   user.Usecase
	storageHelper *storage.StorageHelper
}

func InitTopicController(ac auth.Usecase, tc topic.Usecase, uc user.Usecase, sh *storage.StorageHelper) *TopicController {
	return &TopicController{
		authUsecase:   ac,
		topicUsecase:  tc,
		userUsecase:   uc,
		storageHelper: sh,
	}
}

func (cr *TopicController) CreateTopic(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	var fileName string

	topicName := c.FormValue("name")
	description := c.FormValue("description")
	rules := c.FormValue("rules")
	profileImage, err := c.FormFile("profile_image")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	newTopic := request.NewTopic{}
	if err != http.ErrMissingFile {
		fileName, uploadErr := cr.storageHelper.StoreFile(profileImage)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		newTopic.ProfileImage = fileName
	} else {
		newTopic.ProfileImage = fileName
	}

	newTopic.Name = topicName
	newTopic.Description = description
	newTopic.Rules = rules

	topicDomain, err := cr.topicUsecase.CreateTopic(newTopic.ToDomain(), claims.UserID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(topicDomain))
}

func (cr *TopicController) UpdateTopic(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	var fileName string

	topicName := c.FormValue("name")
	description := c.FormValue("description")
	rules := c.FormValue("rules")
	profileImage, err := c.FormFile("profile_image")
	if err != nil && err != http.ErrMissingFile {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	newTopic := request.NewTopic{}
	if err != http.ErrMissingFile {
		fileName, uploadErr := cr.storageHelper.StoreFile(profileImage)
		if uploadErr != nil {
			return controllers.FailureResponse(c, http.StatusUnprocessableEntity, uploadErr.Error())
		}
		newTopic.ProfileImage = fileName
	} else {
		newTopic.ProfileImage = fileName
	}

	newTopic.Name = topicName
	newTopic.Description = description
	newTopic.Rules = rules

	topicDomain, err := cr.topicUsecase.UpdateTopic(newTopic.ToDomain(), topicId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, response.FromDomain(topicDomain))
}

func (cr *TopicController) CheckAvailibility(c echo.Context) error {
	topicName := c.QueryParam("topicname")

	exist := cr.topicUsecase.CheckTopicAvailibility(topicName)

	if !exist {
		return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
			"message": "Topic name is available to use",
		})
	} else {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: topic already exist")
	}
}

func (cr *TopicController) GetModerators(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	userDomains, err := cr.userUsecase.GetModerators(topicId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromUserDomains(&userDomains))
}

func (cr *TopicController) GetTopics(c echo.Context) error {
	userId, err := strconv.Atoi(c.QueryParam("userId"))
	if err == nil {
		topicDomains, err := cr.topicUsecase.GetSubscribedTopics(userId)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&topicDomains))
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	sort_by := c.QueryParam("sort_by")

	topicDomains, err := cr.topicUsecase.GetTopics(limit, offset, sort_by)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomains(&topicDomains))
}

func (cr *TopicController) GetTopicDetails(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	topicDetailDomain, err := cr.topicUsecase.GetTopicDetails(topicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(topicDetailDomain))
}

func (cr *TopicController) Subscribe(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "Error getting 'userId' path parameter")
	}

	claims := middleware.ExtractUserClaims(c)
	userId := claims.UserID

	err = cr.topicUsecase.Subscribe(userId, topicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success subscribing to topic",
	})
}

func (cr *TopicController) Unsubscribe(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "Error getting 'userId' path parameter")
	}

	claims := middleware.ExtractUserClaims(c)
	userId := claims.UserID

	err = cr.topicUsecase.Unsubscribe(userId, topicId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return controllers.FailureResponse(c, http.StatusNotFound, err.Error())
		}
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success unsubscribing to topic",
	})
}
