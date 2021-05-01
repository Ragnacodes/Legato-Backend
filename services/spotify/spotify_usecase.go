package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"

	"golang.org/x/oauth2"
)

type SpotifyUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewSpotifyUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.SpotifyUseCase {
	return &SpotifyUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (sp *SpotifyUseCase)GetUserToken(userInfo api.UserInfo) (token *oauth2.Token, err error){
	tokenDb, err :=  sp.db.GetSpotifyTokeByUserID(userInfo.ID)
	if err!= nil{
		return nil, err
	}
	token = converter.DbTokenToOauth2(tokenDb)
	return token, nil
}

func (sp *SpotifyUseCase) AddToScenario(u *api.UserInfo, scenarioId uint, ns api.NewServiceNode) (api.ServiceNode, error) {
	user, err := sp.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := sp.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	var spotify legatoDb.Spotify
	spotify.Service = converter.NewServiceNodeToServiceDb(ns)
	// spotify.Token.Token = ns.Data.(map[string]interface{})["Token"].(string)

	h, err := sp.db.CreateSpotify(&scenario, spotify)
	if err != nil {
		return api.ServiceNode{}, err
	}

	return converter.ServiceDbToServiceNode(h.Service), nil
}

func (sp *SpotifyUseCase) CreateSpotifyToken(user api.UserInfo, token *oauth2.Token) (error){
	dbToken := converter.Oauth2ToDbToken(token)
	err := sp.db.NewSpotifyToken(user.ID, dbToken)
	if err != nil {
		return err
	}
	return nil
}

func (sp *SpotifyUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, ns api.NewServiceNode) error {
	user, err := sp.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	scenario, err := sp.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	var spotify legatoDb.Spotify
	spotify.Service = converter.NewServiceNodeToServiceDb(ns)
	

	err = sp.db.UpdateSpotify(&scenario, serviceId, spotify)
	if err != nil {
		return err
	}

	return nil
}
