package services

import (
	account "hciengserver/src/apps/account/services"
	"hciengserver/src/jwt"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
)

func OauthLogin(googleJwt string) (*account.Account, error) {
	claims, err := getClaims(googleJwt)
	if err != nil {
		return nil, err
	}

	account, err := account.GetAccount(claims.Email)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func getClaims(googleJwt string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	return jwt.ValidateGoogleJWT(googleJwt)
}
