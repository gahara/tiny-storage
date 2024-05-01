package main

import (
	"log"
	"s3/src/internal/customTypes"
	"s3/src/internal/server/storage"
)

func main() {
	log.Println("Starting migration")
	database := storage.Init("gorm.db")

	err := database.AutoMigrate(&customTypes.File{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Migration done")
}
