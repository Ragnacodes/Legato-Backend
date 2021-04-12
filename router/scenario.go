package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"legato_server/api"
	"net/http"
)

var scenarioRG = routeGroup{
	name: "User Scenario",
	routes: routes{
		route{
			name:        "Add a user scenario",
			method:      POST,
			pattern:     "/users/:username/scenarios",
			handlerFunc: addScenario,
		},
		route{
			name:        "Get user scenarios",
			method:      GET,
			pattern:     "/users/:username/scenarios",
			handlerFunc: getUserScenarios,
		},
		route{
			name:        "update a single scenarios",
			method:      POST,
			pattern:     "/users/:username/scenarios/:scenario_id",
			handlerFunc: updateScenario,
		},
		route{
			name:        "Get all of details of a single scenarios",
			method:      GET,
			pattern:     "/users/:username/scenarios/:scenario_id",
			handlerFunc: getFullScenario,
		},
	},
}

func addScenario(c *gin.Context) {
	username := c.Param("username")

	newScenario := api.NewScenario{}
	_ = c.BindJSON(&newScenario)

	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Add scenario
	createdScenario, err := resolvers.ScenarioUseCase.AddScenario(loginUser, &newScenario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not create scenario: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "scenario created successfully.",
		"scenario": createdScenario,
	})
}

func getUserScenarios(c *gin.Context) {
	username := c.Param("username")

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Get scenarios
	briefUserScenarios, err := resolvers.ScenarioUseCase.GetUserScenarios(loginUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not fetch user scenarios: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"scenarios": briefUserScenarios,
	})
}

func getFullScenario(c *gin.Context) {
	username := c.Param("username")
	scenarioId := c.Param("scenario_id")

	//Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Get single scenario details
	scenario, err := resolvers.ScenarioUseCase.GetUserScenarioById(loginUser, scenarioId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not fetch this scenario: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"scenario": scenario,
	})
}

func updateScenario(c *gin.Context) {
	username := c.Param("username")
	scenarioId := c.Param("scenario_id")

	updatedScenario := api.FullScenarioGraph{}
	_ = c.BindJSON(&updatedScenario)

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Update that scenario
	err := resolvers.ScenarioUseCase.UpdateUserScenarioById(loginUser, scenarioId, updatedScenario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not update this scenario: %s", err),
		})
		return
	}

	scenario, err := resolvers.ScenarioUseCase.GetUserScenarioById(loginUser, scenarioId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not fetch this scenario: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "update scenario successfully",
		"scenario": scenario,
	})
}
