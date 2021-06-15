package controllers

import (
	"fmt"
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/apps/auth/standard/services"
	"hciengserver/src/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	loginData := new(accounts.Account)

	if err := c.ShouldBindJSON(loginData); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	fmt.Println(loginData.Pass, "ok")

	userAccount, err := services.GetAccount(loginData)
	if err != nil {
		if err.Error() == "unauthorized" {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		c.AbortWithError(http.StatusInternalServerError, err)
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
