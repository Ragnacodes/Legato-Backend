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
			connectionAuthUrl,
		},
		route{
			"Add a connection",
			POST,
			"users/:username/add/connection",
			addConnection,
		},
		route{
			"Get a connection",
			GET,
			"users/:username/get/connection/:id",
			returnConnection,
		},
		route{
			"Get user connections",
			GET,
			"users/:username/get/connections",
			GetConnections,
		},
		route{
			"Update Token",
			PUT,
			"users/:username/update/connection/token/:id",
			UpdateTokenFieldByID,
		},
		route{
			"Check Token",
			GET,
			"users/:username/check/connection/:id",
			checkConnection,
		},
		route{
			"Delete Token",
			DELETE,
			"users/:username/connection/delete/:id",
			DeleteConnectionByID,
		},
		route{
			"Ufpdate Token",
			PUT,
			"users/:username/update/connection/name/:id",
			UpdateConnectionNameByID,
		},
	},
}

func connectionAuthUrl(c *gin.Context) {
	service := c.Param("service")
	switch service {
	case "spotify":
		c.JSON(200, gin.H{
			"url": env.SpotifyAuthenticateUrl,
		})
	case "google":
		c.JSON(200, gin.H{
			"url": env.GoogleAuthenticateUrl,
		})
	case "github":
		c.JSON(200, gin.H{
			"url": env.GitAuthenticateUrl,
		})
	case "discord":
		c.JSON(200, gin.H{
			"url": env.DiscordAuthenticateUrl,
		})
	}
}

func addConnection(c *gin.Context) {
	// this function add connection

	username := c.Param("username")
	userToken := api.Connection{}
	_ = c.BindJSON(&userToken)

	// Authenticate
	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}
	connection, err := resolvers.UserUseCase.AddConnectionToDB(username, userToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not add token: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Name": connection.Name,
		"Data": connection.Data,
		"Id":   connection.ID,
	})
}

func returnConnection(c *gin.Context) {
	//this function return a connection
	userToken := api.Connection{}
	username := c.Param("username")
	id := c.Param("id")
	_ = c.BindJSON(&userToken)

	// Authenticate
	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}
	i, _ := strconv.Atoi(id)
	token, err := resolvers.UserUseCase.GetConnectionByID(username, uint(i))
	if err == nil && !strings.EqualFold(token.Data, "") {
		c.JSON(200, gin.H{
			"token": token.Data,
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
	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}
	type Connection struct {
		Name string
		Data string
		Id   uint
	}
	var Connections []Connection
	connections, err := resolvers.UserUseCase.GetConnections(username)
	for _, connection := range connections {
		con := Connection{}
		con.Id = connection.ID
		con.Name = connection.Name
		con.Data = connection.Data
		Connections = append(Connections, con)
	}
	if err == nil {
		c.JSON(200, gin.H{
			"connections": Connections,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not find connections: %s", err),
		})
	}
}

func UpdateConnectionNameByID(c *gin.Context) {
	// this function update name of connection with id
	userToken := api.Connection{}
	username := c.Param("username")
	id := c.Param("id")
	_ = c.BindJSON(&userToken)

	// Authenticate
	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}
	i, _ := strconv.Atoi(id)
	userToken.ID = i
	err := resolvers.UserUseCase.UpdateUserConnectionNameById(username, userToken)
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
	// Auth
	loginUser := checkAuth(c, []string{username})
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
	// Auth
	loginUser := checkAuth(c, []string{username})
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
	userToken := api.Connection{}
	username := c.Param("username")
	id := c.Param("id")
	_ = c.BindJSON(&userToken)

	// Authenticate
	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}
	i, _ := strconv.Atoi(id)
	userToken.ID = i
	err := resolvers.UserUseCase.UpdateTokenFieldByID(username, userToken)
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
