package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// A testing scenario
var initialRG = routeGroup{
	"Initial routers",
	routes{
		route{
			"Ping Pong Test",
			GET,
			"ping",
			ping,
		},
		route{
			"Get Default User",
			GET,
			"auth/admin",
			getDefaultUser,
		},
	},
}

func ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getDefaultUser(c *gin.Context) {
	user, _ := resolvers.UserUseCase.GetUserByUsername("admin")
	c.JSON(http.StatusOK, user)
}
