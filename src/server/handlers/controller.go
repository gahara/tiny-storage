package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/files", AddFile)
	router.POST("/dirs", MakeDir)
	router.GET("/dirs/:name", GetDir)
	router.GET("/files", GetFiles)
	router.GET("/files/:id", GetFile)
	router.DELETE("files/:id", DeleteFile)
	router.GET("/ping", Ping)
}
