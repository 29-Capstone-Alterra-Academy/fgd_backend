package moderator

import (
	"fgd/core/moderator"
	"fgd/drivers/databases/topic"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type persistenceModeratorRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceModeratorRepository) ApplyPromotion(userId uint, topicId uint) (moderator.Domain, error) {
	user := user.User{Model: gorm.Model{ID: userId}}
	topic := topic.Topic{Model: gorm.Model{ID: topicId}}
	userErr := rp.Conn.Take(&user).Error
	if userErr != nil {
		return moderator.Domain{}, userErr
	}
	topicErr := rp.Conn.Take(&topic).Error
	if topicErr != nil {
		return moderator.Domain{}, topicErr
	}

	modReq := ModeratorRequest{
		User:  user,
		Topic: topic,
	}

	err := rp.Conn.Create(&modReq).Error

	return modReq.toDomain(), err
}

func (rp *persistenceModeratorRepository) ApprovePromotion(promotionId uint) error {
	modReq := ModeratorRequest{ID: promotionId}
	fetchErr := rp.Conn.Joins("User").Joins("Topic").Take(&modReq).Error
	if fetchErr != nil {
		return fetchErr
	}

	promoteErr := rp.Conn.Model(&modReq.Topic).Association("ModeratedBy").Append(&modReq.User)
	if promoteErr != nil {
		return promoteErr
	}

	return rp.Conn.Delete(&modReq).Error
}

func (rp *persistenceModeratorRepository) GetPromotionRequest() ([]moderator.Domain, error) {
	modReqs := []ModeratorRequest{}
	domains := []moderator.Domain{}
	err := rp.Conn.Joins("User").Joins("Topic").Find(&modReqs).Error
	if err != nil {
		return domains, err
	}

	for _, req := range modReqs {
		domains = append(domains, req.toDomain())
	}

	return domains, nil
}

func (rp *persistenceModeratorRepository) RejectPromotion(promotionId uint) error {
	modReq := ModeratorRequest{ID: promotionId}
	return rp.Conn.Delete(&modReq).Error
}

func (rp *persistenceModeratorRepository) RemoveModerator(userId, topicId uint) error {
	user := user.User{Model: gorm.Model{ID: userId}}
	topic := topic.Topic{Model: gorm.Model{ID: topicId}}

	return rp.Conn.Model(&topic).Association("ModeratedBy").Delete(&user)
}

func InitPersistenceModeratorRepository(c *gorm.DB) moderator.Repository {
	return &persistenceModeratorRepository{
		Conn: c,
	}
}
