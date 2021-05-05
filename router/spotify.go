package router

import (
	"context"
	"fmt"
	"legato_server/env"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)


var spotifyRG = routeGroup{
	name: "spotify",
	routes: routes{
		route{
			"log in spotify",
			GET,
			"/users/:username/spotify",
			loginSpotify,
		},

		route{
			"spotify auth callback",
			GET,
			"/callback",
			completeAuth,
		},

		route{
			"get user playlists",
			GET,
			"/users/:username/spotify/playlists",
			ReadUserPlaylists,
		},
		route{
			"get a track info",
			GET,
			"/services/spotify/track/:id",
			getTrackInfo,
		},
	},
}



var (
	redirectURI = getRedirectURI
	scopes = []string{spotify.ScopePlaylistModifyPrivate, spotify.ScopeUserReadPrivate, spotify.ScopePlaylistModifyPublic, spotify.ScopeUserTopRead}
	auth  = getAuth
	ch    = make(chan *oauth2.Token)
	state = "abc123"
)

func getRedirectURI() string{
	return fmt.Sprintf("%s/api/callback/", env.ENV.WebUrl)
}
 
func getAuth() spotify.Authenticator{
	return spotify.NewAuthenticator(redirectURI(), scopes...)
}

func ReadUserPlaylists(c *gin.Context) {
	username := c.Param("username")
	// 
	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}
	token, err := resolvers.SpotifyUseCase.GetUserToken(*loginUser)
	if err!= nil{
		c.JSON(http.StatusNotFound, gin.H{
				"message": err,
		})
		return
	}

	client := auth().NewClient(token)
	
	playLists, err := client.CurrentUsersPlaylists()
	if err!= nil{
		c.JSON(http.StatusNotFound, gin.H{
				"message": err,
		})
		return
	}
	
	c.JSON(http.StatusOK, playLists.Playlists)

}


func loginSpotify(c *gin.Context) {
	username := c.Param("username")

	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}
	fmt.Println(getRedirectURI())
	url := auth().AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	token := <-ch
	
	err := resolvers.SpotifyUseCase.CreateSpotifyToken(*loginUser, token)
	// use the token to get an authenticated client
	client := auth().NewClient(token)
	// use the client to make calls that require authorization
	user, err := client.CurrentUser()
	if err != nil {
		log.Fatal(err)
		return
	}
	message := fmt.Sprintf("You are logged in as: %s", user.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func completeAuth(c *gin.Context) {
	tok, err := auth().Token(state, c.Request)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": fmt.Sprintf("Couldn't get token"),
		})
		log.Fatal(err)
		return
	}
	// if st := c.GetString("state"); st != state {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"message": fmt.Sprintf("State mismatch"),
	// 	})
	// 	log.Fatalf("State mismatch: %s != %s\n", st, state)
	// 	return
	// }

	log.Println("Login Completed!")
	ch <- tok
}

func getTrackInfo(c *gin.Context) {

	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	
	trackId := c.Param("id")
	// handle track info
	
	trackInfo, err := client.GetTrack(spotify.ID(trackId))
	if err!= nil{
		c.JSON(http.StatusNotFound, gin.H{
				"message": err,
		})
		return
	}
	
	c.JSON(http.StatusOK, trackInfo)
}
