package jwt

import (
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/hciengserver"
	"net/http"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

func MakeJWT(accont *accounts.Account) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":   accont.EmailAddr,
		"first":   accont.FirstName,
		"last":    accont.LastName,
		"isadmin": accont.Admin,
	})

	signedTkn, err := tkn.SignedString(hciengserver.JWT_SECRET)
	if err != nil {
		return "", err
	}
	return signedTkn, nil
}

func ValidateGoogleJWT(tokenString string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	verifier := googleAuthIDTokenVerifier.Verifier{}
	aud := "835439685490-8j1kg7tk53vhflhp5n9ifmrs164mmbom.apps.googleusercontent.com"
	err := verifier.VerifyIDToken(tokenString, []string{aud})
	if err != nil {
		return nil, err
	}

	claimSet, err := googleAuthIDTokenVerifier.Decode(tokenString)
	if err != nil {
		return nil, err
	}

	return claimSet, nil
}

func SetCookie(c *gin.Context, tkn string) {
	cookie := &http.Cookie{
		Name:     "authtoken",
		Value:    tkn,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(c.Writer, cookie)
}
