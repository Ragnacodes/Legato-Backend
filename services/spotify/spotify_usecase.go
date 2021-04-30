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

func (tu *SpotifyUseCase)GetUserToken(userInfo api.UserInfo) (token *oauth2.Token, err error){
	tokendb, err :=  tu.db.GetSpotifyTokeByUserID(userInfo.ID)
	if err!= nil{
		return nil, err
	}
	token = converter.DbTokenToOauth2(tokendb)
	return token, nil
}

func (tu *SpotifyUseCase) AddToScenario(u *api.UserInfo, scenarioId uint, nt api.NewServiceNode) (api.ServiceNode, error) {
	user, err := tu.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := tu.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	var spotify legatoDb.Spotify
	spotify.Service = converter.NewServiceNodeToServiceDb(nt)
	// spotify.Token.Token = nt.Data.(map[string]interface{})["Token"].(string)

	h, err := tu.db.CreateSpotify(&scenario, spotify)
	if err != nil {
		return api.ServiceNode{}, err
	}

	return converter.ServiceDbToServiceNode(h.Service), nil
}

func (tu *SpotifyUseCase) CreateSpotifyToken(user api.UserInfo, token *oauth2.Token) (error){
	dbToken := converter.Oauth2ToDbToken(token)
	err := tu.db.NewSpotifyToken(user.ID, dbToken)
	if err != nil {
		return err
	}
	return nil
}

func (tu *SpotifyUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, nt api.NewServiceNode) error {
	user, err := tu.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	scenario, err := tu.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	var spotify legatoDb.Spotify
	spotify.Service = converter.NewServiceNodeToServiceDb(nt)
	

	err = tu.db.UpdateSpotify(&scenario, serviceId, spotify)
	if err != nil {
		return err
	}

	return nil
}
