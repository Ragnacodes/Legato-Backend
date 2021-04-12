package router

import (
	"fmt"
	"legato_server/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// A testing scenario
var ConnectionRG = routeGroup{
	"Connection routers",
	routes{
		route{
			"spotiy token",
			GET,
			"user/connection/access/url",
			connection_auth_url,
		},
		route{
			"Add Token",
			POST,
			"users/:username/connections/addtoken",
			AddTokens,
		},
		route{
			"Get Token",
			POST,
			"users/:username/connection/gettoken",
			GetToken,
		},
		route{
			"Get Tokens",
			POST,
			"users/:username/connection/gettokens",
			GetTokens,
		},
		route{
			"update Token",
			PUT,
			"users/:username/connection/update/token/name",
			UpdateTokenNameByID,
		},
		route{
			"check Token",
			POST,
			"users/:username/connection/check/token",
			CheckToken,
		},
		route{
			"delete Token",
			POST,
			"users/:username/connection/delete/token",
			DeleteConnectionByID,
		},
		route{
			"update Token",
			PUT,
			"users/:username/connection/update/token/token",
			UpdateUserTokenByName,
		},
	},
}

const spotify_authenticate_url = "https://accounts.spotify.com/authorize?client_id=74049abbf6784599a1564060e7c9dc12&redirect_uri=http://localhost:8080/callback&response_type=code&scope=user-read-private&state=abc123"
const google_authenticate_url = "https://accounts.google.com/o/oauth2/v2/auth?client_id=906955768602-u0nu3ruckq6pcjvune1tulkq3n0kfvrl.apps.googleusercontent.com&response_type=code&scope=https://www.googleapis.com/auth/gmail.readonly&redirect_uri=http://localhost:8080/callback&access_type=offline"
const git_authenticate_url = "https://github.com/login/oauth/authorize?access_type=online&client_id=Iv1.9f22bc1a9e8e6822&response_type=code&scope=user%3Aemail+repo&state=thisshouldberandom"
const discord_authenticate_url = "https://discord.com/api/oauth2/authorize?access_type=online&client_id=830463353079988314&redirect_uri=http://localhost:8080/callback&response_type=code&scope=identify+email&state=h8EecvhXJqHsG5EQ3K0gei4EUrWpaFj_HqH3WNZdrzrX1BX1COQRsTUv3-yGi3WmHQbw0EHJ58Rx1UOkvwip-Q%3D%3D"

func connection_auth_url(c *gin.Context) {
	username := c.Param("username")
	loginUser := checkAuthforconnection(c, []string{username}, "addtoken")
	if loginUser == nil {
		return
	}
	usertoken := models.UserAddToken{}
	if strings.EqualFold(usertoken.Token_type, "spotify") {
		c.JSON(200, gin.H{
			"spotify_url": spotify_authenticate_url,
		})
	}
	if strings.EqualFold(usertoken.Token_type, "google") {
		c.JSON(200, gin.H{
			"google_url": google_authenticate_url,
		})
	}
	if strings.EqualFold(usertoken.Token_type, "git") {
		c.JSON(200, gin.H{
			"git_url": git_authenticate_url,
		})
	}
	if strings.EqualFold(usertoken.Token_type, "discord") {
		c.JSON(200, gin.H{
			"discord_url": discord_authenticate_url,
		})
	}
}

func AddTokens(c *gin.Context) {
	username := c.Param("username")
	usertoken := models.UserAddToken{}
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "addtoken")
	if loginUser == nil {
		return
	}

	// Add token
	err := resolvers.UserUseCase.AddTokenDB(username, usertoken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not add token: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "token added successfully.",
	})
}

func GetToken(c *gin.Context) {
	usertoken := models.UserGetToken{}
	username := c.Param("username")
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "gettoken")
	if loginUser == nil {
		return
	}
	token, err := resolvers.UserUseCase.GetTokenByUsername(username, usertoken)
	if err == nil {
		c.JSON(200, gin.H{
			"token": token.Token,
		})
	} else {
		c.JSON(200, gin.H{
			"error": "can not find token",
		})
	}
}

func GetTokens(c *gin.Context) {
	usertoken := models.UserGetToken{}
	username := c.Param("username")
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "gettoken")
	if loginUser == nil {
		return
	}
	connections, err := resolvers.UserUseCase.GetTokensByUsername(username, usertoken)

	if err == nil {
		c.JSON(200, gin.H{
			"connections": connections,
		})
	} else {
		c.JSON(200, gin.H{
			"error": "can not find tokens",
		})
	}
}

func UpdateTokenNameByID(c *gin.Context) {
	usertoken := models.UserGetToken{}
	username := c.Param("username")
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "updatetoken")
	if loginUser == nil {
		return
	}
	err := resolvers.UserUseCase.UpdateUserTokenById(username, usertoken)
	if err == nil {
		c.JSON(200, gin.H{
			"message": "update token successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"error": "can not update tokens",
		})
	}
}

func CheckToken(c *gin.Context) {
	usertoken := models.UserGetToken{}
	username := c.Param("username")
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "checktoken")
	if loginUser == nil {
		return
	}
	err := resolvers.UserUseCase.CheckTokenByID(username, usertoken)
	if err == nil {
		c.JSON(200, gin.H{
			"message": "correct token",
		})
	} else {
		c.JSON(200, gin.H{
			"error": "there is no token with this id",
		})
	}
}
func DeleteConnectionByID(c *gin.Context) {
	usertoken := models.UserGetToken{}
	username := c.Param("username")
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "deletetoken")
	if loginUser == nil {
		return
	}
	err := resolvers.UserUseCase.DeleteUserTokenById(username, usertoken)
	if err == nil {
		c.JSON(200, gin.H{
			"message": "deleted token successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"error": "can not delete token with this ID",
		})
	}
}

func UpdateUserTokenByName(c *gin.Context) {
	usertoken := models.UserGetToken{}
	username := c.Param("username")
	_ = c.BindJSON(&usertoken)

	// Authenticate
	loginUser := checkAuthforconnection(c, []string{username}, "updateotoken")
	if loginUser == nil {
		return
	}
	err := resolvers.UserUseCase.UpdateUserTokenByName(username, usertoken)
	if err == nil {
		c.JSON(200, gin.H{
			"message": "update token successfully",
		})
	} else {
		c.JSON(200, gin.H{
			"error": "can not update tokens",
		})
	}
}
