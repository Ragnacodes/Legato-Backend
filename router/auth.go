package router

import (
	"github.com/gin-gonic/gin"
	"legato_server/middleware"
	"legato_server/models"
	"net/http"
)

const Authorization = "Authorization"

// authRG includes all of the routes that is related to
// signing up or authorizing a user.
var authRG = routeGroup{
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
	loginUser := checkAuth(c, nil)
	if loginUser == nil {
		return
	}

	users, _ := resolvers.UserUseCase.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

/*
	Helper functions
*/

// checkAuth was written because of DRY (Don't Repeat Yourself).
// Each time it authenticate the user and handle the errors that might occur.
// validUsers is list of usernames that the api is accessible for them.
// nil validUsers means that any authenticated user can use api.
// Return the logged in user.
func checkAuth(c *gin.Context, validUsers []string) *models.UserInfo {
	// Get the user
	rawData := c.MustGet(middleware.UserKey)
	if rawData == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
		return nil
	}

	loginUser := rawData.(*models.UserInfo)
	if loginUser == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
		return nil
	}

	// Check if validUsers is nil
	// If it is nil it means any authenticated user is accepted.
	if validUsers == nil {
		return loginUser
	}
	// If it isn't, Check if the user has access.
	for _, validUser := range validUsers {
		if loginUser.Username == validUser {
			return loginUser
		}
	}

	// If the api is not accessible
	c.JSON(http.StatusForbidden, gin.H{
		"message": "access denied: can not add scenario for this user",
	})
	return nil
}
