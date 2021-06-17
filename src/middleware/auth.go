package middleware

import (
	"context"
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/database"
	"hciengserver/src/hciengserver"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type claims struct {
	Id string
	jwt.StandardClaims
}

func WithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tknStr, err := c.Cookie("authtoken")
		if err != nil {
			if err == http.ErrNoCookie {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				c.AbortWithError(http.StatusBadRequest, err)
			}
			return
		}
		reqClaims := new(claims)

		tkn, err := jwt.ParseWithClaims(tknStr, reqClaims, func(tkn *jwt.Token) (interface{}, error) {
			return hciengserver.JWT_SECRET, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatus(http.StatusUnauthorized)
			} else {
				c.AbortWithError(http.StatusBadRequest, err)
			}
			return
		}

		if !tkn.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user, err := retrieveUserData(reqClaims.Id)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		c.Set("user", user)
	}
}

func retrieveUserData(id string) (*accounts.Account, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var userData accounts.Account
	query := bson.M{
		"_id": objId,
	}
	err = database.GetMongoDBConn().Client().
		Database(hciengserver.DB_NAME).
		Collection(hciengserver.ACCOUNT_COLL).
		FindOne(context.Background(), query).
		Decode(&userData)
	if err != nil {
		return nil, err
	}

	return &userData, nil
}
