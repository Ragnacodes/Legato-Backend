package scheduler

import (
	"context"
	"github.com/gin-gonic/gin"
	"legato_server/api"
	"log"
	"net/http"
	"strconv"
)

var schedulerRoutes = routes{
	{
		name:        "Health check",
		method:      GET,
		pattern:     "ping",
		handlerFunc: pingPong,
	},
	{
		name:        "Schedule to start scenario",
		method:      POST,
		pattern:     "schedule/scenario/:scenario_id",
		handlerFunc: scheduleStartScenario,
	},
}

func pingPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func scheduleStartScenario(c *gin.Context) {
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	sss := api.NewStartScenarioSchedule{}
	_ = c.BindJSON(&sss)

	log.Printf("Request new schedule for scenairo %d in %+v", scenarioId, sss.ScheduledTime)

	// Adding to the main queue
	msg := Tasks[StartScenarioTask].WithArgs(context.Background(), scenarioId, sss.Token)
	msg.Delay = sss.ScheduledTime.Sub(sss.SystemTime)
	err := MainQueue.Add(msg)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Your scenario scheduled successfully.",
	})
}
