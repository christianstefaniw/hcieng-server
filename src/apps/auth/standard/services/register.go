package services

import (
	account "hciengserver/src/apps/account/services"

	"golang.org/x/crypto/bcrypt"
)

func AddNewRegisterToDb(registerData *account.Account) error {
	var err error

	registerData.Pass, err = hashPassword(registerData.Pass)
	if err != nil {
		return err
	}

	err = account.ValidateAndAddAccountToDb(registerData)
	if err != nil {
		return err
	}

	return nil
}

func hashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 5)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
