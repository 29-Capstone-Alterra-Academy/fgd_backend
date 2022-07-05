package topic

import "time"

type Domain struct {
	ID               int
	Name             string
	ProfileImage     *string
	Description      string
	Rules            *string
	ActivityCount    int
	ContributorCount int
	ModeratorCount   int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Usecase interface {
	CreateTopic(data *Domain, userId int) (Domain, error)
	CheckTopicAvailibility(topicName string) bool
	GetTopics(limit, offset int, sort_by string) ([]Domain, error)
	GetTopicDetails(topicId int) (Domain, error)
	GetModerators(topicId int) ([]Domain, error) // TODO
	UpdateTopic(data *Domain, topicId int) (Domain, error)
	Subscribe(userId, topicId int) error
	Unsubscribe(userId, topicId int) error
}

type Repository interface {
	CreateTopic(data *Domain) (Domain, error)
	CheckTopicAvailibility(topicName string) bool
	GetTopics(limit, offset int, sort_by string) ([]Domain, error)
	GetTopicDetails(topicId int) (Domain, error)
	GetModerators(topicId int)
	UpdateTopic(data *Domain, topicId int) (Domain, error)
	Subscribe(userId, topicId int) error
	Unsubscribe(userId, topicId int) error
}
