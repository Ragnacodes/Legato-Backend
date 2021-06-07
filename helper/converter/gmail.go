package converter

import (
	"encoding/json"
	"legato_server/api"
	legatoDb "legato_server/db"

	"golang.org/x/oauth2"
)

func GmailInfoToGmailDb(g *api.GmailInfo) legatoDb.Gmail {
	var git legatoDb.Gmail
	s, _ := json.Marshal(g.Token)
	git.Token = string(s)
	return git
}

func GmailDbToGitInfo(git *legatoDb.Gmail) api.GmailInfo {
	var g api.GmailInfo
	g.Id = git.ID
	data := TokenData{}
	json.Unmarshal([]byte(git.Token), &data)
	condata := oauth2.Token{
		AccessToken: data.AccessToken,
		TokenType:   data.TokenType,
		Expiry:      data.Expiry,
	}
	g.Token = &condata
	g.ConnectionID = git.ConnectionID

	return g
}
