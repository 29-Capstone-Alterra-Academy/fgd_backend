package user

import (
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
	GetUsers(limit, offset int) ([]Domain, error)
	CreateUser(data *Domain) (Domain, error)
	GetPersonalProfile(userId int) (Domain, error)
	GetProfileByID(userId int) (Domain, error)
	UpdatePersonalProfile(data *Domain, userId int) (Domain, error)
	UpdatePassword(newPassword string, userId int) error
	UpdateProfileImage(data *Domain) (Domain, error)
	CheckRoleByID(userId int) (string, error)
	CheckUserAvailibility(username string) (bool, error)
	FollowUser(userId, targetId int) error
	UnfollowUser(userId, targetId int) error
}

type Repository interface {
	GetUsers(limit, offset int) ([]Domain, error)
	CreateUser(data *Domain) (Domain, error)
	GetPersonalProfile(userId int) (Domain, error)
	GetProfileByID(userId int) (Domain, error)
	UpdatePersonalProfile(data *Domain, userId int) (Domain, error)
	UpdatePassword(hashedPassword string, userId int) error
	UpdateProfileImage(data *Domain, userId int) error
	CheckRoleByID(userId int) (string, error)
	CheckAvailibility(username string) (bool, error)
	FollowUser(userId, targetId int) error
	UnfollowUser(userId, targetId int) error
}
