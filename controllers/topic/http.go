package topic

import (
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/topic/request"
	"fgd/controllers/topic/response"
	"fgd/core/auth"
	"fgd/core/topic"
	"fgd/core/user"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TopicController struct {
	authUsecase  auth.Usecase
	topicUsecase topic.Usecase
	userUsecase  user.Usecase
}

func InitTopicController(ac auth.Usecase, tc topic.Usecase, uc user.Usecase) *TopicController {
	return &TopicController{
		authUsecase:  ac,
		topicUsecase: tc,
		userUsecase:  uc,
	}
}

func (cr *TopicController) CreateTopic(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)
	userId := claims.UserID

	newTopic := request.NewTopic{}
	err := c.Bind(&newTopic)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	topicDomain, err := cr.topicUsecase.CreateTopic(newTopic.ToDomain())
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = cr.topicUsecase.Subscribe(userId, topicDomain.ID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, topicDomain)
}

func (cr *TopicController) CheckAvailibility(c echo.Context) error {
	topicName := c.QueryParam("topicname")

	available, err := cr.topicUsecase.CheckTopicAvailibility(topicName)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !available {
		return controllers.FailureResponse(c, http.StatusBadRequest, "topic is already exist")
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

// func (cr *TopicController) GetModerators(c echo.Context) error {}

func (cr *TopicController) RequestPromotion(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("topicId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "Error getting 'userId' path parameter")
	}

	claims := middleware.ExtractUserClaims(c)
	userId := claims.UserID

	// TODO Should be promote
	err = cr.topicUsecase.Subscribe(userId, topicId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, nil)
}

func (cr *TopicController) GetTopics(c echo.Context) error {
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
	topics := []response.Topic{}
	for _, topicDomain := range topicDomains {
		topics = append(topics, response.FromDomain(topicDomain))
	}

	return controllers.SuccessResponse(c, http.StatusOK, topics)
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
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
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
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}
