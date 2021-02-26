package validator

import (
	"gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate = validator.New()

func IsPhoneNum2(phone string) bool {
	err := validate.Var(phone, "required,phone")
	if err != nil {
		return false
	} else {
		return true
	}
}
