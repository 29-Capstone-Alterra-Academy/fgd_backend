package user

import (
	"fgd/app/middleware"
	"time"
)

type Domain struct {
	ID             int
	Role           string
	Username       string
	Email          string
	Password       string
	Bio            *string
	ThreadCount    int
	FollowingCount int
	FollowersCount int
	ProfileImage   *string
	Gender         *string
	ModeratedTopic *[]int
	IsVerified     bool
	BirthDate      *time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type Usecase interface {
	CreateToken(username, email, password string) (middleware.CustomToken, error)
	CreateUser(data *Domain) (Domain, error)
	CheckUserAvailibility(username string) bool
	CheckEmailAvailibility(email string) bool
	GetPersonalProfile(userId int) (Domain, error)
	GetProfileByID(userId int) (Domain, error)
	GetUsers(limit, offset int) ([]Domain, error)
	GetUsersByKeyword(keyword string, limit, offset int) ([]Domain, error)
	GetFollowers(userId int) ([]Domain, error)
	GetFollowing(userId int) ([]Domain, error)
	GetModerators(topicId int) ([]Domain, error)
	UpdatePersonalProfile(data *Domain, userId int) (Domain, error)
	UpdatePassword(newPassword string, userId int) error
	FollowUser(userId, targetId int) error
	UnfollowUser(userId, targetId int) error
}

type Repository interface {
	CreateUser(data *Domain) (Domain, error)
	CheckUserAvailibility(username string) bool
	CheckEmailAvailibility(email string) bool
	GetPersonalProfile(userId int) (Domain, error)
	GetProfileByID(userId int) (Domain, error)
	GetUsers(limit, offset int) ([]Domain, error)
	GetUsersByKeyword(keyword string, limit, offset int) ([]Domain, error)
	GetUserByEmail(email string) (Domain, error)
	GetUserByUsername(username string) (Domain, error)
	GetFollowers(userId int) ([]Domain, error)
	GetFollowing(userId int) ([]Domain, error)
	GetModerators(topicId int) ([]Domain, error)
	GetModeratedTopic(userId int) (Domain, error)
	UpdatePersonalProfile(data *Domain, userId int) (Domain, error)
	UpdatePassword(hashedPassword string, userId int) error
	CheckIsAdmin(userId int) (bool, error)
	FollowUser(userId, targetId int) error
	UnfollowUser(userId, targetId int) error
}
