package router

import (
	"github.com/gin-gonic/gin"
	"legato_server/domain"
)

// Each route is a single api.
// name may use in logging and testing.
// method is an enum of request methods.
// pattern is the api address.
// handlerFunc is gin.HandlerFunc that gets *gin.Context.
type route struct {
	name        string
	method      string
	pattern     string
	handlerFunc gin.HandlerFunc
}

// routes is an array of route.
type routes []route

// groupRoute is each single separated scenario of apis.
// name is the scenario name and we can use it in testing and logging.
// routes is array of related apis to that scenario.
type groupRoute struct {
	name   string
	routes routes
}

// groupRoutes includes all of the scenarios in our app.
type groupRoutes []groupRoute

// Resolver includes all of our use cases to handle requests
type Resolver struct {
	UserUseCase domain.UserUseCase
}

// This Resolver includes all of our use cases so we can handle incoming requests
var resolvers *Resolver

// Use all of your scenarios for the server here in legatoRoutesGroups
var legatoRoutesGroups = groupRoutes{
	initialRoutesGroups,
}

// NewRouter get the resolvers and create *gin.Engine that can handle all
// of the request and responses.
func NewRouter(res *Resolver) *gin.Engine {
	r := gin.Default()

	resolvers = res

	for _, routers := range legatoRoutesGroups {
		for _, route := range routers.routes {
			switch route.method {
			case "GET":
				r.GET(route.pattern, route.handlerFunc)
				break
			case "POST":
				r.POST(route.pattern, route.handlerFunc)
				break
			case "PUT":
				r.PUT(route.pattern, route.handlerFunc)
				break
			case "DELETE":
				r.DELETE(route.pattern, route.handlerFunc)
				break
			}
		}
	}

	return r
}
