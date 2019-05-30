package postgres

import (
	"docugraphy/config"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type Repository struct {
	DbSchema   string
	DbName     string
	DbUser     string
	DbPassword string
	Connection *pg.DB
}

func Build() *Repository {
	dbConfig := config.GetConfig().DbConfig
	return &Repository{
		DbSchema:   dbConfig.DbSchema,
		DbName:     dbConfig.DbName,
		DbUser:     dbConfig.DbUser,
		DbPassword: dbConfig.DbPassword,
	}
}

func setSchema(schemaName string) {
	orm.SetTableNameInflector(func(name string) string {
		return schemaName + "." + name
	})
}
