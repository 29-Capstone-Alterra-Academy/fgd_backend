package search

import (
	"fgd/core/search"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type persistenceSearchRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceSearchRepository) QuerySearchHistory(userId uint, keyword string, limit int) ([]search.Domain, error) {
	historyDatas := []SearchHistory{}
	domains := []search.Domain{}

	err := rp.Conn.Limit(limit).Where("user_id = ? AND  query LIKE ?", userId, keyword+"%").Find(&historyDatas).Error
	if err != nil {
		return domains, err
	}

	for _, data := range historyDatas {
		domains = append(domains, data.toDomain())
	}

	return domains, nil
}

func (rp *persistenceSearchRepository) GetLastSearchHistory(userId uint, limit int) ([]search.Domain, error) {
	historyDatas := []SearchHistory{}
	domains := []search.Domain{}

	err := rp.Conn.Limit(limit).Where("user_id = ? ", userId).Find(&historyDatas).Error
	if err != nil {
		return domains, err
	}

	for _, data := range historyDatas {
		domains = append(domains, data.toDomain())
	}

	return domains, nil
}

func (rp *persistenceSearchRepository) StoreSearchKeyword(userId uint, data *search.Domain) error {
	userAcc := user.User{}
	fetchErr := rp.Conn.Take(&userAcc, userId).Error
	if fetchErr != nil {
		return fetchErr
	}

	keywordData := fromDomain(*data)
	keywordData.UserID = userAcc.ID

	return rp.Conn.Create(&keywordData).Error
}

func InitPersistenceSearchRepository(c *gorm.DB) search.Repository {
	return &persistenceSearchRepository{
		Conn: c,
	}
}
