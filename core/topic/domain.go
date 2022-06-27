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
	CheckTopicAvailibility(topicName string) error
	GetTopics(limit, offset int) ([]Domain, error)
	UpdateTopic(data *Domain) (Domain, error)
	Subscribe(topicId int) error
	Unsubscribe(topicId int) error
}

type Repository interface {
	CreateTopic(data *Domain)
	CheckTopicAvailibility(topicName string)
	GetTopics(limit, offset int) ([]Domain, error)
	GetModerators(topicId int)
	UpdateTopic(data *Domain, userId int)
	Subscribe(userId, topicId int) error
	Unsubscribe(userId, topicId int) error
}
