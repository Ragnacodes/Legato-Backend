package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var sshRG = routeGroup{
	name: "Ssh",
	routes: routes{
		// route{
		// 	"Create new http",
		// 	POST,
		// 	"/users/:username/services/ssh",
		// 	createSshService,
		// },
		route{
			"Get user Sshes",
			GET,
			"/users/:username/services/sshes",
			getUserSshs,
		},
		route{
			"Get user Sshes",
			POST,
			"/extract/sshkey/file",
			uploadFile,
		},
	},
}

// func createSshService(c *gin.Context) {
// 	username := c.Param("username")
// 	apissh := api.SshInfo{}
// 	_ = c.BindJSON(&apissh)

// 	// Authenticate
// 	// Auth
// 	loginUser := checkAuth(c, []string{username})
// 	if loginUser == nil {
// 		return
// 	}

// 	ssh, err := resolvers.SshUseCase.AddSsh(username, &apissh)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": fmt.Sprintf("can not add ssh: %s", err),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"Host":     ssh.Host,
// 		"Username": ssh.Username,
// 	})
// }
func getUserSshs(c *gin.Context) {
	username := c.Param("username")

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Get Webhooks
	usersshs, err := resolvers.SshUseCase.GetSshs(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not fetch user ssh: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Sshes": usersshs,
	})
}
func uploadFile(c *gin.Context) {

	file, err := c.FormFile("File")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	byteContainer := make([]byte, 10000000)
	f, _ := file.Open()
	f.Read(byteContainer)

	c.JSON(http.StatusInternalServerError, gin.H{
		"SshKey": string(byteContainer[0:file.Size]),
	})
}
