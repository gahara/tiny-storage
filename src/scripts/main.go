package main

import (
	"log"
	"s3/src/internal/customTypes"
	"s3/src/internal/server/helpers"
	"s3/src/internal/server/storage"
)

const production = "production"
const testing = "testing"

func main() {
	log.Println("Starting migration")
	dbConnectionString := helpers.CreateDbConfig(helpers.ResolveDbConf(production))
	database := storage.Init(dbConnectionString)

	err := database.AutoMigrate(&customTypes.File{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Migration done")
}
