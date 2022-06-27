package verify

import (
	"fgd/core/verify"

	"gorm.io/gorm"
)

type persistenceVerifyRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceVerifyRepository) DeleteVerifyData(email string) error {
	res := rp.Conn.Where("email = ?", email).Delete(&Verify{})
	return res.Error
}

func (rp *persistenceVerifyRepository) FetchVerifyData(email string) (verify.Domain, error) {
	verify := Verify{}
	res := rp.Conn.Where("email = ?", email).Take(&verify)

	return verify.toDomain(), res.Error
}

func (rp *persistenceVerifyRepository) StoreVerifyData(data verify.Domain) error {
	verify := fromDomain(data)
	res := rp.Conn.Create(&verify)
	return res.Error
}

func InitPersistenceVerifyRepository(c *gorm.DB) verify.Repository {
	return &persistenceVerifyRepository{
		Conn: c,
	}
}
