package topic

import (
	"errors"
	"fgd/core/topic"
	"fgd/drivers/databases/user"

	"gorm.io/gorm"
)

type persistenceTopicRepository struct {
	Conn *gorm.DB
}

func (rp *persistenceTopicRepository) CheckTopicAvailibility(topicName string) bool {
	topic := Topic{}

	err := rp.Conn.Where("name = ?", topicName).Take(&topic).Error

	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (rp *persistenceTopicRepository) CreateTopic(data *topic.Domain) (topic.Domain, error) {
	newTopic := fromDomain(*data)

	tx := rp.Conn.Begin()

	err := tx.Create(&newTopic).Error
	if err != nil {
		tx.Rollback()
		return topic.Domain{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		return topic.Domain{}, err
	}

	return newTopic.toDomain(), nil
}

func (rp *persistenceTopicRepository) GetModerators(topicId int) {
	panic("unimplemented")
}

func (rp *persistenceTopicRepository) GetTopicDetails(topicId int) (topic.Domain, error) {
	existingTopic := Topic{}
	err := rp.Conn.Take(&existingTopic, topicId).Error
	if err != nil {
		return topic.Domain{}, err
	}

	topicDomain := existingTopic.toDomain()

	var threadCount int64
	var replyCount int64
	var contributorCount int64
	var moderatorCount int64

	rp.Conn.Table("threads").Where("topic_id = ?", topicDomain.ID).Count(&threadCount)
	rp.Conn.Table("replies").Where("topic_id = ?", topicDomain.ID).Count(&replyCount)
	topicDomain.ActivityCount = int(threadCount + replyCount)
	rp.Conn.Table("threads").Where("topic_id = ?", topicDomain.ID).Distinct("author_id").Count(&contributorCount)
	topicDomain.ContributorCount = int(contributorCount)
	rp.Conn.Table("topic_moderator").Where("user_id = ?", topicDomain.ID).Count(&moderatorCount)
	topicDomain.ModeratorCount = int(moderatorCount)

	return topicDomain, nil
}

func (rp *persistenceTopicRepository) GetSubscribedTopics(userId int) ([]topic.Domain, error) {
	topics := []Topic{}
	topicDomains := []topic.Domain{}

	err := rp.Conn.Select("topics.id", "topics.name", "topics.profile_image").Joins("left join subscribed_topic on subscribed_topic.topic_id = topics.id").Where("subscribed_topic.user_id = ?", userId).Find(&topics).Error
	if err != nil {
		return topicDomains, err
	}

	for _, topic := range topics {
		topicDomains = append(topicDomains, topic.toDomain())
	}

	return topicDomains, nil
}

func (rp *persistenceTopicRepository) GetTopics(limit, offset int, sort_by string) ([]topic.Domain, error) {
	topics := []Topic{}

	// TODO Handle sort_by
	err := rp.Conn.Limit(limit).Offset(offset).Omit("ModeratedBy", "SubscribedBy", "Rules").Find(&topics).Error
	if err != nil {
		return []topic.Domain{}, err
	}

	topicDomains := []topic.Domain{}
	for _, topic := range topics {
		topicDomain := topic.toDomain()

		var threadCount int64
		var replyCount int64
		var contributorCount int64
		var moderatorCount int64

		rp.Conn.Table("threads").Where("topic_id = ?", topicDomain.ID).Count(&threadCount)
		rp.Conn.Table("replies").Where("topic_id = ?", topicDomain.ID).Count(&replyCount)
		topicDomain.ActivityCount = int(threadCount + replyCount)
		rp.Conn.Table("threads").Where("topic_id = ?", topicDomain.ID).Distinct("author_id").Count(&contributorCount)
		topicDomain.ContributorCount = int(contributorCount)
		rp.Conn.Table("topic_moderator").Where("user_id = ?", topicDomain.ID).Count(&moderatorCount)
		topicDomain.ModeratorCount = int(moderatorCount)

		topicDomains = append(topicDomains, topicDomain)
	}

	return topicDomains, nil
}

func (rp *persistenceTopicRepository) GetTopicsByKeyword(keyword string, limit, offset int) ([]topic.Domain, error) {
	topics := []Topic{}

	err := rp.Conn.Limit(limit).Offset(offset).Select("ID", "Name", "ProfileImage").Where("UPPER(name) LIKE UPPER(?)", "%"+keyword+"%").Find(&topics).Error
	if err != nil {
		return []topic.Domain{}, err
	}

	topicDomains := []topic.Domain{}
	for _, topic := range topics {
		topicDomain := topic.toDomain()
		var threadCount int64

		rp.Conn.Table("threads").Where("topic_id = ?", topicDomain.ID).Count(&threadCount)
		topicDomain.ActivityCount = int(threadCount)

		topicDomains = append(topicDomains, topicDomain)
	}

	return topicDomains, nil
}

func (rp *persistenceTopicRepository) Subscribe(userId int, topicId int) error {
	user := user.User{Model: gorm.Model{ID: uint(userId)}}
	topic := Topic{Model: gorm.Model{ID: uint(topicId)}}

	tx := rp.Conn.Begin()

	err := tx.
		Model(&topic).
		Association("SubscribedBy").
		Append(&user)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceTopicRepository) Unsubscribe(userId int, topicId int) error {
	user := user.User{Model: gorm.Model{ID: uint(userId)}}
	topic := Topic{Model: gorm.Model{ID: uint(topicId)}}

	tx := rp.Conn.Begin()

	err := tx.
		Model(&topic).
		Association("SubscribedBy").
		Delete(&user)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (rp *persistenceTopicRepository) UpdateTopic(data *topic.Domain, topicId int) (topic.Domain, error) {
	existingTopic := Topic{}

	tx := rp.Conn.Begin()

	err := tx.Take(&existingTopic, topicId).Error
	if err != nil {
		return topic.Domain{}, err
	}

	updatedTopic := fromDomain(*data)

	existingTopic.Description = updatedTopic.Description
	existingTopic.Rules = updatedTopic.Rules

	err = tx.Save(&existingTopic).Error
	if err != nil {
		tx.Rollback()
		return topic.Domain{}, err
	}

	err = tx.Commit().Error
	if err != nil {
		return topic.Domain{}, err
	}

	return existingTopic.toDomain(), nil
}

func InitPersistenceTopicRepository(c *gorm.DB) topic.Repository {
	return &persistenceTopicRepository{
		Conn: c,
	}
}
