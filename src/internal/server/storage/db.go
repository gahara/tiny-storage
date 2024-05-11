package storage

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"s3/src/internal/constants"
	"s3/src/internal/customTypes"
	"s3/src/internal/server/helpers"
)

func Init(dsn string) *gorm.DB {
	if dsn == constants.GormDB {
		db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		return db
	} else {
		db, err := gorm.Open(postgres.Open(dsn))
		if err != nil {
			log.Fatalln(err)
		}
		return db
	}
}

func SetDb(environmentsVariables customTypes.EnvironmentalVariables) *gorm.DB {
	dbConnectionString := helpers.CreateDbConfig(helpers.ResolveDbConf(environmentsVariables.Env))

	dbHandler := Init(dbConnectionString)
	log.Println("connected to db: ", dbConnectionString)
	return dbHandler
}
