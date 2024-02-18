package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/", AddFile)
	router.GET("/", GetFiles)
	router.GET("/:id", GetFile)
	router.DELETE("/:id", DeleteFile)
	router.GET("/ping", Ping)
}
