package verify

import "time"

type Domain struct {
	Email     string
	Code      string
	Type      string
	ExpiresAt time.Time
}

type Usecase interface {
	StoreVerifyData(email string, verify_type string, data Domain) error
	FetchVerifyData(email string) (Domain, error)
	DeleteVerifyData(email string) error
}
type Repository interface {
	StoreVerifyData(email string, data Domain) error
	FetchVerifyData(email string) (Domain, error)
	DeleteVerifyData(email string) error
}
