package postgres

import "github.com/go-pg/pg/orm"

func setSchema(schemaName string) {
	orm.SetTableNameInflector(func(name string) string {
		return schemaName + "." + name
	})
}
