package reply

import "fgd/core/user"

type replyUsecase struct {
	replyRepository Repository
	userUsecase     user.Usecase
}

func (uc *replyUsecase) CreateReplyReply(data *Domain, userId int, replyId int) (Domain, error) {
	return uc.replyRepository.CreateReplyReply(data, userId, replyId)
}

func (uc *replyUsecase) CreateReplyThread(data *Domain, userId int, threadId int) (Domain, error) {
	return uc.replyRepository.CreateReplyThread(data, userId, threadId)
}

func (uc *replyUsecase) DeleteReply(userId int, replyId int) error {
	return uc.replyRepository.DeleteReply(userId, replyId)
}

func (uc *replyUsecase) EditReply(data *Domain, userId int, replyId int) (Domain, error) {
	return uc.replyRepository.EditReply(data, userId, replyId)
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

func InitReplyUsecase(r Repository, uc user.Usecase) Usecase {
	return &replyUsecase{
		replyRepository: r,
		userUsecase:     uc,
	}
}
