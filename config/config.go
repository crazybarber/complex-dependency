package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	DbName     string `json:"db_name"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_pass"`
}

var config = Configuration{}

func Load(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)

	if err != nil {
		return err
	}
	return nil
}

func GetConfig() Configuration {
	return config
}
