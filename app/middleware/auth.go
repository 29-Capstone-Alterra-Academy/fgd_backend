package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	jwt.StandardClaims
	UserID     int   `json:"user_id"`
	IsAdmin    bool  `json:"is_admin"`
	Moderating []int `json:"moderating"`
}

type JWTConfig struct {
	Secret        string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
}

type CustomToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (c *JWTConfig) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(c.Secret),
	}
}

func (c *JWTConfig) GenerateToken(userId int, isAdmin bool, moderatedTopic []int) (CustomToken, error) {
	accessClaims := JWTCustomClaims{
		UserID:     userId,
		Moderating: moderatedTopic,
		IsAdmin:    isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(c.AccessExpiry).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
			Issuer:    "nomizo",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	t, err := accessToken.SignedString([]byte(c.Secret))
	if err != nil {
		return CustomToken{}, err
	}

	refreshClaims := jwt.StandardClaims{
		ExpiresAt: time.Now().Local().Add(c.RefreshExpiry).Unix(),
		IssuedAt:  time.Now().Local().Unix(),
		Issuer:    "nomizo",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	rt, err := refreshToken.SignedString([]byte(c.Secret))
	if err != nil {
		return CustomToken{}, err
	}

	return CustomToken{
		AccessToken:  t,
		RefreshToken: rt,
	}, nil
}

func ExtractUserClaims(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	userClaims := user.Claims.(*JWTCustomClaims)
	return userClaims
}