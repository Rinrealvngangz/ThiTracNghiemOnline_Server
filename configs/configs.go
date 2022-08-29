package config

import (
	util "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/utils"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

// Configuration contains static info required to run the apps
// It contains DB info
type Configuration struct {
	Port                  string `env:"GO_PORT"`
	DatabaseConnectionURL string `env:"CONNECTION_URL,required"`
}

// NewConfig will read the config data from given .env file
func NewConfig(files ...string) *Configuration {
	logger := util.Gologger()
	err := godotenv.Load(files...) // Loading config from env file

	if err != nil {
		logger.Error("No .env file could be found %q\n", files)
	}

	cfg := Configuration{}
	// Parse env to configuration
	err = env.Parse(&cfg)
	if err != nil {
		logger.Error("%+v\n", err)
	}
	return &cfg
}
