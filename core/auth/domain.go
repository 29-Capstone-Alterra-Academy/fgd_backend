package auth

import "fgd/app/middleware"

type Domain struct {
	AuthUUID    string
	RefreshUUID string
}

type Usecase interface {
	FetchAuth(userId int) (Domain, error)
	DeleteAuth(userId int) error
	StoreAuth(userId int, auth middleware.CustomToken) error
}

type Repository interface {
	FetchAuth(userId int) (Domain, error)
	DeleteAuth(userId int) error
	StoreAuth(userId int, auth middleware.CustomToken) error
}
