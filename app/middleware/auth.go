package middleware

import (
	"fgd/core/auth"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
	jwt.StandardClaims
	UserID     int
	UUID       string
	IsAdmin    bool
	Moderating []int
}

type JWTConfig struct {
	AuthUsecase   auth.Usecase
	Secret        string
	AccessExpiry  time.Duration
	RefreshExpiry time.Duration
}

type CustomToken struct {
	AccessExpiry  time.Duration `json:"-"`
	AccessToken   string        `json:"access_token"`
	AccessUUID    string        `json:"-"`
	RefreshExpiry time.Duration `json:"-"`
	RefreshToken  string        `json:"refresh_token"`
	RefreshUUID   string        `json:"-"`
}

func (c *JWTConfig) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(c.Secret),
	}
}

func (c *JWTConfig) GenerateToken(userId int, isAdmin bool, moderatedTopic []int) (CustomToken, error) {
	var token CustomToken
	var err error

	token.AccessExpiry = c.AccessExpiry
	token.AccessUUID = uuid.New().String()
	accessClaims := JWTCustomClaims{
		UserID:     userId,
		UUID:       token.AccessUUID,
		Moderating: moderatedTopic,
		IsAdmin:    isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(c.AccessExpiry).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
			Issuer:    "nomizo",
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	token.AccessToken, err = accessToken.SignedString([]byte(c.Secret))
	if err != nil {
		return token, err
	}

	token.RefreshExpiry = c.RefreshExpiry
	token.RefreshUUID = uuid.New().String()
	refreshClaims := JWTCustomClaims{
		UserID: userId,
		UUID:   token.RefreshUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(c.RefreshExpiry).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
			Issuer:    "nomizo",
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	token.RefreshToken, err = refreshToken.SignedString([]byte(c.Secret))
	if err != nil {
		return CustomToken{}, err
	}

	return token, nil
}

func (c *JWTConfig) CustomKeyFunc(token *jwt.Token) (interface{}, error) {
	if token.Method.Alg() != "HS256" {
		return CustomToken{}, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	claims := token.Claims.(*JWTCustomClaims)

	if eqIss := claims.VerifyIssuer("nomizo", false); !eqIss {
		return CustomToken{}, fmt.Errorf("error parsing token: invalid issuer")
	}

	cacheErr := c.AuthUsecase.CheckAuth(claims.UserID, claims.UUID)
	if cacheErr != nil {
		return CustomToken{}, fmt.Errorf("error parsing token: token already expired")
	}

	return []byte(c.Secret), nil
}

func ExtractUserClaims(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)
	userClaims := user.Claims.(*JWTCustomClaims)
	return userClaims
}
