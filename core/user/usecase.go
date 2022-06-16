package user

type userUsecase struct {
	userRepository Repository
}

func InitUserUsecase(r Repository) Usecase {
	return &userUsecase{
		userRepository: r,
	}
}
