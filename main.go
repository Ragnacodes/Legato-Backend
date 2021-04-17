package main

import (
	"github.com/spf13/viper"
	"legato_server/authenticate"
	"legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	"legato_server/router"
	scenarioUC "legato_server/scenario/usecase"
	httpUC "legato_server/services/http"
	serviceUC "legato_server/services/usecase"
	webhookUC "legato_server/services/webhook"
	userUC "legato_server/user/usecase"
	"time"
)

var userUseCase domain.UserUseCase
var scenarioUseCase domain.ScenarioUseCase
var serviceUseCase domain.ServiceUseCase
var webhookUseCase domain.WebhookUseCase
var httpUseCase domain.HttpUseCase

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
	serviceUseCase = serviceUC.NewServiceUseCase(appDB, timeoutContext)
	webhookUseCase = webhookUC.NewWebhookUseCase(appDB, timeoutContext)
	httpUseCase = httpUC.NewHttpUseCase(appDB, timeoutContext)

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
		ServiceUseCase:  serviceUseCase,
		WebhookUseCase:  webhookUseCase,
		HttpUserCase:    httpUseCase,
	}

	_ = router.NewRouter(&resolvers).Run()
}
