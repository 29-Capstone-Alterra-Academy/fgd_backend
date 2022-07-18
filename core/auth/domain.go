package auth

import "time"

type Domain struct {
	AuthUUID    string
	RefreshUUID string
}

type Usecase interface {
	CheckAuth(uuid string) error
	DeleteAuth(uuid string) error
	StoreAuth(userId int, accessUuid, refreshUuid string, accessExpiry, refreshExpiry time.Duration) error
}

type Repository interface {
	FetchAuth(uuid string) error
	DeleteAuth(uuid string) error
	StoreAuth(userId int, accessUuid, refreshUuid string, accessExpiry, refreshExpiry time.Duration) error
}
