package auth

import (
	"fmt"
	"time"
)

type authUsecase struct {
	authRepository Repository
}

func (uc *authUsecase) CheckAuth(userId int, uuid string) error {
	cacheId, err := uc.authRepository.FetchAuth(uuid)
	if err != nil || cacheId == 0 {
		return err
	}

	if userId != cacheId {
		return fmt.Errorf("error: token data mismatch")
	}

	return nil
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
