package user

import (
	"fgd/core/user"
	"fgd/drivers/databases/notification"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Role         string
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Password     string
	ProfileImage *string
	Gender       string
	BirthDate    *time.Time
	Bio          string
	IsVerified   bool `gorm:"default:false"`

	Following []*User `gorm:"many2many:user_follow"`

	Notifications []notification.Notification
}

type UserModeratedTopic struct {
	ID []int
}

func (rec *User) toDomain() user.Domain {
	return user.Domain{
		ID:             int(rec.ID),
		Role:           rec.Role,
		Username:       rec.Username,
		Email:          rec.Email,
		Password:       rec.Password,
		ProfileImage:   rec.ProfileImage,
		Gender:         &rec.Gender,
		ModeratedTopic: &[]int{},
		BirthDate:      rec.BirthDate,
		IsVerified:     rec.IsVerified,
		CreatedAt:      rec.CreatedAt,
		UpdatedAt:      rec.UpdatedAt,
		DeletedAt:      rec.DeletedAt.Time,
	}
}

func (rec *UserModeratedTopic) toDomain() user.Domain {
	return user.Domain{
		ModeratedTopic: &rec.ID,
	}
}

func fromDomain(userDomain user.Domain) *User {
	return &User{
		Model: gorm.Model{
			ID:        uint(userDomain.ID),
			CreatedAt: userDomain.CreatedAt,
			UpdatedAt: userDomain.UpdatedAt,
			DeletedAt: gorm.DeletedAt{
				Time: userDomain.DeletedAt,
			},
		},
		Role:         userDomain.Role,
		Username:     userDomain.Username,
		Email:        userDomain.Email,
		Password:     userDomain.Password,
		ProfileImage: userDomain.ProfileImage,
		Gender:       *userDomain.Gender,
		BirthDate:    userDomain.BirthDate,
		IsVerified:   userDomain.IsVerified,
	}
}
