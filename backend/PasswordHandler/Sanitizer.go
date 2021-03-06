package PasswordHandler

import (
	"errors"
	"net/mail"
	"strings"

	Defaults "github.com/ElpidioL/Poke-Web/Defaults"
)

func CheckPassword(password string) error {
	spaces := strings.Fields(password)
	if len(password) < 8 || len(password) > 100 || len(spaces) > 1 || strings.HasSuffix(password, " ") || strings.HasPrefix(password, " ") {
		return errors.New("password Invalid")
	}
	return nil
}
func CheckEmail(email string) error {
	if len(email) == 0 {
		return errors.New("email Invalid")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("email not valid")
	}
	return nil
}
func CheckName(name string) error {
	if len(name) < 6 || len(name) > 100 {
		return errors.New("name Invalid")
	}
	return nil
}

func Sanitizer(registerUser Defaults.Register) (Defaults.Register, error) {

	//basic verifications.
	if registerUser.Intent == "register" {
		err := CheckPassword(registerUser.Password)
		if err != nil {
			return registerUser, err
		}
		err = CheckEmail(registerUser.Email)
		if err != nil {
			return registerUser, err
		}
		err = CheckName(registerUser.Name)
		if err != nil {
			return registerUser, err
		}
		passwordHash, err := Hash(registerUser.Password)
		if err != nil {
			return registerUser, errors.New("fail to hash")
		}
		if registerUser.Intent == "register" {
			registerUser.Password = passwordHash
		}

	} else if registerUser.Intent == "login" {

		err := CheckPassword(registerUser.Password)
		if err != nil {
			return registerUser, err
		}
		err = CheckEmail(registerUser.Email)
		if err != nil {
			return registerUser, err
		}
	}
	return registerUser, nil
}
