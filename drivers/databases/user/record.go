package user

import (
	"fgd/core/user"
	// "fgd/drivers/databases/notification"
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
	Gender       *string
	BirthDate    *time.Time
	Bio          *string
	IsVerified   bool `gorm:"default:false"`

	Following []*User `gorm:"many2many:user_follow"`

	UserReports []*User `gorm:"many2many:user_reports;foreignKey:ID;joinForeignKey:SuspectID;references:ID;constraint:OnDelete:CASCADE"`

	// Notifications []notification.Notification
}

type UserModeratedTopic struct {
	TopicID []int
}

func (rec *User) toDomain() user.Domain {
	user := user.Domain{
		ID:           int(rec.ID),
		Role:         rec.Role,
		Username:     rec.Username,
		Email:        rec.Email,
		Password:     rec.Password,
		Bio:          rec.Bio,
		ProfileImage: rec.ProfileImage,
		Gender:       rec.Gender,
		IsVerified:   rec.IsVerified,
		BirthDate:    rec.BirthDate,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
	}

	if rec.DeletedAt.Valid {
		user.DeletedAt = &rec.DeletedAt.Time
	}

	return user
}

func (rec *UserModeratedTopic) toDomain() user.Domain {
	return user.Domain{
		ModeratedTopic: &rec.TopicID,
	}
}

func fromDomain(userDomain user.Domain) *User {
	return &User{
		Model:        gorm.Model{ID: uint(userDomain.ID)},
		Role:         userDomain.Role,
		Username:     userDomain.Username,
		Email:        userDomain.Email,
		Password:     userDomain.Password,
		ProfileImage: userDomain.ProfileImage,
		Gender:       userDomain.Gender,
		BirthDate:    userDomain.BirthDate,
		Bio:          userDomain.Bio,
		IsVerified:   userDomain.IsVerified,
	}
}
