package verify

type verifyUsecase struct {
	verifyRepository Repository
}

func (uc *verifyUsecase) DeleteVerifyData(email string) error {
	return uc.verifyRepository.DeleteVerifyData(email)
}

func (uc *verifyUsecase) FetchVerifyData(email string) (Domain, error) {
	return uc.verifyRepository.FetchVerifyData(email)
}

func (uc *verifyUsecase) StoreVerifyData(email string, verify_type string, data Domain) error {
	data.Type = verify_type
	return uc.verifyRepository.StoreVerifyData(email, data)
}

func InitVerifyUsecase(r Repository) Usecase {
	return &verifyUsecase{
		verifyRepository: r,
	}
}
