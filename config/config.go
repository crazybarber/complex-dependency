package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	ListenPort string          `json:"listen_port"`
	DbType     string          `json:"type"`
	DbConfig   DbConfiguration `json:"db_config"`
}

type DbConfiguration struct {
	DbSchema   string `json:"schema"`
	DbName     string `json:"name"`
	DbUser     string `json:"user"`
	DbPassword string `json:"pass"`
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
