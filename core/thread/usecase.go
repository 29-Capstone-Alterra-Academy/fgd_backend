package thread

import (
	"fgd/app/config"
	"fgd/core/topic"
	"fgd/core/user"
	"fgd/helper/format"
)

type threadUsecase struct {
	config           config.Config
	threadRepository Repository
	topicUsecase     topic.Usecase
	userUsecase      user.Usecase
}

func (uc *threadUsecase) CreateThread(data *Domain, userId int, topicId int) (Domain, error) {
	newThread, err := uc.threadRepository.CreateThread(data, userId, topicId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(newThread.Image1, uc.config)
	format.FormatImageLink(newThread.Image2, uc.config)
	format.FormatImageLink(newThread.Image3, uc.config)
	format.FormatImageLink(newThread.Image4, uc.config)
	format.FormatImageLink(newThread.Image5, uc.config)

	return newThread, nil
}

func (uc *threadUsecase) DeleteThread(userId int, threadId int) error {
	return uc.threadRepository.DeleteThread(userId, threadId)
}

func (uc *threadUsecase) GetThreadByAuthorID(userId, limit, offset int) ([]Domain, error) {
	return uc.threadRepository.GetThreadByAuthorID(userId, limit, offset)
}

func (uc *threadUsecase) GetThreadByTopicID(topicId, limit, offset int) ([]Domain, error) {
	return uc.threadRepository.GetThreadByTopicID(topicId, limit, offset)
}

func (uc *threadUsecase) Like(userId int, threadId int) error {
	return uc.threadRepository.Like(userId, threadId)
}

func (uc *threadUsecase) UndoLike(userId int, threadId int) error {
	return uc.threadRepository.UndoLike(userId, threadId)
}

func (uc *threadUsecase) UndoUnlike(userId int, threadId int) error {
	return uc.threadRepository.UndoUnlike(userId, threadId)
}

func (uc *threadUsecase) Unlike(userId int, threadId int) error {
	return uc.threadRepository.Unlike(userId, threadId)
}

func (uc *threadUsecase) UpdateThread(data *Domain, userId, threadId int) (Domain, error) {
	updatedThread, err := uc.threadRepository.UpdateThread(data, userId, threadId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(updatedThread.Image1, uc.config)
	format.FormatImageLink(updatedThread.Image2, uc.config)
	format.FormatImageLink(updatedThread.Image3, uc.config)
	format.FormatImageLink(updatedThread.Image4, uc.config)
	format.FormatImageLink(updatedThread.Image5, uc.config)

	return updatedThread, nil
}

func InitThreadUsecase(r Repository, tc topic.Usecase, uc user.Usecase, conf config.Config) Usecase {
	return &threadUsecase{
		config:           conf,
		threadRepository: r,
		topicUsecase:     tc,
		userUsecase:      uc,
	}
}
