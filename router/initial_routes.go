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
			"Ping Pong test for schedule server",
			GET,
			"ping-schedule",
			pingSchedule,
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

func pingSchedule(c *gin.Context) {
	_, err := http.Post("http://192.168.1.20:8090/api/schedule/scenario/1", "application/json", nil)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func getDefaultUser(c *gin.Context) {
	user, _ := resolvers.UserUseCase.GetUserByUsername("legato")
	c.JSON(http.StatusOK, user)
}
