package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"legato_server/models"
	"legato_server/db"
	"legato_server/services"
	"net/http"
	"regexp"
	"github.com/gin-gonic/gin"
)


const Webhook = "Webhook"

var webhookRG = routeGroup{
	name: "Webhook",
	routes: routes{
		route{
			"Create Webhook",
			POST,
			"services/webhook/",
			handleNewWebhook,
		},
		route{
			"Webhook",
			POST,
			"services/webhook/:webhookid/",
			handleWebhookData,
		},
		route{
			"Create Webhook",
			PATCH,
			"services/webhook/:webhookid/",
			handleUpdateWebhook,
		},
	},
}

func handleWebhookData(c *gin.Context) {
	  
	param := c.Param("webhookid")

	wh, err := webhookExists(param)
	if err!= nil{
		c.JSON(http.StatusBadRequest,
			gin.H{"message": err,},
			)
			return
	}
	if !wh.Enable {
		c.AbortWithStatusJSON(http.StatusForbidden,gin.H{
			"message": "webhook is not enabled",
		})
	}
	webhookData := make(map[string]interface{})
	err = json.NewDecoder(c.Request.Body).Decode(&webhookData)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": err.Error(),
		})
	}
	fmt.Println("webhook with id ", param)
	fmt.Println("got payload: ")
	for k, v := range webhookData {
		fmt.Printf("%s : %v\n", k, v)
	}

	services.NewWebhook(wh.Service.Name, wh.Service.Children).Next(webhookData)
}

func handleNewWebhook(c *gin.Context){
	req := models.NewWebhook{}
	c.BindJSON(req)
	url, err := resolvers.WebhookUseCase.Create(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, url)
}

func IsValidUUID(uuid string) bool {
    r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
    return r.MatchString(uuid)
}


func handleUpdateWebhook(c *gin.Context){
	param := c.Param("webhookid")
	_, err := webhookExists(param)
	if err!= nil{
		c.JSON(http.StatusBadRequest,
			gin.H{"message": err,},
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
	resolvers.WebhookUseCase.Update(param, dataMap)
	c.JSON(http.StatusOK, gin.H{
		"message": "updated successfully",
	})
}

func webhookExists(WebhookID string) (legatoDb.Webhook, error){
	
	if !IsValidUUID(WebhookID) {
		return legatoDb.Webhook{},errors.New("bad request")
	}
	wh, err := resolvers.WebhookUseCase.Exists(WebhookID)
	if err!=nil{
		return legatoDb.Webhook{},errors.New("no webhook with this id")
	}
	return wh, nil 
} 