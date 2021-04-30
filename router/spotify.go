package router

import (
	"fmt"
	"legato_server/env"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
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
			"spotify auth callback",
			GET,
			"/users/:username/spotify/playlists",
			ReadUserPlaylists,
		},
	},
}



var (
	auth  = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadPrivate)
	ch    = make(chan *oauth2.Token)
	state = "abc123"
	redirectURI = fmt.Sprintf("%s/api/callback/", env.ENV.WebUrl)
)

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

	client := auth.NewClient(token)
	
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

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	token := <-ch
	
	err := resolvers.SpotifyUseCase.CreateSpotifyToken(*loginUser, token)
	// use the token to get an authenticated client
	client := auth.NewClient(token)
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
	tok, err := auth.Token(state, c.Request)
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