package converter

import (
	"legato_server/db"
	"golang.org/x/oauth2"
)

func Oauth2ToDbToken(token *oauth2.Token) legatoDb.Token{
	tk := legatoDb.Token{}
	tk.AccessToken = token.AccessToken
	tk.RefreshToken = token.RefreshToken
	tk.Expiry = token.Expiry
	tk.TokenType = token.TokenType

	return tk
}


func DbTokenToOauth2(token legatoDb.Token) *oauth2.Token{
	tk := oauth2.Token{}
	tk.AccessToken = token.AccessToken
	tk.RefreshToken = token.RefreshToken
	tk.Expiry = token.Expiry
	tk.TokenType = token.TokenType

	return &tk
}