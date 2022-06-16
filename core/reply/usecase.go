package reply

import "fgd/core/user"

type replyUsecase struct {
	replyRepository Repository
	userUsecase     user.Usecase
}

func InitReplyUsecase(r Repository, uc user.Usecase) Usecase {
	return &replyUsecase{
		replyRepository: r,
		userUsecase:     uc,
	}
}
