package validator

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
)

func IdStrValid(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if str, ok := field.Interface().(string); ok {
		if !IsIdStr(str) {
			return false
		}
	}
	return true
}
func PasswordValid(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if str, ok := field.Interface().(string); ok {
		if !IsPassword(str) {
			return false
		}
	}
	return true
}
func LoginNameValid(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if str, ok := field.Interface().(string); ok {
		if !IsLoginName(str) {
			return false
		}
	}
	return true
}
func PhoneNumValid(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if str, ok := field.Interface().(string); ok {
		if !IsPhoneNum(str) {
			return false
		}
	}
	return true
}
