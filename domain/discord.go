package domain

import "legato_server/api"

type DiscordUseCase interface {
	AddToScenario(u *api.UserInfo, scenarioId uint, nh api.NewServiceNode) (api.ServiceNode, error)
	Update(u *api.UserInfo, scenarioId uint, nodeId uint, nt api.NewServiceNode) error
	GetGuildTextChannels(guildId string) (channels api.Channels, err error)
	GetGuildTextChannelMessages(channelId string) (messages api.Messages, err error)
}
