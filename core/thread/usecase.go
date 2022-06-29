package thread

import (
	"fgd/core/topic"
	"fgd/core/user"
)

type threadUsecase struct {
	threadRepository Repository
	topicUsecase     topic.Usecase
	userUsecase      user.Usecase
}

func (uc *threadUsecase) CreateThread(data *Domain, userId int, topicId int) (Domain, error) {
	return uc.threadRepository.CreateThread(data, userId, topicId)
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
	return uc.threadRepository.UpdateThread(data, userId, threadId)
}

func InitThreadUsecase(r Repository, tc topic.Usecase, uc user.Usecase) Usecase {
	return &threadUsecase{
		threadRepository: r,
		topicUsecase:     tc,
		userUsecase:      uc,
	}
}
