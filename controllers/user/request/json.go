package request

import (
	"fgd/core/user"
	"time"
)

type UserAuth struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserProfile struct {
	Username     string
	Email        string
	Gender       *string
	BirthDate    *time.Time
	ProfileImage *string
}

type TokenRequest struct {
	RefreshToken string `JSON:"refresh_token"`
}

func (r *UserAuth) ToDomain() *user.Domain {
	return &user.Domain{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

func (r *UserProfile) ToDomain() *user.Domain {
	return &user.Domain{
		Username:     r.Username,
		Email:        r.Email,
		Gender:       r.Gender,
		BirthDate:    r.BirthDate,
		ProfileImage: r.ProfileImage,
	}
}
