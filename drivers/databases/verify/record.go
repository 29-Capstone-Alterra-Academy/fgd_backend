package verify

import "time"

type EmailAuth struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	Code      string
	Type      string // PASS_RESET, EMAIL_VERIFY
	ExpiresAt time.Time
}
