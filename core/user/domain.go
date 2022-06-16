package user

import (
	"time"
)

type Domain struct {
	ID             int
	Username       string
	Email          string
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
	GetPersonalProfile() (Domain, error)
	GetProfileByID(userId int) (Domain, error)
	UpdatePersonalProfile(data *Domain) (Domain, error)
	UpdatePassword(newPassword string) error
	UpdateProfileImage(data *Domain) (Domain, error)
	CheckUserAvailibility(username string) error
	FollowUser(userId int) error
	UnfollowUser(userId int) error
}

type Repository interface {
	GetUsers(limit, offset int) ([]Domain, error)
	CreateUser(data *Domain) (Domain, error)
	GetPersonalProfile(userId int) (Domain, error)
	GetProfileByID(userId int) (Domain, error)
	UpdatePersonalProfile(data *Domain, userId int) (Domain, error)
	UpdatePassword(hashedPassword string) error
	UpdateProfileImage(data *Domain, userId int)
	CheckAvailibility(username string) (bool, error)
	FollowUser(userId, targetId int) error
	UnfollowUser(userId, targetId int) error
}
