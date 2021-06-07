package api

import "golang.org/x/oauth2"

type GmailInfo struct {
	Id           uint          `json:"id"`
	Token        *oauth2.Token `json:"token"`
	ConnectionID uint          `json:"connectionId"`
}
