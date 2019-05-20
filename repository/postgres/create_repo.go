package postgres

import (
	"docugraphy/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func (pr *Repository) Connect() {
	pr.Connection = pg.Connect(&pg.Options{
		User:     pr.DbUser,
		Password: pr.DbPassword,
		Database: pr.DbName,
	})
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

func (pr *Repository) AddSourceSystem(sourceSystem model.SourceSystem) error {
	err := pr.Connection.Insert(sourceSystem)
	if err != nil {
		return err
	}
	return nil
}
