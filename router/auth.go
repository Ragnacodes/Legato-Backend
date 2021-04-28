package router

import (
	"legato_server/api"
	"legato_server/middleware"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
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
		route{
			"Get logged in user",
			GET,
			"auth/user",
			getLoggedInUser,
		},
	},
}

// signup creates new users
func signup(c *gin.Context) {
	newUser := api.NewUser{}
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
	userCredentials := api.UserCredential{}
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
	loggedInUser := checkAuth(c, nil)
	if loggedInUser == nil {
		return
	}

	users, _ := resolvers.UserUseCase.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

// getLoggedInUser will get the token in the header.
// Returns loggedInUser
func getLoggedInUser(c *gin.Context) {
	loggedInUser := checkAuth(c, nil)
	if loggedInUser == nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": loggedInUser,
	})
}

/*
	Helper functions
*/

// checkAuth was written because of DRY (Don't Repeat Yourself).
// Each time it authenticate the user and handle the errors that might occur.
// validUsers is list of usernames that the api is accessible for them.
// nil validUsers means that any authenticated user can use api.
// Return the logged in user.
func checkAuth(c *gin.Context, validUsers []string) *api.UserInfo {
	// Get the user
	rawData := c.MustGet(middleware.UserKey)
	if rawData == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
		return nil
	}

	loginUser := rawData.(*api.UserInfo)
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
		"message": "access denied: can not do any actions for this user",
	})
	return nil
}
func checkAuthForConnection(c *gin.Context, validUsers []string, request string) *api.UserInfo {
	// Get the user
	rawData := c.MustGet(middleware.UserKey)
	if rawData == nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied",
		})
		return nil
	}

	loginUser := rawData.(*api.UserInfo)
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
	if strings.EqualFold(request, "addtoken") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied: can not add connection for this user",
		})
	}
	if strings.EqualFold(request, "gettoken") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied: can not get connection for this user",
		})
	}
	if strings.EqualFold(request, "tokenurl") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied: can not go to the this url for this user",
		})
	}
	if strings.EqualFold(request, "tokenurl") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied: can not update this connection for this user",
		})
	}
	if strings.EqualFold(request, "checktoken") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied: can not find connection with this ID for this user",
		})
	}
	if strings.EqualFold(request, "deletetoken") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied: can not delete connection with this ID for this user",
		})
	}
	if strings.EqualFold(request, "updatetoken") {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "access denied: can not update connection with this ID for this user",
		})
	}
	return nil
}
