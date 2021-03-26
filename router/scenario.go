package router

import "github.com/gin-gonic/gin"

var scenarioRG = routeGroup{
	name: "User Scenario",
	routes: routes{
		route{
			name:        "Add a user scenario",
			method:      POST,
			pattern:     "scenario/add",
			handlerFunc: nil,
		},
	},
}

func addUserScenario(c *gin.Context) {

}
