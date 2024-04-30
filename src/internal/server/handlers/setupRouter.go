package handlers

import (
	"github.com/gin-gonic/gin"
	"s3/src/internal/server/helpers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(helpers.SetDBMiddleware())
	router.Use(helpers.SetEnvMiddleware())
	RegisterRoutes(router)
	return router
}
