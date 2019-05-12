package postgres

import (
	"docugraphy/config"
	"github.com/go-pg/pg/orm"
)

var dbConnection *orm.DB

type PostgresRepository struct {
	DbSchema   string
	DbName     string
	DbUser     string
	DbPassword string
}

func Create() PostgresRepository {
	dbConfig := config.GetConfig().DbConfig
	return PostgresRepository {
		DbSchema: dbConfig.DbSchema,
		DbName: dbConfig.DbName,
		DbUser: dbConfig.DbUser,
		DbPassword: dbConfig.DbPassword,
	}

}

func setSchema(schemaName string) {
	orm.SetTableNameInflector(func(name string) string {
		return schemaName + "." + name
	})
}


