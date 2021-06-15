package oauth

import (
	"errors"
	accounts "hciengserver/src/apps/account/services"
	"hciengserver/src/jwt"

	googleAuthIDTokenVerifier "github.com/futurenda/google-auth-id-token-verifier"
)

func oauthLogin(googleJwt string) (*accounts.Account, error) {
	claims, err := getClaimsFromJwt(googleJwt)
	if err != nil {
		return nil, err
	}

	account, err := accounts.GetAccount(claims.Email)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func getAccountFromGoogleJwt(tkn string) (*accounts.Account, error) {
	userAccount, err := oauthLogin(tkn)
	if err != nil {
		if accounts.AccountIsAbsent(err) {
			return nil, errors.New("unauthorized")
		}
		return nil, err
	}

	return userAccount, nil
}

func getClaimsFromJwt(tkn string) (*googleAuthIDTokenVerifier.ClaimSet, error) {
	claims, err := jwt.ValidateGoogleJWT(tkn)
	if err != nil {
		return claims, err
	}
	return claims, nil
}

func getRegisterDataFromGoogleJwt(tkn string) (*accounts.Account, error) {
	claims, err := getClaimsFromJwt(tkn)
	if err != nil {
		return nil, err
	}

	return registerDataFromOauthClaims(claims), nil
}

func registerDataFromOauthClaims(claims *googleAuthIDTokenVerifier.ClaimSet) *accounts.Account {
	return &accounts.Account{
		EmailAddr: claims.Email,
	}
}
