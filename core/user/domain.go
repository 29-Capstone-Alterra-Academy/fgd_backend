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
	ThreadCount    int
	FollowingCount int
	FollowersCount int
	ProfileImage   string
	Gender         string
	BirthDate      time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type Usecase interface {
	CreateToken(userId int, isAdmin bool, moderatedTopic []int) (middleware.CustomToken, error)
	CreateUser(data *Domain) (Domain, error)
	CheckUserAvailibility(username string) (bool, error)
	GetModeratedTopic(userId int) ([]int, error)
	GetPersonalProfile(userId int) (Domain, error)
	GetProfileByID(userId int) (Domain, error)
	GetUsers(limit, offset int) ([]Domain, error)
	UpdatePersonalProfile(data *Domain, userId int) (Domain, error)
	UpdatePassword(newPassword string, userId int) error
	UpdateProfileImage(data *Domain) (Domain, error)
	IsAdmin(userId int) (bool, error)
	FollowUser(userId, targetId int) error
	UnfollowUser(userId, targetId int) error
}

type Repository interface {
	CreateUser(data *Domain) (Domain, error)
	CheckUserAvailibility(username string) (bool, error)
	GetPersonalProfile(userId int) (Domain, error)
	GetProfileByID(userId int) (Domain, error)
	GetUsers(limit, offset int) ([]Domain, error)
	GetModeratedTopic(userId int) ([]int, error)
	UpdatePersonalProfile(data *Domain, userId int) (Domain, error)
	UpdatePassword(hashedPassword string, userId int) error
	UpdateProfileImage(data *Domain, userId int) error
	CheckIsAdmin(userId int) (bool, error)
	FollowUser(userId, targetId int) error
	UnfollowUser(userId, targetId int) error
}
