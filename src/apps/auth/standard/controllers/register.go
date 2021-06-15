package controllers

import (
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/apps/auth/standard/services"
	"hciengserver/src/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	registerData := new(accounts.Account)

	if err := c.BindJSON(registerData); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err := services.AddNewRegisterToDb(registerData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	tkn, err := jwt.MakeJWT(registerData.EmailAddr)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	jwt.SetCookie(c, tkn)

	c.JSON(http.StatusOK, registerData)
}
