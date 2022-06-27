package verify

import (
	"fgd/app/middleware"
	"fgd/controllers"
	"fgd/controllers/verify/request"
	"fgd/core/user"
	"fgd/core/verify"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VerifyController struct {
	userUsecase   user.Usecase
	verifyUsecase verify.Usecase
}

func InitVerifyController(uc user.Usecase, vc verify.Usecase) *VerifyController {
	return &VerifyController{
		userUsecase:   uc,
		verifyUsecase: vc,
	}
}

func (cr *VerifyController) RequestEmailVerification(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)
	user, err := cr.userUsecase.GetProfileByID(claims.UserID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = cr.verifyUsecase.SendVerifyToken(user.Email, "EMAIL_VERIFY")
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *VerifyController) SubmitEmailVerification(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)
	user, err := cr.userUsecase.GetProfileByID(claims.UserID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	verify := request.Verify{}
	c.Bind(&verify)

	status, err := cr.verifyUsecase.CheckVerifyData(user.Email, verify.ToDomain())
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !status {
		return controllers.FailureResponse(c, http.StatusForbidden, "Code mismatch")
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *VerifyController) RequestForgetPassword(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)
	user, err := cr.userUsecase.GetProfileByID(claims.UserID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	err = cr.verifyUsecase.SendVerifyToken(user.Email, "EMAIL_VERIFY")
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

func (cr *VerifyController) SubmitForgetPasswordVerification(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)
	user, err := cr.userUsecase.GetProfileByID(claims.UserID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	verify := request.Verify{}
	c.Bind(&verify)

	status, err := cr.verifyUsecase.CheckVerifyData(user.Email, verify.ToDomain())
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	if !status {
		return controllers.FailureResponse(c, http.StatusForbidden, "Code mismatch")
	}

	// TODO Should redirect
	return controllers.SuccessResponse(c, http.StatusOK, nil)
}

// TODO How to protect this path
func (cr *VerifyController) SubmitNewPassword(c echo.Context) error {
	claims := middleware.ExtractUserClaims(c)

	req := request.PasswordReset{}
	err := c.Bind(&req)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusBadRequest, err.Error())
	}

	err = cr.userUsecase.UpdatePassword(req.NewPassword, claims.UserID)
	if err != nil {
		return controllers.FailureResponse(c, http.StatusInternalServerError, err.Error())
	}

	return controllers.SuccessResponse(c, http.StatusOK, nil)
}
