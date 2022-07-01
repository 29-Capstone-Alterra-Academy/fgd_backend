package reply

import "fgd/core/user"

type replyUsecase struct {
	replyRepository Repository
	userUsecase     user.Usecase
}

func (uc *replyUsecase) CreateReplyReply(data *Domain, userId int, replyId int) (Domain, error) {
	panic("unimplemented")
}

func (uc *replyUsecase) CreateReplyThread(data *Domain, userId int, threadId int) (Domain, error) {
	panic("unimplemented")
}

func (uc *replyUsecase) DeleteReply(userId int, replyId int) error {
	panic("unimplemented")
}

func (uc *replyUsecase) EditReply(data *Domain, userId int, replyId int) (Domain, error) {
	panic("unimplemented")
}

func (uc *replyUsecase) Like(userId int, replyId int) error {
	panic("unimplemented")
}

func (uc *replyUsecase) UndoLike(userId int, replyId int) error {
	panic("unimplemented")
}

func (uc *replyUsecase) UndoUnlike(userId int, replyId int) error {
	panic("unimplemented")
}

func (uc *replyUsecase) Unlike(userId int, replyId int) error {
	panic("unimplemented")
}

func InitReplyUsecase(r Repository, uc user.Usecase) Usecase {
	return &replyUsecase{
		replyRepository: r,
		userUsecase:     uc,
	}
}
