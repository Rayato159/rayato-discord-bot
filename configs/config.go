package configs

import (
	"log"

	"github.com/joho/godotenv"
)

type IConfig interface {
	App() IAppConfig
}

type config struct {
	app *app
}

func NewConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatal("read .env file failed")
	}

	return &config{
		app: &app{
			token: envMap["APP_TOKEN"],
		},
	}
}

type IAppConfig interface {
	GetToken() string
}

type app struct {
	token string
}

func (c *config) App() IAppConfig {
	return c.app
}

func (a *app) GetToken() string {
	return a.token
}
