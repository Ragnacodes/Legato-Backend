package main

import (
	"legato_server/authenticate"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	"legato_server/router"
	scenarioUC "legato_server/scenario/usecase"
	githubUC "legato_server/services/github"
	httpUC "legato_server/services/http"
	spotifyUC "legato_server/services/spotify"
	sshUC "legato_server/services/ssh"
	telegramUC "legato_server/services/telegram"
	serviceUC "legato_server/services/usecase"
	webhookUC "legato_server/services/webhook"
	userUC "legato_server/user/usecase"

	"time"

	"github.com/spf13/viper"
)

var userUseCase domain.UserUseCase
var scenarioUseCase domain.ScenarioUseCase
var serviceUseCase domain.ServiceUseCase
var webhookUseCase domain.WebhookUseCase
var httpUseCase domain.HttpUseCase
var telegramUseCase domain.TelegramUseCase
var spotifyUseCase domain.SpotifyUseCase
var sshUseCase domain.SshUseCase
var githubUseCase domain.GitUseCase

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
	telegramUseCase = telegramUC.NewTelegramUseCase(appDB, timeoutContext)
	spotifyUseCase = spotifyUC.NewSpotifyUseCase(appDB, timeoutContext)
	sshUseCase = sshUC.NewHttpUseCase(appDB, timeoutContext)
	githubUseCase = githubUC.NewHttpUseCase(appDB, timeoutContext)

	// Defaults
	_ = userUseCase.CreateDefaultUser()

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
		TelegramUseCase: telegramUseCase,
		SpotifyUseCase:  spotifyUseCase,
		SshUseCase:      sshUseCase,
		GithubUseCase:   githubUseCase,
	}

	_ = router.NewRouter(&resolvers).Run()
}
