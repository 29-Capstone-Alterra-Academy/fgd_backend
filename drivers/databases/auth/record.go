package auth

import "fgd/core/auth"

type Auth struct {
	ID          uint   `gorm:"primaryKey; auto_increment"`
	UserID      uint   `gorm:"not null;"`
	AccessUUID  string `gorm:"size:255; not null;"`
	RefreshUUID string `gorm:"size:255; not null;"`
}

func (rec *Auth) toDomain() auth.Domain {
	return auth.Domain{
		AuthUUID:    rec.AccessUUID,
		RefreshUUID: rec.RefreshUUID,
	}
}

func fromDomain(authDomain auth.Domain) *Auth {
	return &Auth{
		AccessUUID:  authDomain.AuthUUID,
		RefreshUUID: authDomain.RefreshUUID,
	}
}
