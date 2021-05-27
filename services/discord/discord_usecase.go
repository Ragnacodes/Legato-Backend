package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	"legato_server/helper/converter"
	"log"
	"net/http"
	"time"
)

const discordUrl = "https://discord.com/api/"

type DiscordUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewDiscordUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.DiscordUseCase {
	return &DiscordUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (du DiscordUseCase) AddToScenario(u *api.UserInfo, scenarioId uint, nh api.NewServiceNode) (api.ServiceNode, error) {
	user, err := du.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := du.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	var discord legatoDb.Discord
	discord.Service = converter.NewServiceNodeToServiceDb(nh)

	h, err := du.db.CreateDiscord(&scenario, discord)
	if err != nil {
		return api.ServiceNode{}, err
	}

	return converter.ServiceDbToServiceNode(h.Service), nil
}

func (du DiscordUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, nd api.NewServiceNode) error {
	user, err := du.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	scenario, err := du.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	var discord legatoDb.Discord
	discord.Service = converter.NewServiceNodeToServiceDb(nd)

	err = du.db.UpdateDiscord(&scenario, serviceId, discord)
	if err != nil {
		return err
	}

	return nil
}

func (du DiscordUseCase) GetGuildTextChannels(guildId string) (channels api.Channels, err error) {
	token := env.DiscordBotToken
	getChannelsUrl := fmt.Sprintf("%sguilds/%s/channels", discordUrl, guildId)
	log.Println(getChannelsUrl)
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, getChannelsUrl, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", token)
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// read the payload, in this case, Jhon's info
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	// this is where the magic happens, I pass a pointer of type Person and Go will do the rest
	allChannels := &api.Channels{}
	err = json.Unmarshal(body, &allChannels)
	if err != nil {
		return
	}

	for _, ch := range *allChannels {
		if ch.Type == api.DiscordChannelTypeGuildText {
			channels = append(channels, ch)
		}
	}

	return
}
