package user

import (
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/user/request"
	"fgd/controllers/user/response"
	"fgd/core/auth"
	"fgd/core/user"
	"fgd/core/verify"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	authUsecase   auth.Usecase
	userUsecase   user.Usecase
	verifyUsecase verify.Usecase
}

func InitUserController(ac auth.Usecase, uc user.Usecase, vc verify.Usecase) *UserController {
	return &UserController{
		authUsecase:   ac,
		userUsecase:   uc,
		verifyUsecase: vc,
	}
}

func (cr *UserController) Login(c echo.Context) error {
	user := request.User{}
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
	user := request.User{}
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
		return []byte("secret"), nil
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

	exist, err := cr.userUsecase.CheckUserAvailibility(username)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, "Error checking username availibility: "+err.Error())
	}

	if !exist {
		return controllers.SuccessResponse(c, http.StatusOK, nil)
	} else {
		// TODO Is this right ?
		return controllers.FailureResponse(c, http.StatusBadRequest, "")
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

func (cr *UserController) UpdateProfileImage(c echo.Context) error {
	file, err := c.FormFile("profileImage")
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}
