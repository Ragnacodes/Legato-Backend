package scheduler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

var schedulerRoutes = routes{
	{
		name:        "Schedule to start scenario",
		method:      POST,
		pattern:     "schedule/scenario/:scenario_id",
		handlerFunc: scheduleStartScenario,
	},
}

func scheduleStartScenario(c *gin.Context) {
	scenarioId, _ := strconv.Atoi(c.Param("scenario_id"))

	sss := NewStartScenarioSchedule{}
	_ = c.BindJSON(&sss)

	log.Printf("Request new schedule for scenairo %d in %+v", scenarioId, sss.ScheduledTime)

	// Adding to the main queue
	msg := Tasks[StartScenarioTask].WithArgs(context.Background(), scenarioId)
	msg.Delay = sss.ScheduledTime.Sub(sss.SystemTime)
	err := MainQueue.Add(msg)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(500, gin.H{
		"message": "Your scenario scheduled successfully.",
	})
}
