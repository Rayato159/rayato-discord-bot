package main

import (
	"github.com/Rayato159/rayato-discord-bot/configs"
	"github.com/Rayato159/rayato-discord-bot/modules/server"
)

func main() {
	cfg := configs.NewConfig("/bin/.env")

	server.NewDiscordServer(cfg).Start()
}
