package handlers

import (
	"github.com/gin-gonic/gin"
	"s3/src/internal/server/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandler())
	router.Use(middleware.SetDBMiddleware())
	router.Use(middleware.SetEnvMiddleware())
	RegisterRoutes(router)
	return router
}
