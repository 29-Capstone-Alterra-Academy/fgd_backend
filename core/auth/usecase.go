package auth

import "time"

type authUsecase struct {
	authRepository Repository
}

func (uc *authUsecase) CheckAuth(uuid string) error {
	return uc.authRepository.FetchAuth(uuid)
}

func (uc *authUsecase) DeleteAuth(uuid string) error {
	return uc.authRepository.DeleteAuth(uuid)
}

func (uc *authUsecase) StoreAuth(userId int, accessUuid string, refreshUuid string, accessExpiry time.Duration, refreshExpiry time.Duration) error {
	return uc.authRepository.StoreAuth(userId, accessUuid, refreshUuid, accessExpiry, refreshExpiry)
}

func InitAuthUsecase(r Repository) Usecase {
	return &authUsecase{
		authRepository: r,
	}
}
