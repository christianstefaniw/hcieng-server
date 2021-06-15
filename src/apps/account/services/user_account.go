package services

import (
	"context"
	"errors"
	"hciengserver/src/constants"
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Account struct {
	EmailAddr string `json:"email" bson:"email"`
	FirstName string `json:"first" bson:"first"`
	LastName  string `json:"last" bson:"last"`
	Pass      string `json:"pass" bson:"pass"`
	Admin     bool   `json:"admin" bson:"admin"`
}

func CreateAccount(email, pass, first, last string, isAdmin bool) *Account {
	return &Account{
		EmailAddr: email,
		Pass:      pass,
		FirstName: first,
		LastName:  last,
		Admin:     isAdmin,
	}
}

// set default configurations for an account
func SetDefaults(account *Account) {
	account.Admin = false
}

func GetAccount(email string) (*Account, error) {
	var accountFromDb Account

	err := database.GetMongoDBConn().Client().Database(hciengserver.DB_NAME).
		Collection(hciengserver.ACCOUNT_COLL).
		FindOne(context.Background(), bson.M{
			"email": email,
		}).
		Decode(&accountFromDb)
	if err != nil {
		return nil, err
	}

	return &accountFromDb, nil
}

func AccountIsAbsent(err error) bool {
	return err.Error() == constants.NO_DOC_FOUND_ERR
}

func addAccountToDb(accountData *Account) (*mongo.InsertOneResult, error) {
	return database.GetMongoDBConn().
		Client().Database(hciengserver.DB_NAME).
		Collection(hciengserver.ACCOUNT_COLL).
		InsertOne(context.Background(), accountData)
}

// check that the new account is unique and add to database if it is
func ValidateAndAddAccountToDb(newAcccount *Account) error {
	_, err := GetAccount(newAcccount.EmailAddr)
	if err != nil {
		if AccountIsAbsent(err) {
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
