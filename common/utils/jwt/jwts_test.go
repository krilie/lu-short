package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"lu-short/common/utils/jsonutil"
	"testing"
	"time"
)

func TestNewJwtToken(t *testing.T) {
	fmt.Println("jwt jwt jwt jwt")
	var userClaims UserClaims
	userClaims.ClientId = "123"
	userClaims.Iss = "sys"
	userClaims.UserId = "34"
	userClaims.Jti = "45"
	userClaims.Iat = time.Now().Unix()
	//userClaims.Exp = time.Now().Add(time.Hour).Unix()
	userClaims.Exp = time.Now().Unix() + -1
	jwtToken, e := GetNewJwtToken([]byte{0x23, 0x45, 0x34}, &userClaims)
	if e != nil {
		t.Error(e)
	} else {
		t.Log(jwtToken)
	}
	claims, e := CheckJwtToken([]byte{0x23, 0x45, 0x34}, jwtToken)
	if e != nil {
		t.Error(e)
		eV := e.(*jwt.ValidationError)
		if eV.Inner == ErrIatTime {
			t.Log("ok err time iat exp")
		} else if eV.Inner == ErrTimeExp {
			t.Log("ok err time exp")
		}
	} else {
		t.Log(claims)
	}
}

func TestCheckJwtToken2(t *testing.T) {
	claims, err := CheckJwtToken([]byte{0x34}, "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRfaWQiOiIiLCJ1c2VyX2lkIjoiMTI1OTQ5MTY2MTA1MDgxMDM2OCIsImlhdCI6MTU4OTIwMzYxOSwiZXhwIjoxNTg5ODA4NDE5LCJqdGkiOiI0N2Q3NzQ0NDcxYTk0Njk2YThlOWQ3MjM0MTljYjdmMSIsImlzcyI6InN5cyJ9.ggFsvQS5WOPLHImsPmiJswrVx6fE7HGrgq9KSIdHRo0")
	println(jsonutil.ToJson(claims))
	println(jsonutil.ToJson(err.Error()))
	err1 := err.(*jwt.ValidationError)
	err2 := err.(*jwt.ValidationError).Inner
	println(errors.Is(err2, ErrTimeExp))
	println(err1.Inner, err2)
}
