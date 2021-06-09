package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func email(c *gin.Context) {
	emailData := emailData{
		Name:         c.PostForm("name"),
		EmailAddress: c.PostForm("email-address"),
		Message:      c.PostForm("message"),
	}

	if err := send(emailData); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"sent": true,
	})
}
