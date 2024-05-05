package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"s3/src/internal/server/helpers"
)

func SetEnvMiddleware() gin.HandlerFunc {
	envs := helpers.GetEnvironmentalVariables()
	log.Println("Get environmental variables")
	return func(context *gin.Context) {
		context.Set(helpers.ENIRONMENTAL_VARIABLES_KEY, envs)
		context.Next()
	}
}
