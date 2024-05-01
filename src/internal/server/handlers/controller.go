package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.POST("/files", AddFile)
	router.POST("/dirs", MakeDir)
	router.GET("/dirs/:name", GetDirInsides)
	router.GET("/files", GetFiles)
	router.GET("/files/:id", GetFile)
	router.DELETE("files/:id", DeleteFile)
	router.GET("/ping", Ping)
}
