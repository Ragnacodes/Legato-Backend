package router

import (
	"fmt"
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
			"/:username/logs/:scenario_id/",
			getScenarioHistoriesById,
		},
		route{
			"get a history message list",
			GET,
			"/:username/logs/:scenario_id/histories/:history_id",
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

	c.JSON(http.StatusOK, gin.H{
		"histories": historyList,
	})
}


func getHistoryLogsById(c *gin.Context){

	username := c.Param("username")
	historyID, err := strconv.Atoi(c.Param("scenario_id"))
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

	c.JSON(http.StatusOK, gin.H{
		"logs": logs,
	})

}