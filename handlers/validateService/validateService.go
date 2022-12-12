package validateService

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	validate = binding.Validator.Engine().(*validator.Validate)
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func TranslateError(err error) map[string]string {
	validatorErrs := err.(validator.ValidationErrors)
	m := validatorErrs.Translate(trans)
	return m
}
