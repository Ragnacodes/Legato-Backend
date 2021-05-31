package router

import (
	"fmt"
	"legato_server/api"
	"legato_server/logging"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


var logRG = routeGroup{
	name: "logging",
	routes: routes{
		route{
			"event stream",
			GET,
			"/events/:scid",
			eventHandler,
		},
		route{
			"get history list",
			GET,
			"/users/:username/logs/:scenario_id/",
			getScenarioHistoriesById,
		},
		route{
			"get a history message list",
			GET,
			"/users/:username/logs/:scenario_id/histories/:history_id",
			getHistoryLogsById,
		},
	},
}


func eventHandler(c *gin.Context) {
	logging.SSE.EventServer.ServeHTTP(c.Writer, c.Request)
}



func getScenarioHistoriesById(c *gin.Context) {
	username := c.Param("username")
	scid, err := strconv.Atoi(c.Param("scenario_id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	historyList, err := resolvers.LoggerUseCase.GetScenarioHistoriesById(uint(scid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not get scenario histories: %s", err),
		})
		return
	}
	scenario, err := resolvers.ScenarioUseCase.GetUserScenarioById(loginUser, uint(scid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not get scenario detail: %s", err),
		})
		return
	}
	scenarioJson := api.ScenarioDetail{
		ID:       scenario.ID,
		Name:     scenario.Name,
		IsActive: scenario.IsActive,
	}
	if historyList == nil{
		response := []int{}
		c.JSON(http.StatusOK, gin.H{
			"scenario" : scenarioJson,
			"histories": response,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"scenario" : scenarioJson,
		"histories": historyList,
		
	})
}


func getHistoryLogsById(c *gin.Context){

	username := c.Param("username")
	historyID, err := strconv.Atoi(c.Param("history_id"))
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s", err),
		})
		return
	}
	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	logs, err := resolvers.LoggerUseCase.GetHistoryLogsById(uint(historyID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not get history logs: %s", err),
		})
		return
	}

	history, err := resolvers.LoggerUseCase.GetHistoryById(uint(historyID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not get history detail: %s", err),
		})
		return
	}

	if logs == nil{
		response := []int{}
		c.JSON(http.StatusOK, gin.H{
			"history":history,
			"logs": response,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"history":history,
		"logs": logs,
	})

}