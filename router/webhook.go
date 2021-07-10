package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"legato_server/api"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

var webhookRG = routeGroup{
	name: "Webhook",
	routes: routes{
		route{
			"Add a new separate webhook",
			POST,
			"/users/:username/services/webhooks",
			createNewWebhook,
		},
		route{
			"Webhook Trigger",
			POST,
			"/services/webhook/:webhookid",
			handleWebhookData,
		},
		route{
			"Update Webhook",
			PUT,
			"/users/:username/services/webhooks/:webhook_id",
			updateUserWebhooks,
		},
		route{
			"Get user webhooks",
			GET,
			"/users/:username/services/webhooks",
			getUserWebhooks,
		},
		route{
			"Get user webhook by id",
			GET,
			"/users/:username/services/webhooks/:webhook_id",
			getUserWebhookById,
		},
		route{
			"Delete a webhook",
			DELETE,
			"/users/:username/services/webhooks/:webhook_id",
			deleteUserWebhook,
		},
		route{
			"Get Webhook history",
			GET,
			"/users/:username/services/webhooks/:webhook_id/histories",
			getWebhookHistories,
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
	resolvers.WebhookUseCase.TriggerWebhook(param, webhookData)

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
	})
}

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

func updateUserWebhooks(c *gin.Context) {
	username := c.Param("username")
	webhookId, _ := strconv.Atoi(c.Param("webhook_id"))

	nwh := api.NewSeparateWebhook{}
	_ = c.BindJSON(&nwh)

	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	err := resolvers.WebhookUseCase.UpdateSeparateWebhook(loginUser, uint(webhookId), nwh)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not update this separate webhook: %s", err),
		})
		return
	}

	updatedWebhook, err := resolvers.WebhookUseCase.GetUserWebhookById(loginUser, uint(webhookId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not update this separate webhook: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "webhook is updated successfully",
		"webhook": updatedWebhook,
	})
}

func getUserWebhookById(c *gin.Context) {
	username := c.Param("username")
	webhookId, _ := strconv.Atoi(c.Param("webhook_id"))

	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	wh, err := resolvers.WebhookUseCase.GetUserWebhookById(loginUser, uint(webhookId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not update this separate webhook: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"webhook": wh,
	})
}
func deleteUserWebhook(c *gin.Context) {
	username := c.Param("username")
	webhookId, _ := strconv.Atoi(c.Param("webhook_id"))

	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	err := resolvers.WebhookUseCase.DeleteUserWebhookById(loginUser, uint(webhookId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not delete this separate webhook: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "webhook is deleted successfully",
	})
}

func webhookExists(WebhookID string) (api.WebhookInfo, error) {
	if !IsValidUUID(WebhookID) {
		return api.WebhookInfo{}, errors.New("bad request")
	}
	wh, err := resolvers.WebhookUseCase.Exists(WebhookID)
	if err != nil {
		return api.WebhookInfo{}, err
	}
	return *wh, nil
}

func getWebhookHistories(c *gin.Context) {
	username := c.Param("username")
	webhookId, _ := strconv.Atoi(c.Param("webhook_id"))

	// Authenticate
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	logsList, err := resolvers.WebhookUseCase.GetUserWebhookHistoryById(loginUser, uint(webhookId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("can not find webhook data: %s", err),
		})
		return
	}

	if logsList == nil {
		response := []int{}
		c.JSON(http.StatusOK, gin.H{
			"logs": response,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{

		"logs": logsList,
	})
}
