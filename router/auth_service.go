package router

import (
	"fmt"
	"legato_server/api"
	"legato_server/env"
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
			"Connection access urls",
			GET,
			"user/connection/access/token/:service",
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

func connection_auth_url(c *gin.Context) {
	service := c.Param("service")
	if strings.EqualFold(service, "spotify") {
		c.JSON(200, gin.H{
			"url": env.Spotify_authenticate_url,
		})
	}
	if strings.EqualFold(service, "google") {
		c.JSON(200, gin.H{
			"url": env.Google_authenticate_url,
		})
	}
	if strings.EqualFold(service, "github") {
		c.JSON(200, gin.H{
			"url": env.Git_authenticate_url,
		})
	}
	if strings.EqualFold(service, "discord") {
		c.JSON(200, gin.H{
			"url": env.Discord_authenticate_url,
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
