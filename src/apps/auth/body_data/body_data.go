package bodyData

import (
	accounts "hciengserver/src/apps/account/services"
)

// I know these types are redundant but I like the
// extra clarity the names provide

type RegisterData struct {
	GoogleJWT string `json:"jwt"`
	*accounts.Account
}

func NewRegisterData() RegisterData {
	return RegisterData{
		Account: new(accounts.Account),
	}
}

type LoginData struct {
	GoogleJWT string `json:"jwt"`
	*accounts.Account
}

func NewLoginData() LoginData {
	return LoginData{
		Account: new(accounts.Account),
	}
}

func (r RegisterData) HasJwt() bool {
	return r.GoogleJWT != ""
}

func (l LoginData) HasJwt() bool {
	return l.GoogleJWT != ""
}
