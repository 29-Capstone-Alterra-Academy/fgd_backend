package middleware

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

func AdminValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := ExtractUserClaims(c)
		isAdmin := claims.IsAdmin

		if !isAdmin {
			return echo.ErrUnauthorized
		}

		return next(c)
	}
}

func ModeratorValidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		topicId, _ := strconv.Atoi(c.Param("topicId"))
		claims := ExtractUserClaims(c)

		for _, topic := range claims.Moderating {
			if topicId == topic {
				return next(c)
			}
		}

		return echo.ErrUnauthorized
	}
}
