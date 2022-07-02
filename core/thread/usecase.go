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
	format.FormatImageLink(newThread.Topic.ProfileImage, uc.config)

	format.FormatImageLink(
		uc.config,
		newThread.Image1,
		newThread.Image2,
		newThread.Image3,
		newThread.Image4,
		newThread.Image5,
	)

	return newThread, nil
}

func (uc *threadUsecase) DeleteThread(userId int, threadId int) error {
	return uc.threadRepository.DeleteThread(userId, threadId)
}

func (uc *threadUsecase) GetThreadByID(threadId int) (Domain, error) {
	thread, err := uc.threadRepository.GetThreadByID(threadId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(thread.Topic.ProfileImage, uc.config)

	format.FormatImageLink(thread.Image1, uc.config)
	format.FormatImageLink(thread.Image2, uc.config)
	format.FormatImageLink(thread.Image3, uc.config)
	format.FormatImageLink(thread.Image4, uc.config)
	format.FormatImageLink(thread.Image5, uc.config)

	return thread, nil
}

func (uc *threadUsecase) GetThreadByAuthorID(userId, limit, offset int) ([]Domain, error) {
	threads, err := uc.threadRepository.GetThreadByAuthorID(userId, limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, thread := range threads {
		format.FormatImageLink(thread.Topic.ProfileImage, uc.config)

		format.FormatImageLink(thread.Image1, uc.config)
		format.FormatImageLink(thread.Image2, uc.config)
		format.FormatImageLink(thread.Image3, uc.config)
		format.FormatImageLink(thread.Image4, uc.config)
		format.FormatImageLink(thread.Image5, uc.config)
	}

	return threads, nil
}

func (uc *threadUsecase) GetThreadByTopicID(topicId, limit, offset int) ([]Domain, error) {
	threads, err := uc.threadRepository.GetThreadByTopicID(topicId, limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, thread := range threads {
		format.FormatImageLink(thread.Topic.ProfileImage, uc.config)

		format.FormatImageLink(thread.Image1, uc.config)
		format.FormatImageLink(thread.Image2, uc.config)
		format.FormatImageLink(thread.Image3, uc.config)
		format.FormatImageLink(thread.Image4, uc.config)
		format.FormatImageLink(thread.Image5, uc.config)
	}

	return threads, nil
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
	format.FormatImageLink(updatedThread.Topic.ProfileImage, uc.config)

	format.FormatImageLink(
		uc.config,
		updatedThread.Image1,
		updatedThread.Image2,
		updatedThread.Image3,
		updatedThread.Image4,
		updatedThread.Image5,
	)

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
