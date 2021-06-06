package legatoDb

import (
	"encoding/json"
	"errors"
	"fmt"
	"legato_server/env"
	"log"
	"time"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

const spotifyType string = "spotifies"
const addTrackToPlaylist string = "addToPlaylist"
const getTopTracks string = "getTopTracks"
 
var (
	scopes = []string{spotify.ScopePlaylistModifyPrivate, spotify.ScopeUserReadPrivate}
	auth  = getAuth
	redirectURI = getRedirectURI
) 


func getRedirectURI() string{
	return fmt.Sprintf("%s/api/callback/", env.ENV.WebUrl)
}
 
func getAuth() spotify.Authenticator{
	return spotify.NewAuthenticator(redirectURI(), scopes...)
}

type Spotify struct {
	gorm.Model
	TokenID uint
	Token   Token
	Service Service `gorm:"polymorphic:Owner;"`
}

type Token struct {
	gorm.Model
	AccessToken string 
	TokenType string 
	RefreshToken string 
	Expiry time.Time 
	UserID uint
  	User  User
}

type addToPlaylistData struct {
	PlaylistId string   `json:"PlaylistId"`
	TrackId   string `json:"TrackId"`
}

func (sp *Spotify) String() string {
	return fmt.Sprintf("(@Spotify: %+v)", *sp)
}

// Database methods
func (ldb *LegatoDB) NewSpotifyToken(UserID uint, token Token) error{
	token.UserID = UserID
	err := ldb.db.Create(&token).Error
	if err!=nil {
		return err
	}
	return nil
}


func (ldb *LegatoDB)GetSpotifyTokeByUserID(userID uint) (tk Token, err error){
	err = ldb.db.Where(&Token{UserID:userID}).Find(&tk).Error
	if err!=nil{
		return Token{}, err
	}
	return tk, nil
}

func (ldb *LegatoDB) CreateSpotify(s *Scenario, spotify Spotify) (*Spotify, error) {
	var tk Token
	spotify.Service.UserID = s.UserID
	spotify.Service.ScenarioID = &s.ID
	ldb.db.Where(&Token{UserID:spotify.Service.UserID}).Find(&tk)
	spotify.TokenID = tk.ID
	ldb.db.Create(&spotify)
	ldb.db.Save(&spotify)

	return &spotify, nil
}


func (ldb *LegatoDB) UpdateSpotify(s *Scenario, servId uint, nsp Spotify) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var sp Spotify
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&sp).Error
	if err != nil {
		return err
	}
	if sp.Service.ID != servId {
		return errors.New("the spotify service is not in this scenario")
	}

	ldb.db.Model(&serv).Updates(nsp.Service)
	ldb.db.Model(&sp).Updates(nsp)

	if nsp.Service.ParentID == nil {
		legatoDb.db.Model(&serv).Select("parent_id").Update("parent_id", nil)
	}

	return nil
}


func (ldb *LegatoDB) GetSpotifyByService(serv Service) (*Spotify, error) {
	var t Spotify
	err := ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&t).Error
	if err != nil {
		return nil, err
	}
	if t.ID != uint(serv.OwnerID) {
		return nil, errors.New("the spotify service is not in this scenario")
	}

	return &t, nil
}

// Service Interface for Http
func (sp Spotify) Execute(...interface{}) {
	
	err := legatoDb.db.Preload("Service").Preload("Token").Find(&sp).Error
	if err != nil {
		panic(err)
	}
	SendLogMessage("*******Starting Spotify Service*******", *sp.Service.ScenarioID, nil)
	
	logData := fmt.Sprintf("Executing type (%s) : %s\n", spotifyType, sp.Service.Name)
	SendLogMessage(logData, *sp.Service.ScenarioID, nil)

	var nextData interface{}
	token := DbTokenToOauth2(sp.Token)
	client := auth().NewClient(&token)

	switch sp.Service.SubType {
		case addTrackToPlaylist:
			var data addToPlaylistData
			err = json.Unmarshal([]byte(sp.Service.Data), &data)
			if err != nil {
				log.Fatal(err)
			}
			addTrackToPlaylistHandler(&client, data)
			break

		case getTopTracks:

			nextData = getUserTopTracksHandler(&client)
			break
			
		default:
			break
	}

	
	SendLogMessage(nextData.(string), *sp.Service.ScenarioID, &sp.Service.ID)

	sp.Next(nextData)
}

func (sp Spotify) Post() {
	log.Printf("Executing type (%s) node in background : %s\n", spotifyType, sp.Service.Name)
}

func (sp Spotify) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&sp).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing \"%s\" Children \n", sp.Service.Name)

	for _, node := range sp.Service.Children {
		serv, err := node.Load()
		if err != nil {
			log.Println("error in loading services in Next()")
			return
		}
		serv.Execute()
	}

	logData := fmt.Sprintf("*******End of \"%s\"*******",sp.Service.Name)
	SendLogMessage(logData, *sp.Service.ScenarioID, nil)
}


func addTrackToPlaylistHandler(client *spotify.Client, data addToPlaylistData){
	
	_, err := client.AddTracksToPlaylist(spotify.ID(data.PlaylistId), spotify.ID(data.TrackId))
	if err!= nil{
		fmt.Println(err)
	}
}

func getUserTopTracksHandler(client *spotify.Client) *spotify.FullTrackPage{
	list, _ := client.CurrentUsersTopTracks()
	return list
}


func DbTokenToOauth2(token Token) oauth2.Token{
	tk := oauth2.Token{}
	tk.AccessToken = token.AccessToken
	tk.RefreshToken = token.RefreshToken
	tk.Expiry = token.Expiry
	tk.TokenType = token.TokenType

	return tk
}
