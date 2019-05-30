package repository

import (
	"docugraphy/config"
	"docugraphy/model"
	"docugraphy/repository/postgres"
	"fmt"
)

var DbRepo Repository

func Create() error {
	err := DbRepo.SetupSchema()
	return err
}

func Connect() error {
	var err error = nil
	switch dbType := config.GetConfig().DbType; dbType {
	case TypeNamePostgres:
		DbRepo = postgres.Build()
	default:
		err = fmt.Errorf("Unknown database id %s\n", dbType)
		return err
	}
	DbRepo.Connect()
	return err
}

func GetSourceSystems() ([]model.SourceSystem, error) {
	return DbRepo.GetSourceSystems()
}

func AddSourceSystem(sourceSystem *model.SourceSystem) error {
	return DbRepo.AddSourceSystem(sourceSystem)
}
