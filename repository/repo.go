package repository

import (
	"docugraphy/config"
	"docugraphy/model"
	"docugraphy/repository/postgres"
	"fmt"
)

var dbRepo Repository

func Create() error {
	err := dbRepo.SetupSchema()
	return err
}

func Connect() error {
	var err error = nil
	switch dbType := config.GetConfig().DbType; dbType {
	case TypeNamePostgres:
		dbRepo = postgres.Build()
	default:
		err = fmt.Errorf("Unknown database id %s\n", dbType)
		return err
	}
	dbRepo.Connect()
	return err
}

func GetSourceSystems() ([]model.SourceSystem, error) {
	return dbRepo.GetSourceSystems()
}

func AddSourceSystem(sourceSystem *model.SourceSystem) error {
	return dbRepo.AddSourceSystem(sourceSystem)
}
