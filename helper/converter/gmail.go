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

func GmailDbToGitInfo(gmail *legatoDb.Gmail) api.GmailInfo {
	var g api.GmailInfo
	g.Id = gmail.ID
	data := TokenData{}
	_ = json.Unmarshal([]byte(gmail.Token), &data)
	convData := oauth2.Token{
		AccessToken: data.AccessToken,
		TokenType:   data.TokenType,
		Expiry:      data.Expiry,
	}
	g.Token = &convData
	g.ConnectionID = gmail.ConnectionID

	return g
}
