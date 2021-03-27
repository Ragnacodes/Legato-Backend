package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"legato_server/middleware"
	"legato_server/models"
	"net/http"
)

var scenarioRG = routeGroup{
	name: "User Scenario",
	routes: routes{
		route{
			name:        "Add a user scenario",
			method:      POST,
			pattern:     "scenarios",
			handlerFunc: addScenario,
		},
	},
}

func addScenario(c *gin.Context) {
	newScenario := models.NewScenario{}
	_ = c.BindJSON(&newScenario)

	// Get the user
	rawData := c.MustGet(middleware.UserKey)
	if rawData == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
		return
	}

	loginUser := rawData.(*models.UserInfo)
	if loginUser == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
		return
	}

	// Add scenario
	err := resolvers.ScenarioUseCase.AddScenario(loginUser, &newScenario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not create scenario: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scenario created successfully.",
	})
}
