package config

import "github.com/tkanos/gonfig"

type Configuration struct {
	DB_NAME     string
	DB_USERNAME string
	DB_PASSWORD string
	DB_HOST     string
	DB_PORT     string
}

func GetConfig() Configuration {
	config := Configuration{}
	gonfig.GetConf("config/config.json", &config)
	return config
}