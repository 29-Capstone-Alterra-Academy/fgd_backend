package verify

import (
	"fgd/helper/mail"
	"fgd/helper/token"
	"fmt"
	"time"
)

type verifyUsecase struct {
	verifyRepository Repository
	mailHelper       mail.MailHelper
}

// TODO Should we preserve this ?
func (uc *verifyUsecase) DeleteVerifyData(email string) error {
	return uc.verifyRepository.DeleteVerifyData(email)
}

func (uc *verifyUsecase) CheckVerifyData(email string, data Domain) (bool, error) {
	dbData, err := uc.verifyRepository.FetchVerifyData(email)
	if err != nil {
		return false, err
	}

	if !dbData.ExpiresAt.Before(time.Now()) {
		err = uc.verifyRepository.DeleteVerifyData(email)
		if err != nil {
			return false, err
		}
		return false, nil
	}

  // TODO Mark user as Verified

	if data.Code != dbData.Code {
		return false, nil
	}



	err = uc.verifyRepository.DeleteVerifyData(email)
	if err != nil {
		return true, err
	}

	return true, nil
}

func (uc *verifyUsecase) SendVerifyToken(email string, verify_type string) error {
	code := token.GenerateRandomNumber()

	data := Domain{}
	data.Code = code
	data.Email = email
	data.Type = verify_type

	err := uc.verifyRepository.StoreVerifyData(data)
	if err != nil {
		return err
	}

	err = uc.mailHelper.SendVerificationCode(code, email, verify_type)
	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}

	return nil
}

func InitVerifyUsecase(r Repository, m mail.MailHelper) Usecase {
	return &verifyUsecase{
		mailHelper:       m,
		verifyRepository: r,
	}
}
