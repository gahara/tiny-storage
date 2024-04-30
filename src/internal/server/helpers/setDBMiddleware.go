package helpers

import (
	"github.com/gin-gonic/gin"
	"log"
	"s3/src/internal/server/storage"
)

const (
	testDB = "testgorm.db"
	prodDB = "gorm.db"
)

func SetDBMiddleware() gin.HandlerFunc {
	dbHandler := storage.Init(prodDB)
	log.Println("connected to db: ", prodDB)
	return func(context *gin.Context) {
		context.Set("db", dbHandler)
		context.Next()
	}
}
