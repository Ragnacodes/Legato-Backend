package router

import (
	"github.com/gin-gonic/gin"
	"legato_server/middleware"
	"legato_server/models"
	"net/http"
)

const Authorization = "Authorization"

// authRoutesGroup includes all of the routes that is related to
// signing up or authorizing a user.
var authRoutesGroup = groupRoute{
	name: "Authentication",
	routes: routes{
		route{
			"Signup",
			POST,
			"auth/signup",
			signup,
		},
		route{
			"Login",
			POST,
			"auth/login",
			login,
		},
		route{
			"Refresh",
			POST,
			"auth/refresh",
			refresh,
		},
		route{
			"Protected for testing",
			GET,
			"auth/protected",
			protectedPage,
		},
	},
}

// signup creates new users
func signup(c *gin.Context) {
	newUser := models.NewUser{}
	_ = c.BindJSON(&newUser)

	err := resolvers.UserUseCase.RegisterNewUser(newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user created successfully.",
	})
}

// login is for authorizing and generating access token
func login(c *gin.Context) {
	userCredentials := models.UserCredential{}
	_ = c.BindJSON(&userCredentials)

	token, err := resolvers.UserUseCase.Login(userCredentials)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": token,
	})
}

// refresh is a api that grab the access token from header and generate a new one
func refresh(c *gin.Context) {
	token := c.GetHeader(Authorization)

	t, err := resolvers.UserUseCase.RefreshUserToken(token)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, t)
}

// protectedPage is a test api for getting all of the user details
func protectedPage(c *gin.Context) {
	rawData := c.MustGet(middleware.UserKey)
	if rawData == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
		return
	}

	loginUser := rawData.(*models.UserInfo)
	if loginUser == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
		return
	}

	users, _ := resolvers.UserUseCase.GetAllUsers()
	c.JSON(http.StatusOK, users)
}
