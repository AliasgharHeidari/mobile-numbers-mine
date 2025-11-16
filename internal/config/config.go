package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

var AppConfig *Conf

type Conf struct {
	API   API         `yaml:"api"`
	Redis RedisConfig `yaml:"redis"`
}

func InitConfig() {

	AppConfig = &Conf{}

	configData, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configData, AppConfig)
	if err != nil {
		panic(err)
	}
}
