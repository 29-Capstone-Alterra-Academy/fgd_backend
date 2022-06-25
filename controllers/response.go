package controllers

import "github.com/labstack/echo/v4"

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c echo.Context, status int, data interface{}) error {
	r := Response{}
	r.Data = data

	return c.JSON(status, r.Data)
}

func FailureResponse(c echo.Context, status int, errmsg string) error {
	r := Response{}
	r.Message = errmsg

	return c.JSON(status, r.Message)
}
