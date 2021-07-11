package usecase

import (
	"encoding/json"
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"log"
	"time"

	"golang.org/x/oauth2"
)

type SpotifyUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}
type connectionJson struct {
	Connection  *uint `json:"connection"`
}

func NewSpotifyUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.SpotifyUseCase {
	return &SpotifyUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (sp *SpotifyUseCase)GetUserToken(cid int) (token *oauth2.Token, err error){
	tokenString, err :=  sp.db.GetSpotifyTokenByConnectionID(cid)
	if err!= nil{
		return nil, err
	}
	err = json.Unmarshal([]byte(tokenString), &token)
	if err != nil {
		log.Println(err)
	}
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
	result := make(map[string]interface{}) 
	err = json.Unmarshal([]byte(spotify.Service.Data), &result)
	if err == nil {
		
		if _, ok := result["connection"]; ok {
			var data connectionJson
			err = json.Unmarshal([]byte(spotify.Service.Data), &data)
			if err != nil {
				return err
			}
			spotify.ConnectionID = data.Connection
		}
	}
		err = sp.db.UpdateSpotify(&scenario, serviceId, spotify)
		if err != nil {
			return err
	}

	return nil
}
