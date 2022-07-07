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
	BirthDate    *string    `json:"birth_date"`
	Bio          *string    `json:"bio"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}

type UserPublicProfile struct {
	ID             int
	Username       string
	ProfileImage   *string
	Bio            *string
	ThreadCount    int
	FollowingCount int
	FollowersCount int
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at"`
}

func FromDomain(userDomain *user.Domain, scope string) interface{} {
	if scope == "personal" {
		userProfile := UserProfile{
			ID:           userDomain.ID,
			Username:     userDomain.Username,
			Email:        userDomain.Email,
			Gender:       userDomain.Gender,
			ProfileImage: userDomain.ProfileImage,
			Bio:          userDomain.Bio,
			IsVerified:   userDomain.IsVerified,
			CreatedAt:    userDomain.CreatedAt,
			UpdatedAt:    userDomain.UpdatedAt,
			DeletedAt:    userDomain.DeletedAt,
		}
		if userDomain.BirthDate != nil {
			birthDate := userDomain.BirthDate.Format("2006-01-02")
			userProfile.BirthDate = &birthDate
		}
		return userProfile
	} else {
		return UserPublicProfile{
			ID:             userDomain.ID,
			Username:       userDomain.Username,
			ProfileImage:   userDomain.ProfileImage,
			Bio:            userDomain.Bio,
			ThreadCount:    userDomain.ThreadCount,
			FollowingCount: userDomain.FollowingCount,
			FollowersCount: userDomain.FollowersCount,
			CreatedAt:      userDomain.CreatedAt,
			UpdatedAt:      userDomain.UpdatedAt,
			DeletedAt:      userDomain.DeletedAt,
		}
	}
}
