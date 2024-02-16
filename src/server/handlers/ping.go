package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "OK"})
}
