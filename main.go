package main

import (
	"github.com/spf13/viper"
	"legato_server/authenticate"
	"legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	"legato_server/router"
	scenarioUC "legato_server/scenario/usecase"
	webhookUC "legato_server/services/webhook"
	userUC "legato_server/user/usecase"
	"time"
)

var userUseCase domain.UserUseCase
var scenarioUseCase domain.ScenarioUseCase
var webhookUseCase domain.WebhookUseCase

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
	userUseCase = userUC.NewUserUseCase(appDB, timeoutContext)
	scenarioUseCase = scenarioUC.NewScenarioUseCase(appDB, timeoutContext)
	webhookUseCase = webhookUC.NewWebhookUseCase(appDB, timeoutContext)

	// Defaults
	_ = userUseCase.CreateDefaultUser()
	_ = scenarioUseCase.CreateDefaultScenario()

	// Test single scenario
	// go scenarioUseCase.TestScenario()
}

func main() {
	// resolvers include all of our use cases
	resolvers := router.Resolver{
		UserUseCase:     userUseCase,
		ScenarioUseCase: scenarioUseCase,
		WebhookUseCase:  webhookUseCase,
	}

	_ = router.NewRouter(&resolvers).Run()
}
