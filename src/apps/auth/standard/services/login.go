package services

import (
	accounts "hciengserver/src/apps/account/services"

	"golang.org/x/crypto/bcrypt"
)

func login(loginData *accounts.Account) (*accounts.Account, error) {
	accountFromDb, err := verifyUserCreds(loginData)
	if err != nil {
		return nil, err
	}
	return accountFromDb, nil
}

func verifyUserCreds(accountToValidate *accounts.Account) (*accounts.Account, error) {
	accountInDb, err := accounts.GetAccount(accountToValidate.EmailAddr)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(accountInDb.Pass), []byte(accountToValidate.Pass))
	if err != nil {
		return nil, nil
	}

	return accountInDb, nil
}

// this function takes some [loginData] (email and password or Google JWT) and
// retrieves the related account from the database
func GetAccount(loginData *accounts.Account) (*accounts.Account, error) {
	userAccount, err := login(loginData)
	if err != nil {
		return nil, err
	}

	return userAccount, nil
}
