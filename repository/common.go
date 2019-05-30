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
	GetEventTypes() ([]model.EventType, error)
	GetFields() ([]model.Field, error)
	GetEventTypeImplementation(sourceSystemName string) ([]model.EventTypeImplementation, error)
	GetFieldImplementation(sourceSystemName string, eventTypeName string) ([]model.FieldImplementation, error)

	AddSourceSystem(sourceSystem *model.SourceSystem) error
	AddEventType(eventType *model.EventType) error
	AddField(field *model.Field) error
	AddEventTypeImplementation(eventTypeImplementation *model.EventTypeImplementation) error
	AddFieldImplementation(fieldImplementation *model.FieldImplementation) error

	ChangeStatusOfEventType(eventTypeName string, status model.ImplementationStatus) error
	ChangeStatusOfField(FieldName string, status model.ImplementationStatus) error
	ChangeStatusOfEventTypeImplementation(eventTypeName string, sourceSystemName string, status model.ImplementationStatus) error
	ChangeStatusOfFieldImplementation(fieldName string, sourceSystemName string, status model.ImplementationStatus) error
}
