package services

import (
	account "hciengserver/src/apps/account/services"

	"golang.org/x/crypto/bcrypt"
)

// adds the account data present in the register data (from request body)
// to the database. if the register data contains a jwt it will add the claims to the db
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
