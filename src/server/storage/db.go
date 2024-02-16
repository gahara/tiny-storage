package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Init(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
