package configs

import (
	"os"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type ConfigFiber struct {
	Type string `yaml:"type"`
}

type ConfigMongoDB struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  ConfigFiber
	MongoDB ConfigMongoDB
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig("./config.yml", instance); err != nil {
			cleanenv.GetDescription(instance, nil)
		}


		if mongoHost := os.Getenv("MONGODB_HOST"); mongoHost != "" {
			instance.MongoDB.Host = mongoHost
		} else {
			instance.MongoDB.Host = "mongodb-instance" 
		}
	})
	return instance
}
