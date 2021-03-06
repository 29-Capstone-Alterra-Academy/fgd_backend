package topic

import (
	"fgd/app/config"
	"fgd/core/user"
	"fgd/helper/format"
)

type topicUsecase struct {
	config          config.Config
	topicRepository Repository
	userUsecase     user.Usecase
}

func (uc *topicUsecase) GetTopicDetails(topicId int) (Domain, error) {
	topic, err := uc.topicRepository.GetTopicDetails(topicId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(uc.config, topic.ProfileImage)

	return topic, nil
}

func (uc *topicUsecase) CheckTopicAvailibility(topicName string) bool {
	return uc.topicRepository.CheckTopicAvailibility(topicName)
}

func (uc *topicUsecase) CreateTopic(data *Domain, userId int) (Domain, error) {
	newTopic, err := uc.topicRepository.CreateTopic(data)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(uc.config, newTopic.ProfileImage)

	err = uc.topicRepository.Subscribe(userId, newTopic.ID)

	return newTopic, err
}

func (uc *topicUsecase) GetSubscribedTopics(userId int) ([]Domain, error) {
	topics, err := uc.topicRepository.GetSubscribedTopics(userId)
	if err != nil {
		return []Domain{}, nil
	}

	for _, topic := range topics {
		format.FormatImageLink(uc.config, topic.ProfileImage)
	}

	return topics, nil
}

func (uc *topicUsecase) GetTopics(limit, offset int, sort_by string) ([]Domain, error) {
	topics, err := uc.topicRepository.GetTopics(limit, offset, sort_by)
	if err != nil {
		return []Domain{}, err
	}

	for _, topic := range topics {
		format.FormatImageLink(uc.config, topic.ProfileImage)
	}

	return topics, nil
}

func (uc *topicUsecase) GetTopicsByKeyword(keyword string, limit, offset int) ([]Domain, error) {
	topics, err := uc.topicRepository.GetTopicsByKeyword(keyword, limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, topic := range topics {
		format.FormatImageLink(uc.config, topic.ProfileImage)
	}

	return topics, nil
}

func (uc *topicUsecase) Subscribe(userId, topicId int) error {
	return uc.topicRepository.Subscribe(userId, topicId)
}

func (uc *topicUsecase) Unsubscribe(userId, topicId int) error {
	return uc.topicRepository.Unsubscribe(userId, topicId)
}

func (uc *topicUsecase) UpdateTopic(data *Domain, topicId int) (Domain, error) {
	updatedTopic, err := uc.topicRepository.UpdateTopic(data, topicId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(uc.config, updatedTopic.ProfileImage)

	return updatedTopic, nil
}

func InitTopicUsecase(r Repository, u user.Usecase, conf config.Config) Usecase {
	return &topicUsecase{
		config:          conf,
		topicRepository: r,
		userUsecase:     u,
	}
}
