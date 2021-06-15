package middleware

import (
	account "hciengserver/src/apps/account/services"
	"hciengserver/src/hciengserver"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type claims struct {
	*account.Account
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

		c.Set("user", accountFromClaims(reqClaims))
	}
}

func accountFromClaims(newClaims *claims) *account.Account {
	return &account.Account{
		EmailAddr: newClaims.EmailAddr,
		FirstName: newClaims.FirstName,
		LastName:  newClaims.LastName,
	}
}
