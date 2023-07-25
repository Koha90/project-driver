package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/koha90/project-driver/pkg/logging"
)

type Config struct {
	IsDebug       *bool  `json:"is_debug"`
	Addr          string `json:"addr"`
	Port          string `json:"port"`
	StorageConfig struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Database string `json:"database"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"storage_config"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Print("gather config")

		instance = &Config{}
		err := cleanenv.ReadConfig("config.json", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
