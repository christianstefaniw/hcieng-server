package services

import (
	"context"
	"hciengserver/src/constants"
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"

	"go.mongodb.org/mongo-driver/bson"
)

type Account struct {
	EmailAddr string `json:"email" bson:"email"`
	Pass      string `json:"pass" bson:"pass"`
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
