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
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	config        config.Config
	jwtConfig     middleware.JWTConfig
	authUsecase   auth.Usecase
	userUsecase   user.Usecase
	verifyUsecase verify.Usecase
	storageHelper *storage.StorageHelper
}

func InitUserController(ac auth.Usecase, uc user.Usecase, vc verify.Usecase, conf config.Config, jConf middleware.JWTConfig, sh *storage.StorageHelper) *UserController {
	return &UserController{
		config:        conf,
		jwtConfig:     jConf,
		authUsecase:   ac,
		userUsecase:   uc,
		verifyUsecase: vc,
		storageHelper: sh,
	}
}

func (cr *UserController) Login(c echo.Context) error {
	user := request.UserLogin{}
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

	err := cr.authUsecase.DeleteAuth(claims.UUID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusUnprocessableEntity, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Succesfully logged out",
	})
}

func (cr *UserController) Register(c echo.Context) error {
	user := request.UserRegister{}
	err := c.Bind(&user)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	err = c.Validate(&user)
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

	return controllers.SuccessResponse(c, http.StatusCreated, map[string]interface{}{
		"message": "Success registering user",
	})
}

func (cr *UserController) RefreshToken(c echo.Context) error {
	tokenReq := request.TokenRequest{}
	c.Bind(&tokenReq)

	err := c.Validate(&tokenReq)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	customClaims := middleware.JWTCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenReq.RefreshToken, &customClaims, cr.jwtConfig.CustomKeyFunc)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusUnauthorized, err.Error())
	}

	var newToken middleware.CustomToken
	claims, ok := token.Claims.(*middleware.JWTCustomClaims)
	if ok && token.Valid {
		cr.authUsecase.DeleteAuth(claims.UUID)
		user, err := cr.userUsecase.GetPersonalProfile(claims.UserID)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}

		newToken, err = cr.userUsecase.CreateToken(user.Username, user.Email, user.Password)
		if err != nil {
			return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
		}
		return controllers.SuccessResponse(c, http.StatusOK, newToken)
	} else {
		return controllers.FailureResponse(c, http.StatusInternalServerError, "error: failed to extract token claims")
	}
}

func (cr *UserController) CheckAvailibility(c echo.Context) error {
	username := c.QueryParam("username")
	email := c.QueryParam("email")

	var exist bool
	if username != "" {
		exist = cr.userUsecase.CheckUserAvailibility(username)
	} else if email != "" {
		exist = cr.userUsecase.CheckEmailAvailibility(email)
	} else {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: missing required query param of either 'username' or 'email'")
	}

	if !exist {
		return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
			"message": "username/email is available to use",
		})
	} else {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error: username/email already used")
	}
}

func (cr *UserController) Follow(c echo.Context) error {
	targetId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error getting 'userId' path parameter")
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.userUsecase.FollowUser(claims.UserID, targetId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error following user")
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success following user",
	})
}

func (cr *UserController) Unfollow(c echo.Context) error {
	targetId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error getting 'userId' path parameter")
	}

	claims := middleware.ExtractUserClaims(c)

	err = cr.userUsecase.UnfollowUser(claims.UserID, targetId)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "error unfollowing user")
	}

	return controllers.SuccessResponse(c, http.StatusOK, map[string]interface{}{
		"message": "Success unfollowing user",
	})
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

	profileImage, _ := c.FormFile("profile_image")
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
