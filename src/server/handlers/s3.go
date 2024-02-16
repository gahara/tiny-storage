package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"s3/src/server/storage"
)

func AddFile(ctx *gin.Context) {
	var file storage.File
	err := ctx.BindJSON(&file)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	database := ctx.MustGet("db").(*gorm.DB)
	res := database.Create(&file)

	if res.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, res.Error)
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "file": file.Name})
}

func GetFile(ctx *gin.Context) {

}

func GetFiles(ctx *gin.Context) {
	var files []storage.File
	database := ctx.MustGet("db").(*gorm.DB)

	if res := database.Find(&files); res.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, res.Error)
	}
	ctx.IndentedJSON(http.StatusOK, &files)
}

func DeleteFile(ctx *gin.Context) {

}
