package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var discordRG = routeGroup{
	name: "Ssh",
	routes: routes{
		route{
			"Get text channels of that guild",
			GET,
			"/services/discord/guilds/:guildId/channels/text",
			getGuildTextChannels,
		},
		route{
			"Get messages of a single text channels",
			GET,
			"services/discord/channels/:channelId/messages",
			getGuildTextChannelMessages,
		},
	},
}

func getGuildTextChannels(c *gin.Context) {
	guildId := c.Param("guildId")

	// Auth
	// should authenticate with a token for discord connection later
	//loginUser := checkAuth(c, []string{username})
	//if loginUser == nil {
	//	return
	//}

	channels, err := resolvers.DiscordUseCase.GetGuildTextChannels(guildId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("can not fetch text channels: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"channels": channels,
	})
}

func getGuildTextChannelMessages(c *gin.Context) {
	channelId := c.Param("channelId")

	// Auth
	// should authenticate with a token for discord connection later
	//loginUser := checkAuth(c, []string{username})
	//if loginUser == nil {
	//	return
	//}

	channels, err := resolvers.DiscordUseCase.GetGuildTextChannelMessages(channelId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("can not fetch messages: %s", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": channels,
	})
}
