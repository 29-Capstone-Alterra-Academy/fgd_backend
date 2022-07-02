package reply

import (
	"fgd/app/config"
	"fgd/core/user"
	"fgd/helper/format"
)

type replyUsecase struct {
	config          config.Config
	replyRepository Repository
	userUsecase     user.Usecase
}

func (uc *replyUsecase) CreateReplyReply(data *Domain, userId int, replyId int) (Domain, error) {
	newReply, err := uc.replyRepository.CreateReplyReply(data, userId, replyId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(newReply.Image, uc.config)
	return newReply, nil
}

func (uc *replyUsecase) CreateReplyThread(data *Domain, userId int, threadId int) (Domain, error) {
	updatedReply, err := uc.replyRepository.CreateReplyThread(data, userId, threadId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(updatedReply.Image, uc.config)
	return updatedReply, nil
}

func (uc *replyUsecase) DeleteReply(userId int, replyId int) error {
	return uc.replyRepository.DeleteReply(userId, replyId)
}

func (uc *replyUsecase) EditReply(data *Domain, userId int, replyId int) (Domain, error) {
	updatedReply, err := uc.replyRepository.EditReply(data, userId, replyId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(updatedReply.Image, uc.config)
	return updatedReply, nil
}

func (uc *replyUsecase) Like(userId int, replyId int) error {
	return uc.replyRepository.Like(userId, replyId)
}

func (uc *replyUsecase) UndoLike(userId int, replyId int) error {
	return uc.replyRepository.UndoLike(userId, replyId)
}

func (uc *replyUsecase) UndoUnlike(userId int, replyId int) error {
	return uc.replyRepository.UndoUnlike(userId, replyId)
}

func (uc *replyUsecase) Unlike(userId int, replyId int) error {
	return uc.replyRepository.Unlike(userId, replyId)
}

func InitReplyUsecase(r Repository, uc user.Usecase, conf config.Config) Usecase {
	return &replyUsecase{
		config:          conf,
		replyRepository: r,
		userUsecase:     uc,
	}
}
