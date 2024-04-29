package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"s3/src/server/helpers"
	"s3/src/server/storage"
)

// AddFile  godoc
// @Summary add a file
// @Description Get all files across all dirs
// @Tags        files
// @Produce     json
// @Param       file formData file true "File to store"
// @Param       path  formData string true "path to store things"
// @Success     200  {array} storage.File
// @Failure     500
// @Router      /files [post]
func AddFile(ctx *gin.Context) {
	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(helpers.EnvironmentalVariables).StoragePath

	form, err := ctx.MultipartForm()
	log.Println(form)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}
	file := form.File["file"][0]
	dirName := form.Value["path"][0]

	extension := filepath.Ext(file.Filename)
	fileNameForStorage := uuid.New().String() + extension
	dirPath := fmt.Sprintf("%s/%s", storagePath, dirName)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": helpers.DirDoesNotExist})
		return
	}

	fullPath := fmt.Sprintf("%s/%s/%s", storagePath, dirName, fileNameForStorage)

	if err := ctx.SaveUploadedFile(file, fullPath); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	database := helpers.GetDB(ctx)

	dbFile := storage.File{
		Name:        file.Filename,
		StorageName: fileNameForStorage,
		Path:        dirName,
		FullPath:    storagePath + fileNameForStorage,
	}

	res := database.Create(&dbFile)

	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res.Error)
		return
	}
	ctx.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "filename": dbFile.Name,
		"stored_name": dbFile.StorageName, "dir": dbFile.Path})
}

// GetFile  godoc
// @Summary Get file by id
// @Description Get file by id
// @Tags        files
// @Produce     json
// @Success     200  {array} storage.File
// @Failure     500
// @Router      /files/{id} [get]
func GetFile(ctx *gin.Context) {
	database := helpers.GetDB(ctx)
	id := ctx.Param("id")

	var file storage.File

	if res := database.First(&file, id); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, res.Error)
		return
	}

	ctx.IndentedJSON(http.StatusOK, &file)
}

// GetFiles  godoc
// @Summary Get all files
// @Description Get all files across all dirs
// @Tags        files
// @Produce     json
// @Success     200  {array} storage.File
// @Failure     500
// @Router      /files [get]
func GetFiles(ctx *gin.Context) {
	var files []storage.File
	database := helpers.GetDB(ctx)

	if res := database.Find(&files); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res.Error)
		return
	}
	ctx.IndentedJSON(http.StatusOK, &files)
}

// DeleteFile  godoc
// @Summary Delete file
// @Description Delete file by id
// @Tags        files
// @Produce     json
// @Success     200 {object} storage.File
// @Failure     500
// @Router      /files [delete]
func DeleteFile(ctx *gin.Context) {
	id := ctx.Param("id")
	var fileRecord storage.File
	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(helpers.EnvironmentalVariables).StoragePath

	database := helpers.GetDB(ctx)

	if res := database.First(&fileRecord, id); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, res.Error)
		return
	}

	if err := os.Remove(storagePath + fileRecord.StorageName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	database.Delete(&fileRecord)

	ctx.IndentedJSON(http.StatusOK, &fileRecord)
}

// MakeDir  godoc
// @Summary Create dir
// @Description Create dir to store files
// @Tags        dirs
// @Produce     json
// @Success     200  {object} any
// @Failure     500
// @Router      /dirs [post]
func MakeDir(ctx *gin.Context) {
	type dirReq struct {
		Name string
	}
	var dirBody dirReq

	err := ctx.BindJSON(&dirBody)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(helpers.EnvironmentalVariables).StoragePath
	dirPath := fmt.Sprintf("%s/%s/", storagePath, dirBody.Name)

	dirCreationText, dirCreationStatus, err := helpers.CreateDir(dirPath, dirBody.Name)

	if err != nil {
		ctx.AbortWithStatusJSON(dirCreationStatus, err.Error())
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": dirCreationText})
}

// GetDir  godoc
// @Summary Get dir
// @Description Get contents of the dir
// @Tags        dirs
// @Produce     json
// @Success     200  {array} storage.File
// @Failure     500
// @Router      /dirs [get]
func GetDir(ctx *gin.Context) {
	dirName := ctx.Param("name")
	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(helpers.EnvironmentalVariables).StoragePath
	dirPath := fmt.Sprintf("%s/%s/", storagePath, dirName)
	println(dirPath)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": helpers.DirDoesNotExist})
		return
	} else {
		var files []storage.File
		database := helpers.GetDB(ctx)

		if res := database.Where("path = ?", dirName).Find(&files); res.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res.Error)
			return
		}

		ctx.IndentedJSON(http.StatusOK, &files)
	}
}
