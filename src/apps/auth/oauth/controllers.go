package oauth

import (
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type oauthToken struct {
	GoogleJWT string `json:"jwt"`
}

func GoogleAuthLogin(c *gin.Context) {
	var googleOauthToken oauthToken

	if err := c.BindJSON(&googleOauthToken); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userAccount, err := getAccountFromGoogleJwt(googleOauthToken.GoogleJWT)
	if err != nil {
		if !isAuthorized(err) {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tkn, err := jwt.MakeJWT(userAccount)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jwt.SetCookie(c, tkn)

	c.JSON(http.StatusOK, userAccount)
}

func GoogleAuthRegister(c *gin.Context) {
	var googleOauthToken oauthToken

	if err := c.BindJSON(&googleOauthToken); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	registerData, err := getRegisterDataFromGoogleJwt(googleOauthToken.GoogleJWT)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if err := accounts.ValidateAndAddAccountToDb(registerData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	tkn, err := jwt.MakeJWT(registerData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jwt.SetCookie(c, tkn)
	c.Set("user", registerData)

	c.JSON(http.StatusOK, registerData)
}

func isAuthorized(err error) bool {
	return err.Error() == "unauthorized"
}
