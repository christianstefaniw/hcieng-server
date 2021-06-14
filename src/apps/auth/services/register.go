package services

import (
	"context"
	"errors"
	account "hciengserver/src/apps/account/services"
	bodyData "hciengserver/src/apps/auth/body_data"
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// adds the account data present in the register data (from request body)
// to the database. if the register data contains a jwt it will add the claims to the db
func AddNewRegisterToDb(registerData bodyData.RegisterData) error {
	registerData, err := genMissingInfo(registerData)
	if err != nil {
		return err
	}

	err = ValidateAndAddAccountToDb(registerData.Account)
	if err != nil {
		return err
	}

	return nil
}

// check that the new account is unique and add to database if it is
func ValidateAndAddAccountToDb(newAcccount *account.Account) error {
	_, err := account.GetAccount(newAcccount.EmailAddr)
	if err != nil {
		if account.AccountIsAbsent(err) {
			_, err = addAccountToDb(newAcccount)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return errors.New("user already exists")
}

func addAccountToDb(accountData *account.Account) (*mongo.InsertOneResult, error) {
	return database.GetMongoDBConn().
		Client().Database(hciengserver.DB_NAME).
		Collection(hciengserver.ACCOUNT_COLL).
		InsertOne(context.Background(), accountData)
}

// add email data to registration if there is a jwt (meaning oauth was used)
// or generate a password hash if standard login was used
func genMissingInfo(registerData bodyData.RegisterData) (bodyData.RegisterData, error) {
	if registerData.HasJwt() {
		claims, err := getClaims(registerData.GoogleJWT)
		if err != nil {
			return registerData, err
		}
		registerData.EmailAddr = claims.Email
	} else {
		hash, err := bcrypt.GenerateFromPassword([]byte(registerData.Pass), 5)
		if err != nil {
			return registerData, err
		}
		registerData.Pass = string(hash)
	}
	return registerData, nil
}
