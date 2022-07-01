package response

import (
	"fgd/core/user"
	"time"
)

type UserProfile struct {
	ID           int        `json:"id"`
	Username     string     `json:"username"`
	Email        string     `json:"email"`
	Gender       *string    `json:"gender"`
	ProfileImage *string    `json:"profile_image"`
	IsVerified   bool       `json:"is_verified"`
	BirthDate    *time.Time `json:"birth_date"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type UserPublicProfile struct {
	ID             int
	Username       string
	ProfileImage   *string
	ThreadCount    int
	FollowingCount int
	FollowersCount int
}

func FromDomain(userDomain *user.Domain, scope string) interface{} {
	if scope == "personal" {
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
	} else {
		return UserPublicProfile{
			ID:             userDomain.ID,
			Username:       userDomain.Username,
			ProfileImage:   userDomain.ProfileImage,
			ThreadCount:    userDomain.ThreadCount,
			FollowingCount: userDomain.FollowingCount,
			FollowersCount: userDomain.FollowersCount,
		}
	}
}
