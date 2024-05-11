package middleware

import (
	"github.com/gin-gonic/gin"
	"s3/src/internal/customTypes"
	"s3/src/internal/server/storage"
)

func SetDBMiddleware(envVars customTypes.EnvironmentalVariables) gin.HandlerFunc {
	dbHandler := storage.SetDb(envVars)

	return func(context *gin.Context) {
		context.Set("db", dbHandler)
		context.Next()
	}
}
