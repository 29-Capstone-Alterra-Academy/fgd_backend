package auth

import "fgd/app/middleware"

type authUsecase struct {
	authRepository Repository
}

func (uc *authUsecase) DeleteAuth(userId int) error {
	return uc.authRepository.DeleteAuth(userId)
}

func (uc *authUsecase) FetchAuth(userId int) (Domain, error) {
	return uc.authRepository.FetchAuth(userId)
}

func (uc *authUsecase) StoreAuth(userId int, auth middleware.CustomToken) error {
	return uc.authRepository.StoreAuth(userId, auth)
}

func InitAuthUsecase(r Repository) Usecase {
	return &authUsecase{
		authRepository: r,
	}
}
