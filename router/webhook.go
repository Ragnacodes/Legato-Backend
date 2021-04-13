package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"legato_server/api"
	"net/http"
	"regexp"
)

const Webhook = "Webhook"

var webhookRG = routeGroup{
	name: "Webhook",
	routes: routes{
		route{
			"Add a new separate webhook",
			POST,
			"/users/:username/services/webhooks",
			createNewWebhook,
		},
		//route{
		//	"Webhook Trigger",
		//	POST,
		//	"/services/webhook/:webhookid",
		//	handleWebhookData,
		//},
		//route{
		//	"Update Webhook",
		//	PATCH,
		//	"/users/:username/services/webhook/:webhookid",
		//	handleUpdateWebhook,
		//},
		route{
			"Get user webhooks",
			GET,
			"/users/:username/services/webhooks",
			getUserWebhooks,
		},
	},
}

//func handleWebhookData(c *gin.Context) {
//	param := c.Param("webhookid")
//
//	wh, err := webhookExists(param)
//	if err != nil {
//		c.JSON(http.StatusBadRequest,
//			gin.H{"message": err.Error()},
//		)
//		return
//	}
//	if !wh.IsEnable {
//		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
//			"message": "webhook is not enabled",
//		})
//		return
//	}
//	webhookData := make(map[string]interface{})
//	err = json.NewDecoder(c.Request.Body).Decode(&webhookData)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{
//			"message": err.Error(),
//		})
//		return
//	}
//	log.Println("webhook with id ", param)
//	log.Println("got payload: ")
//	for k, v := range webhookData {
//		log.Printf("%s : %v\n", k, v)
//	}
//	wh.Next(webhookData)
//}

func createNewWebhook(c *gin.Context) {
	username := c.Param("username")

	nwh := api.NewSeparateWebhook{}
	_ = c.BindJSON(&nwh)

	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	webhookInfo, err := resolvers.WebhookUseCase.CreateSeparateWebhook(loginUser, nwh)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not add separate webhook: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "webhook is added successfully",
		"webhook": webhookInfo,
	})
}

func IsValidUUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

func getUserWebhooks(c *gin.Context) {
	username := c.Param("username")

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	// Get Webhooks
	userWebhooks, err := resolvers.WebhookUseCase.GetUserWebhooks(loginUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not fetch user webhooks: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"webhooks": userWebhooks,
	})
}

//func handleUpdateWebhook(c *gin.Context){
//	username := c.Param("username")
//	loginUser := checkAuth(c, []string{username})
//	if loginUser == nil {
//		return
//	}
//	param := c.Param("webhookid")
//	_, err := webhookExists(param)
//	if err != nil {
//		c.JSON(http.StatusBadRequest,
//			gin.H{"message": err},
//		)
//		return
//	}
//
//	dataMap := make(map[string]interface{})
//	err = json.NewDecoder(c.Request.Body).Decode(&dataMap)
//	for k, v := range dataMap {
//		log.Printf("%s : %v\n", k, v)
//	}
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{
//			"message": err.Error(),
//		})
//		return
//	}
//
//	err = resolvers.WebhookUseCase.Update(param, dataMap)
//	if err != nil {
//		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
//			"message": err.Error(),
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"message": "updated successfully",
//	})
//}
//
//func webhookExists(WebhookID string) (*legatoDb.Webhook, error) {
//	if !IsValidUUID(WebhookID) {
//		return &legatoDb.Webhook{}, errors.New("bad request")
//	}
//	wh, err := resolvers.WebhookUseCase.Exists(WebhookID)
//	if err != nil {
//		return &legatoDb.Webhook{}, err
//	}
//	return wh, nil
//}
