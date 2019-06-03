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
		(*model.ImplementationStatusDictionary)(nil),
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

func (pr *Repository) GetEventTypes() ([]model.EventType, error) {
	var eventType []model.EventType

	err := pr.Connection.Model(&eventType).Select()
	if err != nil {
		return nil, err
	}
	return eventType, nil
}

func (pr *Repository) getEventType(eventTypeName string) (model.EventType, error) {
	eventType := model.EventType{}

	err := pr.Connection.Model(&eventType).
		Where("event_type.name = ?", eventTypeName).
		Select()
	if err != nil {
		return eventType, err
	}
	return eventType, nil
}

func (pr *Repository) GetFields() ([]model.Field, error) {
	var fields []model.Field

	err := pr.Connection.Model(&fields).Select()
	if err != nil {
		return nil, err
	}
	return fields, nil
}

func (pr *Repository) getField(fieldName string) (model.Field, error) {
	field := model.Field{}

	err := pr.Connection.Model(&field).
		Where("field.name = ?", fieldName).
		Select()
	if err != nil {
		return field, err
	}
	return field, nil
}

func (pr *Repository) GetEventTypeImplementations(sourceSystemName string) ([]model.EventTypeImplementation, error) {
	var eventTypeImplementations []model.EventTypeImplementation

	err := pr.Connection.
		Model(&eventTypeImplementations).
		Where("event_type_implementation.source_system = ?", sourceSystemName).
		Select()
	if err != nil {
		return nil, err
	}
	return eventTypeImplementations, nil
}

func (pr *Repository) getEventTypeImplementation(sourceSystemName string, eventTypeName string) (model.EventTypeImplementation, error) {
	eventTypeImplementation := model.EventTypeImplementation{}

	err := pr.Connection.
		Model(&eventTypeImplementation).
		Where("event_type_implementation.source = ?", sourceSystemName).
		Where("event_type_implementation.event_type_name = ?", eventTypeName).
		Select()
	if err != nil {
		return eventTypeImplementation, err
	}
	return eventTypeImplementation, nil
}

func (pr *Repository) GetFieldImplementations(sourceSystemName string, eventTypeName string) ([]model.FieldImplementation, error) {
	var fieldImplementations []model.FieldImplementation

	err := pr.Connection.
		Model(&fieldImplementations).
		Column("event_type_implementation").
		Where("event_type_implementation.source = ?", sourceSystemName).
		Where("event_type_implementation.event_type_name = ?", eventTypeName).
		Select("field.*")
	if err != nil {
		return nil, err
	}
	return fieldImplementations, nil
}

func (pr *Repository) getFieldImplementation(sourceSystemName string, eventTypeName string, fieldName string) (model.FieldImplementation, error) {
	fieldImplementation := model.FieldImplementation{}

	err := pr.Connection.
		Model(&fieldImplementation).
		Where("field_implementation.source = ?", sourceSystemName).
		Where("field_implementation.event_type_name = ?", eventTypeName).
		Where("field_implementation.field_name = ?", fieldName).
		Select()
	if err != nil {
		return fieldImplementation, err
	}
	return fieldImplementation, nil
}

func (pr *Repository) AddSourceSystem(sourceSystem *model.SourceSystem) error {
	err := pr.Connection.Insert(sourceSystem)
	if err != nil {
		return err
	}
	log.Println("Inserted new Source System: ", sourceSystem.Name)
	return nil
}

func (pr *Repository) AddEventType(eventType *model.EventType) error {
	err := pr.Connection.Insert(eventType)
	if err != nil {
		return err
	}
	log.Println("Inserted new Event Type: ", eventType.Name)
	return nil
}

func (pr *Repository) AddField(field *model.Field) error {
	err := pr.Connection.Insert(field)
	if err != nil {
		return err
	}
	log.Println("Inserted new Field: ", field.Name)
	return nil
}

func (pr *Repository) AddEventTypeImplementation(eventTypeImplementation *model.EventTypeImplementation) error {
	err := pr.Connection.Insert(eventTypeImplementation)
	if err != nil {
		return err
	}
	log.Println("Inserted new Event Type Implementation with Event Type: ",
		eventTypeImplementation.EventTypeName,
		", and source system: ",
		eventTypeImplementation.SourceSystemName)
	return nil
}

func (pr *Repository) AddFieldImplementation(fieldImplementation *model.FieldImplementation) error {
	err := pr.Connection.Insert(fieldImplementation)
	if err != nil {
		return err
	}
	log.Println("Inserted new Event Type Implementation with Event Type: ",
		fieldImplementation.EventTypeName,
		", and field: ",
		fieldImplementation.FieldName,
		", and source system: ",
		fieldImplementation.SourceSystemName)
	return nil
}

func (pr *Repository) ChangeStatusOfEventType(eventTypeName string, status model.ImplementationStatus) error {
	eventType, err := pr.getEventType(eventTypeName)
	if err != nil {
		return err
	}
	eventType.Status = status
	err = pr.Connection.Update(&eventType)
	if err != nil {
		return err
	}
	return nil
}

func (pr *Repository) ChangeStatusOfField(fieldName string, status model.ImplementationStatus) error {
	field, err := pr.getField(fieldName)
	if err != nil {
		return err
	}
	field.Status = status
	err = pr.Connection.Update(&field)
	if err != nil {
		return err
	}
	return nil
}

func (pr *Repository) ChangeStatusOfEventTypeImplementation(eventTypeName string, sourceSystem string, status model.ImplementationStatus) error {
	eventTypeImplementation, err := pr.getEventTypeImplementation(eventTypeName, sourceSystem)
	if err != nil {
		return err
	}
	eventTypeImplementation.Status = status
	err = pr.Connection.Update(&eventTypeImplementation)
	if err != nil {
		return err
	}
	return nil
}

func (pr *Repository) ChangeStatusOfFieldImplementation(fieldName string, eventTypeName string, sourceSystem string, status model.ImplementationStatus) error {
	fieldImplementation, err := pr.getFieldImplementation(fieldName, eventTypeName, sourceSystem)
	if err != nil {
		return err
	}
	fieldImplementation.Status = status
	err = pr.Connection.Update(&fieldImplementation)
	if err != nil {
		return err
	}
	return nil
}
