package server

import (
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/Rayato159/rayato-discord-bot/configs"
	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID        = flag.String("guild", "", "Test guild ID. If not passed - bot registers commands globally")
	RemoveCommands = flag.Bool("rmcmd", true, "Remove all commands after shutdowning or not")
)

type IDiscordServer interface {
	Start()
}

type discordServer struct {
	cfg      configs.IConfig
	dg       *discordgo.Session
	commands []*discordgo.ApplicationCommand
}

func NewDiscordServer(cfg configs.IConfig) IDiscordServer {
	dg, err := discordgo.New("Bot " + cfg.App().GetToken())
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}

	return &discordServer{
		dg:       dg,
		cfg:      cfg,
		commands: make([]*discordgo.ApplicationCommand, 0),
	}
}

func (s *discordServer) Start() {
	s.dg.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	})
	if err := s.dg.Open(); err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	module := ModuleInit(s)
	module.BotinfoModule().Init()

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(s.commands))
	for i, v := range s.commands {
		cmd, err := s.dg.ApplicationCommandCreate(s.dg.State.User.ID, *GuildID, v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}
	defer s.dg.Close()

	s.dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := module.GetCommandHandlers()[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	})

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	if *RemoveCommands {
		log.Println("Removing commands...")
		for _, v := range registeredCommands {
			err := s.dg.ApplicationCommandDelete(s.dg.State.User.ID, *GuildID, v.ID)
			if err != nil {
				log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
			}
		}
	}

	log.Println("Gracefully shutting down.")
}
