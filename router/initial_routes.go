package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"legato_server/env"
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
	_, err := http.Get(fmt.Sprintf("%s/api/ping", env.ENV.SchedulerUrl))
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
