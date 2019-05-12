package postgres

import (
	"docugraphy/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func (pr *PostgresRepository) SetupSchema() error {
	db := pg.Connect(&pg.Options{
		User:     pr.DbUser,
		Password: pr.DbPassword,
		Database: pr.DbName,
	})
	defer db.Close()

	setSchema(pr.DbSchema)

	for _, entity := range []interface{}{
		(*model.SourceSystem)(nil),
		(*model.EventType)(nil),
		(*model.EventTypeImplementation)(nil),
		(*model.Field)(nil),
		(*model.FieldImplementation)(nil),
		(*model.RestrictedValue)(nil),
	} {
		err := db.CreateTable(entity, &orm.CreateTableOptions{
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

func (pr *PostgresRepository) GetSourceSystems() ([]string, error) {
	db := pg.Connect(&pg.Options{
		User:     pr.DbUser,
		Password: pr.DbPassword,
		Database: pr.DbName,
	})
	defer db.Close()

	var sourceSystem = new([]model.SourceSystem)

	err := db.Model(&sourceSystem).Select()
	if err != nil {
		return nil, err
	}

	return nil, nil
}
