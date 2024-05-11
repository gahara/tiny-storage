package handlers

import (
	"github.com/gin-gonic/gin"
	"s3/src/internal/customTypes"
	"s3/src/internal/server/middleware"
)

func SetupRouter(envVars customTypes.EnvironmentalVariables) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.SetEnvMiddleware(envVars))
	router.Use(middleware.SetDBMiddleware(envVars))
	RegisterRoutes(router)
	return router
}
