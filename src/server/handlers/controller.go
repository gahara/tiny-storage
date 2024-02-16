package handlers

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.POST("/", AddFile)
	router.GET("/", GetFiles)
	router.GET("/ping", Ping)
}
