package controllers

import (
	bodyData "hciengserver/src/apps/auth/body_data"
	"hciengserver/src/apps/auth/services"
	"hciengserver/src/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginData := bodyData.NewLoginData()

	c.ShouldBindJSON(&loginData)

	userAccount, err := services.GetAccount(loginData)
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
