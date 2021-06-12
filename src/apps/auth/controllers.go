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

	c.JSON(http.StatusOK, claims)
}
