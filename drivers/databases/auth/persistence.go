package auth

import (
	"fgd/app/middleware"
	"fgd/core/auth"

	"gorm.io/gorm"
)

type persistenceAuthRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceAuthRepository) DeleteAuth(userId int) error {
	return rp.Conn.Where("user_id = ?", userId).Delete(&Auth{}).Error
}

func (rp *persistenceAuthRepository) FetchAuth(userId int) (auth.Domain, error) {
	auth := Auth{}
	res := rp.Conn.Where("user_id = ?", userId).Take(&auth)
	return auth.toDomain(), res.Error
}

func (rp *persistenceAuthRepository) StoreAuth(userId int, auth middleware.CustomToken) error {
	authStore := Auth{
		UserID:      uint(userId),
		AccessUUID:  auth.AccessUUID,
		RefreshUUID: auth.RefreshUUID,
	}

	res := rp.Conn.Create(&authStore)

	return res.Error
}

func InitPersistenceAuthRepository(c *gorm.DB) auth.Repository {
	return &persistenceAuthRepository{
		Conn: c,
	}
}
