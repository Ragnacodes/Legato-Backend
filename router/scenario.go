package router

import (
	"fmt"
	"legato_server/api"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
			method:      PUT,
			pattern:     "/users/:username/scenarios/:scenario_id",
			handlerFunc: updateScenario,
		},
		route{
			name:        "Get all of details of a single scenarios",
			method:      GET,
			pattern:     "/users/:username/scenarios/:scenario_id",
			handlerFunc: getFullScenario,
		},
		route{
			name:        "Delete a single scenario with its services",
			method:      DELETE,
			pattern:     "/users/:username/scenarios/:scenario_id",
			handlerFunc: deleteScenario,
		},
		route{
			name:        "Start a scenario",
			method:      PATCH,
			pattern:     "/users/:username/scenarios/:scenario_id",
			handlerFunc: startScenario,
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
		"message":  "scenario is created successfully.",
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
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	//Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Get single scenario details
	scenario, err := resolvers.ScenarioUseCase.GetUserScenarioById(loginUser, uint(scenarioId))
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
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	updatedScenario := api.NewScenario{}
	_ = c.BindJSON(&updatedScenario)

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Update that scenario
	err := resolvers.ScenarioUseCase.UpdateUserScenarioById(loginUser, uint(scenarioId), updatedScenario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not update this scenario: %s", err),
		})
		return
	}

	scenario, err := resolvers.ScenarioUseCase.GetUserScenarioById(loginUser, uint(scenarioId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not fetch this scenario: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "scenario is updated successfully",
		"scenario": scenario,
	})
}

func deleteScenario(c *gin.Context) {
	username := c.Param("username")
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Delete that scenario
	err := resolvers.ScenarioUseCase.DeleteUserScenarioById(loginUser, uint(scenarioId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not delete this scenario: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scenario is deleted successfully",
	})
}

func startScenario(c *gin.Context) {
	username := c.Param("username")
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Start that scenario
	err := resolvers.ScenarioUseCase.StartScenario(loginUser, uint(scenarioId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not start this scenario: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scenario is started successfully",
	})
}
