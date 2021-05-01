package router

import "github.com/gin-gonic/gin"

const Http = "Http"

var httpRG = routeGroup{
	name: Http,
	routes: routes{
		route{
			"Create new http",
			POST,
			"/users/:username/services/http",
			createHttpService,
		},
		route{
			"Update http",
			PUT,
			"/users/:username/services/http/:http_id",
			updateHttpService,
		},
	},
}

func createHttpService(c *gin.Context) {

}

func updateHttpService(c *gin.Context) {

}
