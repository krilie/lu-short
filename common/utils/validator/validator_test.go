package validator

import (
	"github.com/asaskevich/govalidator"
	"testing"
)

func TestInit(t *testing.T) {
	isEmail := govalidator.IsEmail("me@i.com")
	t.Log(isEmail)
}
