package controllers

import (
	"errors"
	accounts "hciengserver/src/apps/account/services"
	bodyData "hciengserver/src/apps/auth/body_data"
	"hciengserver/src/apps/auth/services"
	"hciengserver/src/constants"
	"hciengserver/src/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var err error

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

	cookie := &http.Cookie{
		Name:     "authtoken",
		Value:    tkn,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, userAccount)
}

func getAccount(loginData bodyData.LoginData) (*accounts.Account, error) {
	var userAccount *accounts.Account
	var err error

	if loginData.HasJwt() {
		userAccount, err = services.OauthLogin(loginData.GoogleJWT)
		if err != nil {
			if constants.NO_DOC_FOUND_ERR == err.Error() {
				return nil, errors.New("unauthorized")
			}
			return nil, err
		}
	} else {
		ok, err := services.Login(loginData.Account)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, errors.New("unauthorized")
		}
	}

	return userAccount, nil
}
