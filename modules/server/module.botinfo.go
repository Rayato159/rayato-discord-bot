package server

import (
	"github.com/Rayato159/rayato-discord-bot/modules/botinfo/botinfoHandlers"
	"github.com/Rayato159/rayato-discord-bot/modules/botinfo/botinfoUsecases"
	"github.com/bwmarrin/discordgo"
)

type IBotinfoModule interface {
	Init()
	Handler() botinfoHandlers.IBotinfoHandler
	Usecase() botinfoUsecases.IBotinfoUsecase
}

type botinfoModule struct {
	*module
	handler botinfoHandlers.IBotinfoHandler
	usecase botinfoUsecases.IBotinfoUsecase
}

func (m *module) BotinfoModule() IBotinfoModule {
	botinfoUsecase := botinfoUsecases.NewBotinfoUsecase()
	botinfoHandler := botinfoHandlers.NewBotinfoHandler(botinfoUsecase)

	return &botinfoModule{
		module:  m,
		handler: botinfoHandler,
		usecase: botinfoUsecase,
	}
}

func (b *botinfoModule) Init() {
	b.module.commands = append(b.module.commands, &discordgo.ApplicationCommand{
		Name:        "help",
		Description: "Just a help menu",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "feature",
				Description: "Just an example",
				Type:        discordgo.ApplicationCommandOptionString,
			},
		},
	})

	b.commandHandlers["help"] = b.handler.Help
}
func (b *botinfoModule) Handler() botinfoHandlers.IBotinfoHandler { return b.handler }
func (b *botinfoModule) Usecase() botinfoUsecases.IBotinfoUsecase { return b.usecase }
