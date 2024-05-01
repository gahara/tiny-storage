package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"s3/src/internal/constants"
	"s3/src/internal/customTypes"
	"s3/src/internal/server/helpers"
	"s3/src/pkg"
)

// AddFile  godoc
// @Summary add a file
// @Description Get all files across all dirs
// @Tags        files
// @Produce     json
// @Param       file formData file true "File to store"
// @Param       path  formData string true "path to store things"
// @Success     200 {object} customTypes.FilesResponse
// @Failure     500
// @Router      /files [post]
func AddFile(ctx *gin.Context) {
	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(helpers.EnvironmentalVariables).StoragePath

	form, err := ctx.MultipartForm()

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
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": constants.DirDoesNotExist})
		return
	}

	fullPath := fmt.Sprintf("%s/%s/%s", storagePath, dirName, fileNameForStorage)

	if err := ctx.SaveUploadedFile(file, fullPath); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	database := helpers.GetDB(ctx)

	dbFile := customTypes.File{
		Name:        file.Filename,
		StorageName: fileNameForStorage,
		Path:        dirName,
		FullPath:    storagePath + fileNameForStorage,
	}

	dbResult := database.Create(&dbFile)

	if dbResult.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dbResult.Error)
		return
	}

	response := pkg.BuildFilesResponse(constants.StatusTextOk, []customTypes.File{dbFile})
	ctx.IndentedJSON(http.StatusOK, response)
}

// GetFile  godoc
// @Summary Get file by id
// @Description Get file by id
// @Tags        files
// @Produce     json
// @Success     200  {object} customTypes.FilesResponse
// @Failure     500
// @Router      /files/{id} [get]
func GetFile(ctx *gin.Context) {
	database := helpers.GetDB(ctx)
	id := ctx.Param("id")

	var file customTypes.File

	if res := database.First(&file, id); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, res.Error)
		return
	}

	response := pkg.BuildFilesResponse(constants.StatusTextOk, []customTypes.File{file})

	ctx.IndentedJSON(http.StatusOK, response)
}

// GetFiles  godoc
// @Summary Get all files
// @Description Get all files across all dirs
// @Tags        files
// @Produce     json
// @Success     200 {object} customTypes.FilesResponse
// @Failure     500
// @Router      /files [get]
func GetFiles(ctx *gin.Context) {
	var files []customTypes.File
	database := helpers.GetDB(ctx)

	if res := database.Find(&files); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res.Error)
		return
	}

	response := pkg.BuildFilesResponse(constants.StatusTextOk, files)
	ctx.IndentedJSON(http.StatusOK, response)
}

// DeleteFile  godoc
// @Summary Delete file
// @Description Delete file by id
// @Tags        files
// @Produce     json
// @Success     200 {object} customTypes.FilesResponse
// @Failure     500
// @Router      /files [delete]
func DeleteFile(ctx *gin.Context) {
	id := ctx.Param("id")
	var fileRecord customTypes.File
	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(helpers.EnvironmentalVariables).StoragePath

	database := helpers.GetDB(ctx)

	if res := database.First(&fileRecord, id); res.Error != nil {
		log.Println(res.Error)
		ctx.AbortWithStatusJSON(http.StatusNotFound, res.Error)
		return
	}

	if err := os.Remove(storagePath + fileRecord.StorageName); err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	database.Delete(&fileRecord)

	response := pkg.BuildFilesResponse(constants.StatusTextDeleted, customTypes.Files{fileRecord})

	ctx.IndentedJSON(http.StatusOK, response)
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
	pkg.ResponseOK(ctx, dirCreationText)
}

// GetDirInsides  godoc
// @Summary Get dir
// @Description Get contents of the dir
// @Tags        dirs
// @Produce     json
// @Success     200 {object} customTypes.FilesResponse
// @Failure     500
// @Router      /dirs/{name} [get]
func GetDirInsides(ctx *gin.Context) {
	dirName := ctx.Param("name")
	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(helpers.EnvironmentalVariables).StoragePath
	dirPath := fmt.Sprintf("%s/%s/", storagePath, dirName)
	println(dirPath)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": constants.DirAlreadyExists})
		return
	} else {
		var files []customTypes.File
		database := helpers.GetDB(ctx)

		if res := database.Where("path = ?", dirName).Find(&files); res.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, res.Error)
			return
		}
		response := pkg.BuildFilesResponse(constants.StatusTextOk, files)

		ctx.IndentedJSON(http.StatusOK, response)
	}
}
