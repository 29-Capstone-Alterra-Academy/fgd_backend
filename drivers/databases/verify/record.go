package verify

import (
	"fgd/core/verify"
	"time"
)

type Verify struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Code      string
	Type      string // PASS_RESET, EMAIL_VERIFY
	ExpiresAt time.Time
}

func (r *Verify) toDomain() verify.Domain {
	return verify.Domain{
		Email:     r.Email,
		Code:      r.Code,
		ExpiresAt: r.ExpiresAt,
	}
}

func fromDomain(verifyDomain verify.Domain) Verify {
	return Verify{
		Email:     verifyDomain.Email,
		Code:      verifyDomain.Code,
		ExpiresAt: verifyDomain.ExpiresAt,
	}
}
