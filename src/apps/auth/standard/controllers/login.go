package controllers

import (
	"hciengserver/src/apps/account/services"
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginData := new(accounts.Account)

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	userAccount, err := services.GetAccount(loginData.EmailAddr)
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
