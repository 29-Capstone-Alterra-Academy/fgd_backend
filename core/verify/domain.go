package verify

import "time"

type Domain struct {
	Email     string
	Code      string
	Type      string
	ExpiresAt time.Time
}

type Usecase interface {
	SendVerifyToken(email string, verify_type string) error
	CheckVerifyData(email string, data Domain) (bool, error)
	DeleteVerifyData(email string) error
}
type Repository interface {
	StoreVerifyData(data Domain) error
	FetchVerifyData(email string) (Domain, error)
	DeleteVerifyData(email string) error
}
