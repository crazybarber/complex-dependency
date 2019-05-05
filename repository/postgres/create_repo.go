package postgres

import (
	"complex-dependency/config"
	"complex-dependency/repository"
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
		(*repository.SourceSystem)(nil),
		(*repository.EventType)(nil),
		(*repository.EventTypeImplementation)(nil),
		(*repository.Field)(nil),
		(*repository.FieldImplementation)(nil),
		(*repository.RestrictedValue)(nil),
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
