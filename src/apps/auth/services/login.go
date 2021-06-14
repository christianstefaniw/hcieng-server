package services

import (
	"errors"
	accounts "hciengserver/src/apps/account/services"
	bodyData "hciengserver/src/apps/auth/body_data"

	"golang.org/x/crypto/bcrypt"
)

func login(loginData bodyData.LoginData) (*accounts.Account, error) {
	accountFromDb, err := verifyUserCreds(loginData.Account)
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
func GetAccount(loginData bodyData.LoginData) (*accounts.Account, error) {
	var userAccount *accounts.Account
	var err error

	if loginData.HasJwt() {
		userAccount, err = OauthLogin(loginData.GoogleJWT)
		if err != nil {
			if accounts.AccountIsAbsent(err) {
				return nil, errors.New("unauthorized")
			}
			return nil, err
		}
	} else {
		userAccount, err = login(loginData)
		if err != nil {
			return nil, err
		}
	}

	return userAccount, nil
}
