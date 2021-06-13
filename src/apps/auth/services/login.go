package services

import (
	account "hciengserver/src/apps/account/services"

	"golang.org/x/crypto/bcrypt"
)

func Login(user *account.Account) (bool, error) {
	hasAccount, err := verifyUserCreds(user)
	if err != nil {
		return false, err
	}
	if !hasAccount {
		return false, err
	}
	return true, nil
}

func verifyUserCreds(accountToValidate *account.Account) (bool, error) {
	accountInDb, err := account.GetAccount(accountToValidate.EmailAddr)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(accountInDb.Pass), []byte(accountToValidate.Pass))
	if err != nil {
		return false, nil
	}

	return true, nil
}
