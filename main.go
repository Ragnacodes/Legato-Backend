package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	"legato_server/user/usecase"
	"net/http"
	"time"
)

func init() {
	// Load environment variables
	env.LoadEnv()

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

}

var userUseCase domain.UserUseCase

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		//userUseCase.GetUserByUsername()
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/auth/admin", func(c *gin.Context) {
		user, _ := userUseCase.GetUserByUsername("admin")
		c.JSON(http.StatusOK, user)
	})

	_ = r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
