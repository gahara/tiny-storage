package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"path/filepath"
	"s3/src/server/helpers"
	"s3/src/server/storage"
)

func AddFile(ctx *gin.Context) {

	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(helpers.EnvironmentalVariables).StoragePath
	log.Println("storage", storagePath)

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	extension := filepath.Ext(file.Filename)

	fileNameForStorage := uuid.New().String() + extension

	if err := ctx.SaveUploadedFile(file, storagePath+fileNameForStorage); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	database := helpers.GetDB(ctx)

	dbFile := storage.File{
		Name:        file.Filename,
		StorageName: fileNameForStorage,
		Path:        storagePath + fileNameForStorage,
	}

	res := database.Create(&dbFile)

	if res.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, res.Error)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "filename": dbFile.Name, "stored_name": dbFile.StorageName})
}

func GetFile(ctx *gin.Context) {
	database := helpers.GetDB(ctx)
	id := ctx.Param("id")

	var file storage.File

	if res := database.First(&file, id); res.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, res.Error)
		return
	}

	ctx.IndentedJSON(http.StatusOK, &file)
}

func GetFiles(ctx *gin.Context) {
	var files []storage.File
	database := helpers.GetDB(ctx)

	if res := database.Find(&files); res.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, res.Error)
	}
	ctx.IndentedJSON(http.StatusOK, &files)
}

func DeleteFile(ctx *gin.Context) {

}
