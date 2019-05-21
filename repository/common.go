package repository

import "docugraphy/model"

const (
	TypeNamePostgres string = "postgres"
)

type Repository interface {
	Connect()
	Disconnect() error
	SetupSchema() error
	GetSourceSystems() ([]model.SourceSystem, error)
	AddSourceSystem(sourceSystem *model.SourceSystem) error
}
