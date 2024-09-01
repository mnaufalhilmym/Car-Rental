package config

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var phoneNumberRegexp *regexp.Regexp

func RegisterCustomValidation(phoneNumberPattern string) {
	phoneNumberRegexp = regexp.MustCompile(phoneNumberPattern)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("phone_number", validatePhoneNumber)
	}
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	return phoneNumberRegexp.MatchString(phoneNumber)
}
