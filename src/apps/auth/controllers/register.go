package controllers

import (
	bodyData "hciengserver/src/apps/auth/body_data"
	"hciengserver/src/apps/auth/services"
	"hciengserver/src/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	registerData := bodyData.NewRegisterData()

	if err := c.BindJSON(&registerData); err != nil {
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
