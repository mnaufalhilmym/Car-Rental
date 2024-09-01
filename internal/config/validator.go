package config

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var phoneNumberRegexp *regexp.Regexp

func NewValidator(phoneNumberPattern string) *validator.Validate {
	phoneNumberRegexp = regexp.MustCompile(phoneNumberPattern)

	v := validator.New(validator.WithRequiredStructEnabled())
	v.RegisterValidation("phone_number", validatePhoneNumber)

	return v
}

func validatePhoneNumber(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	return phoneNumberRegexp.MatchString(phoneNumber)
}
