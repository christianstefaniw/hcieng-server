package controllers

import (
	"errors"
	accounts "hciengserver/src/apps/account/services"
	bodyData "hciengserver/src/apps/auth/body_data"
	"hciengserver/src/apps/auth/services"
	"hciengserver/src/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginData := bodyData.NewLoginData()

	c.ShouldBindJSON(&loginData)

	userAccount, err := getAccount(loginData)
	if err != nil {
		if err.Error() == "unauthorized" {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	tkn, err := jwt.MakeJWT(userAccount.EmailAddr)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jwt.SetCookie(c, tkn)

	c.JSON(http.StatusOK, userAccount)
}

// this function takes some [loginData] (email and password or Google JWT) and
// retrieves the related account from the database
func getAccount(loginData bodyData.LoginData) (*accounts.Account, error) {
	var userAccount *accounts.Account
	var err error

	if loginData.HasJwt() {
		userAccount, err = services.OauthLogin(loginData.GoogleJWT)
		if err != nil {
			if accounts.AccountIsAbsent(err) {
				return nil, errors.New("unauthorized")
			}
			return nil, err
		}
	} else {
		userAccount, err = services.Login(loginData)
		if err != nil {
			return nil, err
		}
	}

	return userAccount, nil
}
