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

func (uc *replyUsecase) GetReplyByThreadID(threadId, limit, offset int) ([]Domain, error) {
	replies, err := uc.replyRepository.GetReplyByThreadID(threadId, limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, reply := range replies {
		format.FormatImageLink(uc.config, reply.Author.ProfileImage, reply.Image)
	}

	return replies, nil
}

func (uc *replyUsecase) GetReplyByID(replyId int, limit int) (Domain, error) {
	reply, err := uc.replyRepository.GetReplyByID(replyId, limit)
	if err != nil {
		return Domain{}, nil
	}

	format.FormatImageLink(uc.config, reply.Author.ProfileImage, reply.Image)

	if reply.Child != nil {
		for _, childReply := range *reply.Child {
			format.FormatImageLink(uc.config, childReply.Author.ProfileImage, childReply.Image)
		}
	}

	return reply, nil
}

func (uc *replyUsecase) GetReplyChilds(replyId, limit, offset int) ([]Domain, error) {
	replyChilds, err := uc.replyRepository.GetReplyChilds(replyId, limit, offset)
	if err != nil {
		return []Domain{}, err
	}

	for _, child := range replyChilds {
		format.FormatImageLink(uc.config, child.Author.ProfileImage, child.Image)
	}

	return replyChilds, nil
}

func (uc *replyUsecase) CreateReplyReply(data *Domain, userId int, replyId int) (Domain, error) {
	newReply, err := uc.replyRepository.CreateReplyReply(data, userId, replyId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(
		uc.config,
		newReply.Author.ProfileImage,
		newReply.Image,
	)

	return newReply, nil
}

func (uc *replyUsecase) CreateReplyThread(data *Domain, userId int, threadId int) (Domain, error) {
	updatedReply, err := uc.replyRepository.CreateReplyThread(data, userId, threadId)
	if err != nil {
		return Domain{}, err
	}

	format.FormatImageLink(
		uc.config,
		updatedReply.Author.ProfileImage,
		updatedReply.Image,
	)

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

	format.FormatImageLink(
		uc.config,
		updatedReply.Author.ProfileImage,
		updatedReply.Image,
	)

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
