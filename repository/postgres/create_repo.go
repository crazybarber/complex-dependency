package postgres

import (
	"docugraphy/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
)

func (pr *Repository) Connect() {
	pr.Connection = pg.Connect(&pg.Options{
		User:     pr.DbUser,
		Password: pr.DbPassword,
		Database: pr.DbName,
	})
	setSchema(pr.DbSchema)
	log.Println("Postgres DB connected")
}

func (pr *Repository) Disconnect() error {
	return pr.Connection.Close()
}

func (pr *Repository) SetupSchema() error {
	setSchema(pr.DbSchema)

	for _, entity := range []interface{}{
		(*model.SourceSystem)(nil),
		(*model.EventType)(nil),
		(*model.EventTypeImplementation)(nil),
		(*model.Field)(nil),
		(*model.FieldImplementation)(nil),
		(*model.RestrictedValue)(nil),
	} {
		err := pr.Connection.CreateTable(entity, &orm.CreateTableOptions{
			FKConstraints: true,
			Varchar:       150,
			Temp:          false,
			IfNotExists:   true,
		})
		if err != nil {
		}
	}
	return nil
}

func (pr *Repository) GetSourceSystems() ([]model.SourceSystem, error) {
	var sourceSystem []model.SourceSystem

	err := pr.Connection.Model(&sourceSystem).Select()
	if err != nil {
		return nil, err
	}
	return sourceSystem, nil
}

func (pr *Repository) AddSourceSystem(sourceSystem *model.SourceSystem) error {
	err := pr.Connection.Insert(sourceSystem)
	if err != nil {
		return err
	}
	log.Println("Inserted new Source System: ", sourceSystem.Name)
	return nil
}

func (pr *Repository) GetEventTypes() ([]model.EventType, error) {
	panic("implement me")
}

func (pr *Repository) GetFields() ([]model.Field, error) {
	panic("implement me")
}

func (pr *Repository) GetEventTypeImplementation(sourceSystemName string) ([]model.EventTypeImplementation, error) {
	panic("implement me")
}

func (pr *Repository) GetFieldImplementation(sourceSystemName string, eventTypeName string) ([]model.FieldImplementation, error) {
	panic("implement me")
}

func (pr *Repository) AddEventType(eventType *model.EventType) error {
	panic("implement me")
}

func (pr *Repository) AddField(field *model.Field) error {
	panic("implement me")
}

func (pr *Repository) AddEventTypeImplementation(eventTypeImplementation *model.EventTypeImplementation) error {
	panic("implement me")
}

func (pr *Repository) AddFieldImplementation(fieldImplementation *model.FieldImplementation) error {
	panic("implement me")
}

func (pr *Repository) ChangeStatusOfEventType(eventTypeName string, status model.ImplementationStatus) error {
	panic("implement me")
}

func (pr *Repository) ChangeStatusOfField(FieldName string, status model.ImplementationStatus) error {
	panic("implement me")
}

func (pr *Repository) ChangeStatusOfEventTypeImplementation(EventTypeName string, sourceSystem string, status model.ImplementationStatus) error {
	panic("implement me")
}

func (pr *Repository) ChangeStatusOfFieldImplementation(FieldName string, sourceSystem string, status model.ImplementationStatus) error {
	panic("implement me")
}
