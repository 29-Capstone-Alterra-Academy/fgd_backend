package response

import (
	"fgd/core/user"
	"time"
)

type UserProfile struct {
	ID           int       `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	Gender       string    `json:"gender"`
	ProfileImage string    `json:"profile_image"`
	IsVerified   bool      `json:"is_verified"`
	BirthDate    time.Time `json:"birth_date"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(userDomain *user.Domain) UserProfile {
	return UserProfile{
		ID:           userDomain.ID,
		Username:     userDomain.Username,
		Email:        userDomain.Email,
		Gender:       userDomain.Gender,
		ProfileImage: userDomain.ProfileImage,
		IsVerified:   userDomain.IsVerified,
		BirthDate:    userDomain.BirthDate,
		UpdatedAt:    userDomain.UpdatedAt,
	}
}
