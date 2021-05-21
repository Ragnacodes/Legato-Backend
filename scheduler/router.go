package scheduler

import (
	"fmt"
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

// Resolver includes all of our use cases to handle requests
func NewRouter() *gin.Engine {
	r := gin.Default()

	// Setup middlewares
	//r.Use(middleware.AuthMiddleware(&resolvers.UserUseCase))
	//r.Use(middleware.CORSMiddleware())

	for _, route := range schedulerRoutes {
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

	return r
}
