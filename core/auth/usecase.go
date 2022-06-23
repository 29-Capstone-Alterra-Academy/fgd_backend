package auth

type authUsecase struct {
	authRepository Repository
}

func InitAuthUsecase(r Repository) Usecase {
	return &authUsecase{
		authRepository: r,
	}
}
