package main

import (
	"github.com/spf13/viper"
	"legato_server/authenticate"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	"legato_server/router"
	"legato_server/scenario/webhook"
	"legato_server/user/usecase"
	"log"
	"time"
)

var userUseCase domain.UserUseCase

func init() {
	// Load environment variables
	env.LoadEnv()

	// Generate random jwt key
	authenticate.GenerateRandomKey()

	// Connect to database
	appDB, err := legatoDb.Connect()
	if err != nil {
		panic(err)
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	// Use Cases
	userUseCase = usecase.NewUserUseCase(appDB, timeoutContext)

	// Defaults
	_ = userUseCase.CreateDefaultUser()

	// Test single scenario
	testScenario()
}

func testScenario() {
	log.Println("---------------------------")
	log.Println("Testing Scenario mode")

	var events []webhook.Event
	events = []webhook.Event{
		{
			Name: "This is event number 1",
		},
		{
			Name: "This is event number 2",
		},
	}

	handler := webhook.Handler{
		Events: events,
	}

	wh := &webhook.Webhook{
		Name:    "Webhook 1",
		Handler: handler,
	}

	log.Println(wh.Name)
	wh.Observe()
	log.Println("---------------------------")
}

func main() {
	// resolvers include all of our use cases
	resolvers := router.Resolver{
		UserUseCase: userUseCase,
	}

	_ = router.NewRouter(&resolvers).Run()
}
