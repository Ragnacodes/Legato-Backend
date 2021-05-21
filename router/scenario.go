package router

import (
	"fmt"
	"legato_server/api"
	"legato_server/scheduler"
	"log"
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
		route{
			name:        "Schedule a scenario",
			method:      POST,
			pattern:     "/users/:username/scenarios/:scenario_id/schedule",
			handlerFunc: scheduleScenario,
		},
		route{
			name:        "Force a scenario to start",
			method:      POST,
			pattern:     "/scenarios/:scenario_id/force",
			handlerFunc: forceStartScenario,
		},
		route{
			name:        "Set an interval for a scenario",
			method:      PUT,
			pattern:     "/users/:username/scenarios/:scenario_id/set-interval",
			handlerFunc: setScenarioInterval,
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
		c.JSON(http.StatusBadRequest, gin.H{
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
	err := resolvers.ScenarioUseCase.StartScenarioInstantly(loginUser, uint(scenarioId))
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

func forceStartScenario(c *gin.Context) {
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	sss := api.NewStartScenarioSchedule{}
	_ = c.BindJSON(&sss)

	// Check if it is scheduler
	// It should be more secure later
	if c.GetHeader("Authorization") != scheduler.AccessToken {
		log.Println(c.GetHeader("Authorization"))
		log.Println(scheduler.AccessToken)
		c.JSON(http.StatusForbidden, gin.H{
			"message": "You do not have the access to force",
		})
		return
	}

	// Start that scenario because of the scheduler signal
	err := resolvers.ScenarioUseCase.ForceStartScenario(uint(scenarioId), sss.Token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not force start this scenario: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "scenario is started successfully",
	})
}

func scheduleScenario(c *gin.Context) {
	username := c.Param("username")
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	sss := api.NewStartScenarioSchedule{}
	_ = c.BindJSON(&sss)

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	log.Printf("Request new schedule for scenairo %d in %+v", scenarioId, sss.ScheduledTime)

	// Schedule that scenario
	err := resolvers.ScenarioUseCase.Schedule(loginUser, uint(scenarioId), &sss)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not schedule this scenario: %s", err),
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
		"message":  fmt.Sprintf("scenario is scheduled successfully for %v", sss.ScheduledTime),
		"scenario": scenario,
	})
}

func setScenarioInterval(c *gin.Context) {
	username := c.Param("username")
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	ni := api.NewScenarioInterval{}
	_ = c.BindJSON(&ni)

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Start that scenario
	err := resolvers.ScenarioUseCase.SetInterval(loginUser, uint(scenarioId), &ni)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not start this scenario: %s", err),
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
		"message":  fmt.Sprintf("interval has been set %d minutes", ni.Interval),
		"scenario": scenario,
	})
}
