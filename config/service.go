package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Service Service
	Cors    Cors
}

type Service struct {
	BaseUrl string `yaml:"baseUrl"`
	LogPath string `yaml:"logPath"`
}

type Cors struct {
	Origin string `yaml:"origin"`
}

func LoadConfig() *Config {
	env := os.Getenv("ENV")
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(err)
	}

	return config
}
