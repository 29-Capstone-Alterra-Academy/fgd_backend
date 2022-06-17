package topic

type Domain struct {
	ID           int
	Name         string
	ProfileImage string
	Description  string
	Rules        string
}

type Usecase interface {
	CreateTopic(data *Domain) (Domain, error)
	CheckTopicAvailibility(topicName string) (bool, error)
	GetTopics(limit, offset int, sort_by string) ([]Domain, error)
	GetModerators(topicId int) ([]Domain, error) // TODO
	UpdateTopic(data *Domain, topicId int) (Domain, error)
	Subscribe(userId, topicId int) error
	Unsubscribe(userId, topicId int) error
}

type Repository interface {
	CreateTopic(data *Domain) (Domain, error)
	CheckTopicAvailibility(topicName string) (bool, error)
	GetTopics(limit, offset int, sort_by string) ([]Domain, error)
	GetModerators(topicId int)
	UpdateTopic(data *Domain, topicId int) (Domain, error)
	Subscribe(userId, topicId int) error
	Unsubscribe(userId, topicId int) error
}
