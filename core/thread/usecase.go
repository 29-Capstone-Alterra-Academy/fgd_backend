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

func InitThreadUsecase(r Repository, tc topic.Usecase, uc user.Usecase) Usecase {
	return &threadUsecase{
		threadRepository: r,
		topicUsecase:     tc,
		userUsecase:      uc,
	}
}
