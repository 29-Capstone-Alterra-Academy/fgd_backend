package user

import (
	"fgd/app/config"
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/user/request"
	"fgd/controllers/user/response"
	"fgd/core/auth"
	"fgd/core/user"
	"fgd/core/verify"
	"fgd/helper/storage"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	config        config.Config
	authUsecase   auth.Usecase
	userUsecase   user.Usecase
	verifyUsecase verify.Usecase
	storageHelper *storage.StorageHelper
}

func InitUserController(ac auth.Usecase, uc user.Usecase, vc verify.Usecase, conf config.Config, sh *storage.StorageHelper) *UserController {
	return &UserController{
		config:        conf,
		authUsecase:   ac,
		userUsecase:   uc,
		verifyUsecase: vc,
		storageHelper: sh,
	}
}

func (cr *UserController) Login(c echo.Context) error {
	user := request.UserAuth{}
	if err := c.Bind(&user); err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := cr.userUsecase.CreateToken(user.Username, user.Email, user.Password)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}

func (cr *UserController) Logout(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	err := cr.authUsecase.DeleteAuth(claims.UserID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusUnprocessableEntity, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *UserController) Register(c echo.Context) error {
	user := request.UserAuth{}
	err := c.Bind(&user)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	userDomain, err := cr.userUsecase.CreateUser(user.ToDomain())
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = cr.verifyUsecase.SendVerifyToken(userDomain.Email, "EMAIL_VERIFY")
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusCreated, nil)
}

func (cr *UserController) RefreshToken(c echo.Context) error {
	tokenReq := request.TokenRequest{}
	c.Bind(&tokenReq)

	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return middleware.CustomToken{}, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(cr.config.JWT_SECRET), nil
	})
	if err != nil {
		return controllers.FailureResponse(c, http.StatusUnauthorized, err.Error())
	}

	var newToken middleware.CustomToken
	if claims, ok := token.Claims.(middleware.JWTCustomClaims); ok && token.Valid {

		user, err := cr.userUsecase.GetPersonalProfile(claims.UserID)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		newToken, err = cr.userUsecase.CreateToken(user.Username, user.Email, user.Password)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}
	}

	return controllers.SuccessResponse(c, http.StatusOK, newToken)
}

func (cr *UserController) CheckAvailibility(c echo.Context) error {
	username := c.QueryParam("username")

	exist := cr.userUsecase.CheckUserAvailibility(username)

	if !exist {
		return controllers.SuccessResponse(c, http.StatusOK, nil)
	} else {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: username already used")
	}
}

func (cr *UserController) Follow(c echo.Context) error {
	targetId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "Error getting 'userId' path parameter")
	}

	claims := middleware.ExtractUserClaims(c)
	userId := claims.UserID

	err = cr.userUsecase.FollowUser(userId, targetId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "Error following user")
	}

	return c.NoContent(http.StatusOK)
}

func (cr *UserController) Unfollow(c echo.Context) error {
	targetId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "Error getting 'userId' path parameter")
	}

	user := c.Get("user").(*jwt.Token)
	userClaims := user.Claims.(*middleware.JWTCustomClaims)
	userId := userClaims.UserID

	err = cr.userUsecase.FollowUser(userId, targetId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "Error unfollowing user")
	}

	return c.NoContent(http.StatusOK)
}

func (cr *UserController) GetProfile(c echo.Context) error {
	userClaims := middleware.ExtractUserClaims(c)
	profile, err := cr.userUsecase.GetPersonalProfile(userClaims.UserID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(&profile, "personal"))
}

func (cr *UserController) GetPublicProfile(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	profile, err := cr.userUsecase.GetProfileByID(userId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(&profile, "public"))
}

func (cr *UserController) UpdateProfile(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	var birthDate time.Time
	var err error

	username := c.FormValue("username")
	email := c.FormValue("email")
	bio := c.FormValue("bio")
	gender := c.FormValue("gender")
	birthDateStr := c.FormValue("birth_date")

	if birthDateStr != "" {
		birthDate, err = time.Parse("2006-01-02", birthDateStr)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
		}
	}

	profileImage, err := c.FormFile("profile_image")
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	var profileImageFilename string
	if profileImage != nil {
		profileImageFilename, err = cr.storageHelper.StoreFile(profileImage)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
		}
	}

	user := request.UserProfile{
		Username:     username,
		Email:        email,
		Gender:       &gender,
		Bio:          &bio,
		BirthDate:    &birthDate,
		ProfileImage: &profileImageFilename,
	}

	updatedUser, err := cr.userUsecase.UpdatePersonalProfile(user.ToDomain(), claims.UserID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, response.FromDomain(&updatedUser, "personal"))
}
