package helpers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func SetEnvMiddleware() gin.HandlerFunc {
	envs := GetEnvironmentalVariables()
	log.Println("Get environmental variables")
	return func(context *gin.Context) {
		context.Set(ENIRONMENTAL_VARIABLES_KEY, envs)
		context.Next()
	}
}
