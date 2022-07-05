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
	res := rp.Conn.Create(&newTopic)

	return newTopic.toDomain(), res.Error
}

func (rp *persistenceTopicRepository) GetModerators(topicId int) {
	panic("unimplemented")
}

func (rp *persistenceTopicRepository) GetTopicDetails(topicId int) (topic.Domain, error) {
	existingTopic := Topic{}
	res := rp.Conn.Take(&existingTopic, topicId)
	if res.Error != nil {
		return topic.Domain{}, res.Error
	}

	topicDomain := existingTopic.toDomain()
	var threadCount int64
	// var replyCount int64
	var contributorCount int64
	var moderatorCount int64

	rp.Conn.Table("threads").Where("topic_id = ?", topicDomain.ID).Count(&threadCount)
	// TODO Check reply in topic
	rp.Conn.Table("threads").Where("topic_id = ?", topicDomain.ID).Distinct("author_id").Count(&contributorCount)
	topicDomain.ContributorCount = int(contributorCount)
	rp.Conn.Table("topic_moderator").Where("user_id = ?", topicDomain.ID).Count(&moderatorCount)
	topicDomain.ModeratorCount = int(moderatorCount)

	return topicDomain, nil
}

func (rp *persistenceTopicRepository) GetTopics(limit, offset int, sort_by string) ([]topic.Domain, error) {
	topics := []Topic{}

	// TODO Handle sort_by
	res := rp.Conn.Limit(limit).Offset(offset).Omit("ModeratedBy", "SubscribedBy", "Rules").Find(&topics)

	if res.Error != nil {
		return []topic.Domain{}, res.Error
	}

	topicDomains := []topic.Domain{}
	for _, topic := range topics {
		topicDomain := topic.toDomain()
		topicDomains = append(topicDomains, topicDomain)
	}

	return topicDomains, nil
}

func (rp *persistenceTopicRepository) Subscribe(userId int, topicId int) error {
	topic := Topic{Model: gorm.Model{ID: uint(topicId)}}
	err := rp.Conn.
		Model(&topic).
		Association("SubscribedBy").
		Append(&user.User{
			Model: gorm.Model{ID: uint(userId)},
		})

	return err
}

func (rp *persistenceTopicRepository) Unsubscribe(userId int, topicId int) error {
	topic := Topic{Model: gorm.Model{ID: uint(topicId)}}
	err := rp.Conn.
		Model(&topic).
		Association("SubscribedBy").
		Delete(&user.User{
			Model: gorm.Model{ID: uint(userId)},
		})

	return err
}

func (rp *persistenceTopicRepository) UpdateTopic(data *topic.Domain, topicId int) (topic.Domain, error) {
	existingTopic := Topic{}
	fetchResult := rp.Conn.Take(&existingTopic, topicId)
	if fetchResult.Error != nil {
		return topic.Domain{}, fetchResult.Error
	}

	updatedTopic := fromDomain(*data)

	existingTopic.Description = updatedTopic.Description
	existingTopic.Rules = updatedTopic.Rules

	res := rp.Conn.Save(&existingTopic)
	return existingTopic.toDomain(), res.Error
}

func InitPersistenceTopicRepository(c *gorm.DB) topic.Repository {
	return &persistenceTopicRepository{
		Conn: c,
	}
}
