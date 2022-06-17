package topic

import "fgd/core/user"

type topicUsecase struct {
	topicRepository Repository
	userUsecase     user.Usecase
}

func InitTopicUsecase(r Repository, u user.Usecase) Usecase {
	return &topicUsecase{
		topicRepository: r,
		userUsecase:     u,
	}
}
