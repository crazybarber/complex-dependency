package repository

import (
	"docugraphy/config"
	"docugraphy/repository/postgres"
	"fmt"
)

func Create() error {
	var err error = nil

	switch dbType := config.GetConfig().DbType; dbType {
	case TypeNamePostgres:
		dbRepo := postgres.Create()
		err = dbRepo.SetupSchema()
	default:
		err = fmt.Errorf("Unknown database id %s\n", dbType)
	}
	return err
}
