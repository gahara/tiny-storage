package main

import (
	"log"
	"s3/src/server/storage"
)

func main() {
	log.Println("Starting migration")
	database := storage.Init("gorm.db")

	err := database.AutoMigrate(&storage.File{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Migration done")
}
