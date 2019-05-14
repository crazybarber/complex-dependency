package postgres

import (
	"docugraphy/config"
	"docugraphy/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func SetupSchema() error {
	db := pg.Connect(&pg.Options{
		User:     config.GetConfig().DbUser,
		Password: config.GetConfig().DbPassword,
		Database: config.GetConfig().DbName,
	})
	defer db.Close()

	setSchema(config.GetConfig().DbSchema)

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
			return err
		}
	}
	return nil
}
