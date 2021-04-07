package router

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"legato_server/db"
	"legato_server/models"
	"log"
	"net/http"
	"regexp"
)

const Webhook = "Webhook"

var webhookRG = routeGroup{
	name: "Webhook",
	routes: routes{
		route{
			"Create Webhook",
			POST,
			"services/webhook",
			handleNewWebhook,
		},
		route{
			"Webhook",
			POST,
			"services/webhook/:webhookid",
			handleWebhookData,
		},
		route{
			"Update Webhook",
			PATCH,
			"services/webhook/:webhookid",
			handleUpdateWebhook,
		},
	},
}

func handleWebhookData(c *gin.Context) {
	param := c.Param("webhookid")

	wh, err := webhookExists(param)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": err.Error()},
		)
		return
	}
	if !wh.IsEnable {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "webhook is not enabled",
		})
		return
	}
	webhookData := make(map[string]interface{})
	err = json.NewDecoder(c.Request.Body).Decode(&webhookData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
		return
	}
	log.Println("webhook with id ", param)
	log.Println("got payload: ")
	for k, v := range webhookData {
		log.Printf("%s : %v\n", k, v)
	}
	wh.Next(webhookData)
}

func handleNewWebhook(c *gin.Context) {
	req := models.NewWebhook{}
	_ = c.BindJSON(req)
	url := resolvers.WebhookUseCase.Create(req.Name)
	c.JSON(http.StatusOK, url)
}

func IsValidUUID(uuid string) bool {
    r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
    return r.MatchString(uuid)
}

func handleUpdateWebhook(c *gin.Context) {
	param := c.Param("webhookid")
	_, err := webhookExists(param)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			gin.H{"message": err},
		)
		return
	}
	dataMap := make(map[string]interface{})
	err = json.NewDecoder(c.Request.Body).Decode(&dataMap)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}
	err = resolvers.WebhookUseCase.Update(param, dataMap)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "updated successfully",
	})
}

func webhookExists(WebhookID string) (*legatoDb.Webhook, error) {
	if !IsValidUUID(WebhookID) {
		return &legatoDb.Webhook{}, errors.New("bad request")
	}
	wh, err := resolvers.WebhookUseCase.Exists(WebhookID)
	if err != nil {
		return &legatoDb.Webhook{}, err
	}
	return wh, nil
}
