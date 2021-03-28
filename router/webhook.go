package router

import (
	"encoding/json"
	"fmt"
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
			"create/webhook",
			handleNewWebhook,
		},
		route{
			"Webhook",
			POST,
			"services/webhook/:webhookid/",
			handleWebhookData,
		},
	},
}

func handleWebhookData(c *gin.Context) {
	  
	param := c.Param("webhookid")
	
	if !IsValidUUID(param) {
		c.JSON(400,gin.H{"message":"bad request"})
		return
	}
	exists, err := resolvers.WebhookUseCase.WebhookExistOr404(param)
	if !exists||err!=nil{
		c.JSON(http.StatusBadRequest,
		gin.H{"message": "no webhook with this id",},
		)
		return
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
}

func handleNewWebhook(c *gin.Context){
	name := c.Param("Name")
	url, err := resolvers.WebhookUseCase.CreateNewWebhook(name)
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
