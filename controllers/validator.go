package controllers

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type CustomValidator struct {
	Validator  *validator.Validate
	Translator ut.Translator
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return cv.translate(err)
	}
	return nil
}

func (cv *CustomValidator) translate(err error) error {
	if err != nil {
		var errorResult string

		validatorErrors := err.(validator.ValidationErrors)
		for _, e := range validatorErrors {
			t := e.Translate(cv.Translator)
			errorResult += t + ","
		}

		return fmt.Errorf("%s", errorResult)
	}
	return nil
}

func InitCustomValidator() CustomValidator {
	validator := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	enTranslations.RegisterDefaultTranslations(validator, trans)

	return CustomValidator{
		Validator:  validator,
		Translator: trans,
	}
}
