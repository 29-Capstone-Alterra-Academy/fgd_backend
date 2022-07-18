package request

import (
	"fgd/core/user"
	"time"
)

type UserLogin struct {
	Username string `json:"username" validate:"omitempty"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"required"`
}

type UserRegister struct {
	Username string `json:"username" validate:"omitempty,alphanum,min=6"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,alphanum,min=8"`
}

type UserProfile struct {
	Username     string
	Email        string
	Gender       *string
	Bio          *string
	BirthDate    *time.Time
	ProfileImage *string
}

type TokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required,jwt"`
}

func (r *UserLogin) ToDomain() *user.Domain {
	return &user.Domain{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}
}

func (r *UserRegister) ToDomain() *user.Domain {
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
		Bio:          r.Bio,
		BirthDate:    r.BirthDate,
		ProfileImage: r.ProfileImage,
	}
}
