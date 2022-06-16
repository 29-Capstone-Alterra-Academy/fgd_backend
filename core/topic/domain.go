package topic

type Domain struct {
	ID           int
	Name         string
	ProfileImage string
	Description  string
	Rules        string
}

type Usecase interface {
	GetTopics(limit, offset int) ([]Domain, error)
	CreateTopic(data *Domain) (Domain, error)
	UpdateTopic(data *Domain) (Domain, error)
	CheckTopicAvailibility(topicName string) error
	Subscribe(topicId int) error
	Unsubscribe(topicId int) error
}

type Repository interface {
	GetTopics(limit, offset int) ([]Domain, error)
	GetModerators(topicId int)
	CreateTopic(data *Domain)
	CheckTopicAvailibility(topicName string)
	UpdateTopic(data *Domain, userId int)
	Subscribe(userId, topicId int) error
	Unsubscribe(userId, topicId int) error
}
