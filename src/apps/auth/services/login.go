package services

import (
	account "hciengserver/src/apps/account/services"
	bodyData "hciengserver/src/apps/auth/body_data"

	"golang.org/x/crypto/bcrypt"
)

func Login(loginData bodyData.LoginData) (*account.Account, error) {
	accountFromDb, err := verifyUserCreds(loginData.Account)
	if err != nil {
		return nil, err
	}
	return accountFromDb, nil
}

func verifyUserCreds(accountToValidate *account.Account) (*account.Account, error) {
	accountInDb, err := account.GetAccount(accountToValidate.EmailAddr)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(accountInDb.Pass), []byte(accountToValidate.Pass))
	if err != nil {
		return nil, nil
	}

	return accountInDb, nil
}
