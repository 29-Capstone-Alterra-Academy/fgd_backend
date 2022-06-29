package topic

import "fgd/core/user"

type topicUsecase struct {
	topicRepository Repository
	userUsecase     user.Usecase
}

func (uc *topicUsecase) GetTopicDetails(topicId int) (Domain, error) {
	return uc.topicRepository.GetTopicDetails(topicId)
}

func (uc *topicUsecase) CheckTopicAvailibility(topicName string) (bool, error) {
	return uc.topicRepository.CheckTopicAvailibility(topicName)
}

func (uc *topicUsecase) CreateTopic(data *Domain) (Domain, error) {
	return uc.topicRepository.CreateTopic(data)
}

func (uc *topicUsecase) GetModerators(topicId int) ([]Domain, error) {
	panic("unimplemented")
}

func (uc *topicUsecase) GetTopics(limit, offset int, sort_by string) ([]Domain, error) {
	return uc.topicRepository.GetTopics(limit, offset, sort_by)
}

func (uc *topicUsecase) Subscribe(userId, topicId int) error {
	return uc.topicRepository.Subscribe(userId, topicId)
}

func (uc *topicUsecase) Unsubscribe(userId, topicId int) error {
	return uc.topicRepository.Unsubscribe(userId, topicId)
}

func (uc *topicUsecase) UpdateTopic(data *Domain, topicId int) (Domain, error) {
	return uc.topicRepository.UpdateTopic(data, topicId)
}

func InitTopicUsecase(r Repository, u user.Usecase) Usecase {
	return &topicUsecase{
		topicRepository: r,
		userUsecase:     u,
	}
}
