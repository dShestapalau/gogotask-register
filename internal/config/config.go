package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Configuration struct {
	ServerPort       string `mapstructure:"SERVER_PORT"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
}

func NewConfiguration() Configuration {
	return Configuration{}
}

func (config *Configuration) LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Error reading config file %d", err)
	}

	viper.Unmarshal(&config)
}
