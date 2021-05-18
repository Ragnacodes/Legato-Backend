package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/ssh"
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
		route{
			"check password",
			POST,
			"/check/ssh/pass",
			checkPassSSH,
		},
		route{
			"check sshkey",
			POST,
			"/check/ssh/key",
			checkKeySSH,
		},
	},
}

type SSHLoginWithPass struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type SSHLoginWithKey struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Sshkey   string `json:"sshKey"`
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

func checkPassSSH(c *gin.Context) {
	data := SSHLoginWithPass{}
	_ = c.BindJSON(&data)
	config := &ssh.ClientConfig{
		User: data.Username,
		Auth: []ssh.AuthMethod{
			ssh.Password(data.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	_, err := ssh.Dial("tcp", data.Host+":22", config)

	if err != nil {
		c.JSON(500, err.Error())
	}
}

func checkKeySSH(c *gin.Context) {
	data := SSHLoginWithKey{}
	_ = c.BindJSON(&data)
	signer, err := ssh.ParsePrivateKey([]byte(data.Sshkey))

	if err != nil {
		c.JSON(500, err)
	}

	config := &ssh.ClientConfig{
		User: data.Username,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	_, err = ssh.Dial("tcp", data.Host+":22", config)

	if err != nil {
		c.JSON(500, err)
	}
}
