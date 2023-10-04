package validation

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate
var Translator ut.Translator

func LoadEnglishTranslator() {
	en := en.New()
	uni := ut.New(en, en)
	Translator, _ = uni.GetTranslator("en")
}

func RegisterTranslations(v *validator.Validate, translations ...[][]string) *validator.Validate {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	for _, collection := range translations {
		for _, trans := range collection {
			from := trans[0]
			to := trans[1]

			v.RegisterTranslation(from, Translator, func(ut ut.Translator) error {
				return ut.Add(from, to, true)
			}, func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T(from, fe.Field())

				return t
			})
		}
	}

	return v
}

func LoadValidator() {
	LoadEnglishTranslator()

	Validator = validator.New()

	translations := [][]string{
		{
			"required",
			"{0} is required",
		},
		{
			"alpha",
			"{0} must be composed of alpha characters only",
		},
		{
			"alphanum",
			"{0} must be composed of alphanumeric characters only",
		},
		{
			"email",
			"{0} must be a valid email address",
		},
		{
			"required_with",
			"{0} is required",
		},
		{
			"required_without_all",
			"{0} is optional, but required if you don't provide any other field",
		},
	}

	Validator = RegisterTranslations(Validator, translations)
}

func Validate[T any](target T) []string {
	err := Validator.Struct(target)
	errors := []string{}

	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, err := range errs {
			errors = append(errors, err.Translate(Translator))
		}
	}

	return errors
}
