package scheduler

import (
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
	log.Printf("Request new schedule on scenairo %d", scenarioId)
}
