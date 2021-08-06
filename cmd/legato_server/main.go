package main

import (
	"legato_server/authenticate"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	"legato_server/logging"
	"legato_server/router"
	scenarioUC "legato_server/scenario/usecase"
	discordUC "legato_server/services/discord"
	githubUC "legato_server/services/github"
	gmailUC "legato_server/services/gmail"
	httpUC "legato_server/services/http"
	spotifyUC "legato_server/services/spotify"
	sshUC "legato_server/services/ssh"
	telegramUC "legato_server/services/telegram"
	toolboxUC "legato_server/services/toolbox"
	serviceUC "legato_server/services/usecase"
	webhookUC "legato_server/services/webhook"
	userUC "legato_server/user/usecase"
	logUC  "legato_server/logging/usecase"
	"legato_server/cache"

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
var loggerUseCase domain.LoggerUseCase
var gmailUseCase domain.GmailUseCase
var githubUseCase domain.GitUseCase
var discordUseCase domain.DiscordUseCase
var toolBoxUseCase domain.ToolBoxUseCase

func init() {
	// Load environment variables
	env.LoadEnv()

	// Generate random jwt key
	authenticate.GenerateRandomKey()
	
	// Make server sent event 
	logging.SSE.Init()

	// Connect to redis
	cache.ConnectToRedis()
	
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
	loggerUseCase = logUC.NewLoggerUseCase(appDB, timeoutContext)
	sshUseCase = sshUC.NewSshUseCase(appDB, timeoutContext)
	gmailUseCase = gmailUC.NewGmailUseCase(appDB, timeoutContext)
	githubUseCase = githubUC.NewGithubUseCase(appDB, timeoutContext)
	discordUseCase = discordUC.NewDiscordUseCase(appDB, timeoutContext)
	toolBoxUseCase = toolboxUC.NewToolBoxUseCase(appDB, timeoutContext)

	// Defaults
	_ = userUseCase.CreateDefaultUser()
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
		LoggerUseCase:   loggerUseCase,
		GmailUseCase:    gmailUseCase,
		GithubUseCase:   githubUseCase,
		DiscordUseCase:  discordUseCase,
		ToolBoxUseCase:  toolBoxUseCase,
	}

	_ = router.NewRouter(&resolvers).Run()
}
