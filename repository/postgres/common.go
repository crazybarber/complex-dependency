package postgres

import "github.com/go-pg/pg/orm"

const (
	DbModuleConfigName string = "postgres"
)

func setSchema(schemaName string) {
	orm.SetTableNameInflector(func(name string) string {
		return schemaName + "." + name
	})
}
