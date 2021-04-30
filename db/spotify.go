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

func (t *Spotify) String() string {
	return fmt.Sprintf("(@Spotify: %+v)", *t)
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


func (ldb *LegatoDB) UpdateSpotify(s *Scenario, servId uint, nt Spotify) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var t Spotify
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&t).Error
	if err != nil {
		return err
	}
	if t.Service.ID != servId {
		return errors.New("the spotify service is not in this scenario")
	}

	ldb.db.Model(&serv).Updates(nt.Service)
	ldb.db.Model(&t).Updates(nt)

	if t.Service.ParentID == nil {
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
func (t Spotify) Execute(...interface{}) {
	log.Println("*******Starting Spotify Service*******")

	err := legatoDb.db.Preload("Service").Preload("Token").Find(&t).Error
	if err != nil {
		panic(err)
	}
	var nextData interface{}
	log.Printf("Executing type (%s) : %s\n", spotifyType, t.Service.Name)
	token := DbTokenToOauth2(t.Token)
	client := auth().NewClient(&token)
	switch t.Service.SubType {
		case addTrackToPlaylist:
			var data addToPlaylistData
			err = json.Unmarshal([]byte(t.Service.Data), &data)
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

	t.Next(nextData)
}

func (t Spotify) Post() {
	log.Printf("Executing type (%s) node in background : %s\n", spotifyType, t.Service.Name)
}

func (t Spotify) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&t).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing \"%s\" Children \n", t.Service.Name)

	for _, node := range t.Service.Children {
		serv, err := node.Load()
		if err != nil {
			log.Println("error in loading services in Next()")
			return
		}
		serv.Execute()
	}

	log.Printf("*******End of \"%s\"*******", t.Service.Name)
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