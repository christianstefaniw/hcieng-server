package auth

import (
	"hciengserver/src/apps/auth/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	type parameters struct {
		GoogleJWT string `json:"jwt"`
	}

	var params parameters
	if err := c.ShouldBindJSON(&params); err != nil {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}

	claims, err := services.ValidateGoogleJWT(params.GoogleJWT)
	if err != nil {
		c.AbortWithError(http.StatusForbidden, err)
		return
	}

	tkn, err := services.MakeJWT(claims.Email)
	if err != nil {
		c.AbortWithError(http.StatusForbidden, err)
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

	c.JSON(http.StatusOK, claims)
}
