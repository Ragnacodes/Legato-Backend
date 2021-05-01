package router

import (
	"fmt"
	"legato_server/domain"
	"legato_server/middleware"

	"github.com/gin-gonic/gin"
)

const API = "api"

// Define methods as an integer.
type Method int

// Method enums
// POST, GET, PUT, DELETE methods.
const (
	POST   Method = iota
	GET    Method = iota
	PUT    Method = iota
	DELETE Method = iota
	PATCH  Method = iota
)

// Each route is a single api.
// name may use in logging and testing.
// method is an enum of request methods.
// pattern is the api address.
// handlerFunc is gin.HandlerFunc that gets *gin.Context.
type route struct {
	name        string
	method      Method
	pattern     string
	handlerFunc gin.HandlerFunc
}

// routes is an array of route.
type routes []route

// routeGroup is each single separated scenario of apis.
// name is the scenario name and we can use it in testing and logging.
// routes is array of related apis to that scenario.
type routeGroup struct {
	name   string
	routes routes
}

// routeGroups includes many scenarios of app.
type routeGroups []routeGroup

// Resolver includes all of our use cases to handle requests
type Resolver struct {
	UserUseCase     domain.UserUseCase
	ScenarioUseCase domain.ScenarioUseCase
	ServiceUseCase  domain.ServiceUseCase
	WebhookUseCase  domain.WebhookUseCase
	HttpUserCase    domain.HttpUseCase
	TelegramUseCase domain.TelegramUseCase
	SpotifyUseCase  domain.SpotifyUseCase
}

// This Resolver includes all of our use cases so we can handle incoming requests
var resolvers *Resolver

// Use all of your scenarios for the server here in legatoRoutesGroups
var legatoRoutesGroups = routeGroups{
	initialRG,
	authRG,
	scenarioRG,
	webhookRG,
	nodeRG,
	httpRG,
	spotifyRG,
	ConnectionRG,
}

// NewRouter get the resolvers and create *gin.Engine that can handle all
// of the request and responses.
func NewRouter(res *Resolver) *gin.Engine {
	resolvers = res

	r := gin.Default()

	// Setup middlewares
	r.Use(middleware.AuthMiddleware(&resolvers.UserUseCase))
	r.Use(middleware.CORSMiddleware())

	for _, routers := range legatoRoutesGroups {
		for _, route := range routers.routes {
			pattern := fmt.Sprintf("/%s/%s", API, route.pattern)

			switch route.method {
			case GET:
				r.GET(pattern, route.handlerFunc)
				break
			case PATCH:
				r.PATCH(pattern, route.handlerFunc)
				break
			case POST:
				r.POST(pattern, route.handlerFunc)
				break
			case PUT:
				r.PUT(pattern, route.handlerFunc)
				break
			case DELETE:
				r.DELETE(pattern, route.handlerFunc)
				break
			}
		}
	}

	return r
}
