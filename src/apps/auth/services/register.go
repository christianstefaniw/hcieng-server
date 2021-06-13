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

func AddAccountToDb(registerData bodyData.RegisterData) error {
	registerData, err := genMissingInfo(registerData)
	if err != nil {
		return err
	}

	_, err = account.GetAccount(registerData.EmailAddr)
	if err != nil {
		if account.AccountIsAbsent(err) {
			_, err = addToDb(registerData.Account)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}
	return errors.New("user already exists")
}

func addToDb(accountData *account.Account) (*mongo.InsertOneResult, error) {
	return database.GetMongoDBConn().
		Client().Database(hciengserver.DB_NAME).
		Collection(hciengserver.ACCOUNT_COLL).
		InsertOne(context.Background(), accountData)
}

func genMissingInfo(registerData bodyData.RegisterData) (bodyData.RegisterData, error) {
	if !registerData.HasJwt() {
		hash, err := bcrypt.GenerateFromPassword([]byte(registerData.Pass), 5)
		if err != nil {
			return registerData, err
		}
		registerData.Pass = string(hash)
	} else {
		claims, err := getClaims(registerData.GoogleJWT)
		if err != nil {
			return registerData, err
		}
		registerData.EmailAddr = claims.Email
	}
	return registerData, nil
}
