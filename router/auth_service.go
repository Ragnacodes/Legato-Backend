package router

import (
	"fmt"
	"legato_server/api"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// A testing scenario
var ConnectionRG = routeGroup{
	"Connection routers",
	routes{
		route{
			"Connection urls",
			POST,
			"user/connection/access/token/urls",
			connection_auth_url,
		},
		route{
			"Add Connection",
			POST,
			"users/:username/add/connection",
			addConnection,
		},
		route{
			"Retrun connection",
			GET,
			"users/:username/get/connection/:id",
			returnConnection,
		},
		route{
			"Return Connections",
			GET,
			"users/:username/get/connections",
			GetConnections,
		},
		route{
			"update Token",
			PUT,
			"users/:username/update/connection/token/:id",
			UpdateTokenFieldByID,
		},
		route{
			"check Token",
			GET,
			"users/:username/check/connection/:id",
			checkConnection,
		},
		route{
			"delete Token",
			DELETE,
			"users/:username/connection/delete/:id",
			DeleteConnectionByID,
		},
		route{
			"update Token",
			PUT,
			"users/:username/update/connection/name/:id",
			UpdateConnectionNameByID,
		},
	},
}

const spotify_authenticate_url = "https://accounts.spotify.com/authorize?client_id=74049abbf6784599a1564060e7c9dc12&redirect_uri=http://localhost:3000/redirect/spotify/&response_type=code&scope=user-read-private&state=abc123"
const google_authenticate_url = "https://accounts.google.com/o/oauth2/v2/auth?client_id=906955768602-u0nu3ruckq6pcjvune1tulkq3n0kfvrl.apps.googleusercontent.com&response_type=code&scope=https://www.googleapis.com/auth/gmail.readonly&redirect_uri=http://localhost:3000/redirect/gmail/&access_type=offline"
const git_authenticate_url = "https://github.com/login/oauth/authorize?access_type=online&client_id=Iv1.9f22bc1a9e8e6822&response_type=code&scope=user%3Aemail+repo&state=thisshouldberandom&redirect_uri=http://localhost:3000/redirect/github/"
const discord_authenticate_url = "https://discord.com/api/oauth2/authorize?access_type=online&client_id=830463353079988314&redirect_uri=http://localhost:3000/redirect/discord/&response_type=code&scope=identify+email&state=h8EecvhXJqHsG5EQ3K0gei4EUrWpaFj_HqH3WNZdrzrX1BX1COQRsTUv3-yGi3WmHQbw0EHJ58Rx1UOkvwip-Q%3D%3D"

func connection_auth_url(c *gin.Context) {
	usertoken := api.Connection{}
	_ = c.BindJSON(&usertoken)
	if strings.EqualFold(usertoken.TokenType, "spotify") {
		c.JSON(200, gin.H{
			"url": spotify_authenticate_url,
		})
	}
	if strings.EqualFold(usertoken.TokenType, "google") {
		c.JSON(200, gin.H{
			"url": google_authenticate_url,
		})
	}
	if strings.EqualFold(usertoken.TokenType, "git") {
		c.JSON(200, gin.H{
			"url": git_authenticate_url,
		})
	}
	if strings.EqualFold(usertoken.TokenType, "discord") {
		c.JSON(200, gin.H{
			"url": discord_authenticate_url,
		})
	}
}

func addConnection(c *gin.Context) {
	// this function add connection

	username := c.Param("username")
	usertoken := api.Connection{}
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "addtoken")
	if loginUser == nil {
		return
	}

	connection, err := resolvers.UserUseCase.AddConnectionToDB(username, usertoken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not add token: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"connection": connection,
	})
}

func returnConnection(c *gin.Context) {
	//this function retuen a connection
	usertoken := api.Connection{}
	username := c.Param("username")
	id := c.Param("id")
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "gettoken")
	if loginUser == nil {
		return
	}
	i, _ := strconv.Atoi(id)
	token, err := resolvers.UserUseCase.GetConnectionByID(username, uint(i))
	if err == nil && !strings.EqualFold(token.Token, "") {
		c.JSON(200, gin.H{
			"token": token.Token,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not find connection : %s", err),
		})
	}
}

func GetConnections(c *gin.Context) {
	// this function return list of all connections of a user
	username := c.Param("username")
	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "gettoken")
	if loginUser == nil {
		return
	}
	connections, err := resolvers.UserUseCase.GetConnections(username)

	if err == nil {
		c.JSON(200, gin.H{
			"connections": connections,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not find connections: %s", err),
		})
	}
}

func UpdateConnectionNameByID(c *gin.Context) {
	// this function update name of connection with id
	usertoken := api.Connection{}
	username := c.Param("username")
	id := c.Param("id")
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "updatetoken")
	if loginUser == nil {
		return
	}
	i, _ := strconv.Atoi(id)
	usertoken.ID = i
	err := resolvers.UserUseCase.UpdateUserConnectionNameById(username, usertoken)
	if err == nil {
		c.JSON(200, gin.H{
			"message": "update connection successfully",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not update connection: %s", err),
		})
	}
}

func checkConnection(c *gin.Context) {
	// this function check if there is a connection with this id for a user
	username := c.Param("username")
	id := c.Param("id")
	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "checktoken")
	if loginUser == nil {
		return
	}
	i, _ := strconv.Atoi(id)
	err := resolvers.UserUseCase.CheckConnectionByID(username, uint(i))
	if err == nil {
		c.JSON(200, gin.H{
			"message": "true",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "false",
		})
	}
}
func DeleteConnectionByID(c *gin.Context) {
	// this function delete a connection with id
	username := c.Param("username")
	id := c.Param("id")
	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "deletetoken")
	if loginUser == nil {
		return
	}
	i, _ := strconv.Atoi(id)
	err := resolvers.UserUseCase.DeleteUserConnectionById(username, uint(i))
	if err == nil {
		c.JSON(200, gin.H{
			"message": "deleted connection successfully",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not delete connection: %s", err),
		})
	}
}

func UpdateTokenFieldByID(c *gin.Context) {
	// this function update token field in connection with id
	usertoken := api.Connection{}
	username := c.Param("username")
	id := c.Param("id")
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "updateotoken")
	if loginUser == nil {
		return
	}
	i, _ := strconv.Atoi(id)
	usertoken.ID = i
	err := resolvers.UserUseCase.UpdateTokenFieldByID(username, usertoken)
	if err == nil {
		c.JSON(200, gin.H{
			"message": "update connection successfully",
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not update connection: %s", err),
		})
	}
}
