package server

import "github.com/bwmarrin/discordgo"

type IModule interface {
	GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	BotinfoModule() IBotinfoModule
}

type module struct {
	*discordServer
	commandHandlers map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
}

func ModuleInit(s *discordServer) IModule {
	return &module{
		discordServer:   s,
		commandHandlers: make(map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)),
	}
}

func (m *module) GetCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	return m.commandHandlers
}
