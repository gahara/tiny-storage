package main

import (
	"log"
	"s3/src/internal/server/storage"
	"s3/src/internal/types"
)

func main() {
	log.Println("Starting migration")
	database := storage.Init("gorm.db")

	err := database.AutoMigrate(&types.File{})

	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Migration done")
}
