package router

import (
	"fmt"
	"legato_server/api"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var nodeRG = routeGroup{
	name: "Scenario Nodes",
	routes: routes{
		route{
			name:        "Add a node to the scenario",
			method:      POST,
			pattern:     "/users/:username/scenarios/:scenario_id/nodes",
			handlerFunc: addNode,
		},
		route{
			name:        "Update a node in the scenario",
			method:      PUT,
			pattern:     "/users/:username/scenarios/:scenario_id/nodes/:node_id",
			handlerFunc: updateNode,
		},
		route{
			name:        "Delete a node in the scenario",
			method:      DELETE,
			pattern:     "/users/:username/scenarios/:scenario_id/nodes/:node_id",
			handlerFunc: deleteNode,
		},
		route{
			name:        "Get details about a node in the scenario",
			method:      GET,
			pattern:     "/users/:username/scenarios/:scenario_id/nodes/:node_id",
			handlerFunc: getNode,
		},
		route{
			name:        "Get all of the nodes in the scenario",
			method:      GET,
			pattern:     "/users/:username/scenarios/:scenario_id/nodes",
			handlerFunc: getScenarioNodes,
		},
	},
}

func getScenarioNodes(c *gin.Context) {
	username := c.Param("username")
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
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
		"nodes": scenario.Services,
	})
}

func addNode(c *gin.Context) {
	username := c.Param("username")
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	newNode := api.NewServiceNode{}
	_ = c.BindJSON(&newNode)

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Service Switch
	// NOTE: handle other non-service state
	var err error
	var addedServ api.ServiceNode
	switch newNode.Type {
	case "webhooks":
		addedServ, err = resolvers.WebhookUseCase.AddWebhookToScenario(loginUser, uint(scenarioId), newNode)
		break
	case "https":
		addedServ, err = resolvers.HttpUserCase.AddToScenario(loginUser, uint(scenarioId), newNode)
		break
	case "telegrams":
		addedServ, err = resolvers.TelegramUseCase.AddToScenario(loginUser, uint(scenarioId), newNode)
		break
	case "spotifies":
		addedServ, err = resolvers.SpotifyUseCase.AddToScenario(loginUser, uint(scenarioId), newNode)
	default:
		break
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("can not create this node: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "node is created successfully.",
		"node":    addedServ,
	})
}

func getNode(c *gin.Context) {
	username := c.Param("username")
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))
	nodeId, _ := strconv.Atoi(c.Param("node_id"))

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	node, err := resolvers.ServiceUseCase.GetServiceNodeById(loginUser, uint(scenarioId), uint(nodeId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not fetch this node: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"node": node,
	})
}

func updateNode(c *gin.Context) {
	username := c.Param("username")
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))
	nodeId, _ := strconv.Atoi(c.Param("node_id"))

	newNode := api.NewServiceNode{}
	_ = c.BindJSON(&newNode)

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Get the existing service and get the type
	serv, err := resolvers.ServiceUseCase.GetServiceNodeById(loginUser, uint(scenarioId), uint(nodeId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not fetch this scenario before updating: %s", err),
		})
		return
	}

	// Service Switch
	// NOTE: handle other non-service state
	switch serv.Type {
	case "webhooks":
		err = resolvers.WebhookUseCase.Update(loginUser, uint(scenarioId), uint(nodeId), newNode)
		break
	case "https":
		err = resolvers.HttpUserCase.Update(loginUser, uint(scenarioId), uint(nodeId), newNode)
		break
	case "telegrams":
		err = resolvers.TelegramUseCase.Update(loginUser, uint(scenarioId), uint(nodeId), newNode)
		break
	default:
		break
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("can not create this node: %s", err),
		})
		return
	}

	updatedServ, err := resolvers.ServiceUseCase.GetServiceNodeById(loginUser, uint(scenarioId), uint(nodeId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not fetch this scenario after updating: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "node is updated successfully.",
		"node":    updatedServ,
	})
}

func deleteNode(c *gin.Context) {
	username := c.Param("username")
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))
	nodeId, _ := strconv.Atoi(c.Param("node_id"))

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	err := resolvers.ServiceUseCase.DeleteServiceNodeById(loginUser, uint(scenarioId), uint(nodeId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not delete this node: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "node is deleted successfully",
	})
}
