package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"s3/src/internal/customTypes"
	"s3/src/internal/server/helpers"
)

func SetEnvMiddleware(envVars customTypes.EnvironmentalVariables) gin.HandlerFunc {
	log.Println("set environmental variables")
	return func(context *gin.Context) {
		context.Set(helpers.ENIRONMENTAL_VARIABLES_KEY, envVars)
		context.Next()
	}
}
