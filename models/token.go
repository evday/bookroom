package models

type Token struct {
	Token string `json:"token"`
	Username string `json:"username"`
	Isadmin bool `json:"isadmin"`
}