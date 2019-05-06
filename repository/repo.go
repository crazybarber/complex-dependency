package repository

import (
	"docugraphy/config"
	"docugraphy/repository/postgres"
	"fmt"
)

func Create() error {
	var err error = nil

	switch dbConfig := config.GetConfig().Db; dbConfig {
	case DbModuleConfigName:
		err = postgres.SetupSchema()
	default:
		err = fmt.Errorf("Unknown database id %s\n", dbConfig)
	}
	return err
}
