package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
// @Message Get all files across all dirs
// @Tags        files
// @Produce     json
// @Param       file formData file true "File to store"
// @Param       path  formData string true "path to store things"
// @Success     200 {object} customTypes.FilesResponse
// @Failure     500 {object} pkg.HttpError
// @Failure     400 {object} pkg.HttpError
// @Router      /files [post]
func AddFile(ctx *gin.Context) {
	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(customTypes.EnvironmentalVariables).StoragePath

	form, err := ctx.MultipartForm()

	if err != nil {
		ctx.Error(pkg.BuildError(constants.BadRequest, http.StatusBadRequest, err.Error()))
		return
	}

	file, dirName, err := pkg.DeconStrucMultipartForm(form)

	if err != nil {
		ctx.Error(pkg.BuildError(constants.BadRequest, http.StatusBadRequest, err.Error()))
		return
	}

	extension := filepath.Ext(file.Filename)
	fileNameForStorage := uuid.New().String() + extension
	dirPath := fmt.Sprintf("%s/%s", storagePath, dirName)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		ctx.Error(pkg.BuildError(constants.DirDoesNotExist, http.StatusNotFound, err.Error()))
		return
	}

	fullPath := fmt.Sprintf("%s/%s/%s", storagePath, dirName, fileNameForStorage)

	if err := ctx.SaveUploadedFile(file, fullPath); err != nil {
		ctx.Error(pkg.BuildError(constants.SomethingWentWrong, http.StatusInternalServerError, err.Error()))
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
		ctx.Error(pkg.BuildError(constants.SomethingWentWrong, http.StatusInternalServerError, err.Error()))
		return
	}

	response := pkg.BuildFilesResponse(constants.StatusTextOk, []customTypes.File{dbFile})
	ctx.IndentedJSON(http.StatusOK, response)
}

// GetFile  godoc
// @Summary Get file by id
// @Message Get file by id
// @Tags        files
// @Produce     json
// @Success     200  {object} customTypes.FilesResponse
// @Failure     404 {object} pkg.HttpError
// @Router      /files/{id} [get]
func GetFile(ctx *gin.Context) {
	database := helpers.GetDB(ctx)
	id := ctx.Param("id")

	var file customTypes.File

	if res := database.First(&file, id); res.Error != nil {
		ctx.Error(pkg.BuildError(constants.NotFoundMessage, http.StatusNotFound, res.Error.Error()))
		return
	}

	response := pkg.BuildFilesResponse(constants.StatusTextOk, []customTypes.File{file})

	ctx.IndentedJSON(http.StatusOK, response)
}

// GetFiles  godoc
// @Summary Get all files
// @Message Get all files across all dirs
// @Tags        files
// @Produce     json
// @Success     200 {object} customTypes.FilesResponse
// @Failure     500 {object} pkg.HttpError
// @Router      /files [get]
func GetFiles(ctx *gin.Context) {
	var files []customTypes.File
	database := helpers.GetDB(ctx)

	if res := database.Find(&files); res.Error != nil {
		ctx.Error(pkg.BuildError(constants.SomethingWentWrong, http.StatusInternalServerError, res.Error.Error()))
		return
	}

	response := pkg.BuildFilesResponse(constants.StatusTextOk, files)
	ctx.IndentedJSON(http.StatusOK, response)
}

// DeleteFile  godoc
// @Summary Delete file
// @Message Delete file by id
// @Tags        files
// @Produce     json
// @Success     200 {object} customTypes.FilesResponse
// @Failure     500 {object} pkg.HttpError
// @Failure     404 {object} pkg.HttpError
// @Router      /files [delete]
func DeleteFile(ctx *gin.Context) {
	id := ctx.Param("id")
	var fileRecord customTypes.File
	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(customTypes.EnvironmentalVariables).StoragePath

	database := helpers.GetDB(ctx)

	if res := database.First(&fileRecord, id); res.Error != nil {
		ctx.Error(pkg.BuildError(constants.NotFoundMessage, http.StatusNotFound, res.Error.Error()))
		return
	}

	if err := os.Remove(storagePath + fileRecord.StorageName); err != nil {
		ctx.Error(pkg.BuildError(constants.SomethingWentWrong, http.StatusInternalServerError, err.Error()))
		return
	}

	database.Delete(&fileRecord)

	response := pkg.BuildFilesResponse(constants.StatusTextDeleted, customTypes.Files{fileRecord})

	ctx.IndentedJSON(http.StatusOK, response)
}

// MakeDir  godoc
// @Summary Create dir
// @Message Create dir to store files
// @Tags        dirs
// @Produce     json
// @Success     200  {object} any
// @Failure     500 {object} pkg.HttpError
// @Failure     400 {object} pkg.HttpError
// @Router      /dirs [post]
func MakeDir(ctx *gin.Context) {
	type dirReq struct {
		Name string
	}
	var dirBody dirReq

	err := ctx.BindJSON(&dirBody)
	if err != nil {
		ctx.Error(pkg.BuildError(constants.BadRequest, http.StatusBadRequest, err.Error()))
		return
	}

	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(customTypes.EnvironmentalVariables).StoragePath

	dirPath := fmt.Sprintf("%s/%s/", storagePath, dirBody.Name)

	err = helpers.CreateDir(dirPath, dirBody.Name)

	if err != nil {
		ctx.Error(pkg.BuildError(constants.BadRequest, http.StatusBadRequest, err.Error()))
		return
	}

	//TODO: bring this to the common format
	pkg.ResponseOK(ctx, constants.DirCreated)
}

// GetDirInsides  godoc
// @Summary Get dir
// @Message Get contents of the dir
// @Tags        dirs
// @Produce     json
// @Success     200 {object} customTypes.FilesResponse
// @Failure     500 {object} pkg.HttpError
// @Failure     404 {object} pkg.HttpError
// @Router      /dirs/{name} [get]
func GetDirInsides(ctx *gin.Context) {
	dirName := ctx.Param("name")
	storagePath := ctx.MustGet(helpers.ENIRONMENTAL_VARIABLES_KEY).(customTypes.EnvironmentalVariables).StoragePath
	dirPath := fmt.Sprintf("%s/%s/", storagePath, dirName)
	println(dirPath)

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		ctx.Error(pkg.BuildError(constants.NotFoundMessage, http.StatusNotFound, err.Error()))
		return
	} else {
		var files []customTypes.File
		database := helpers.GetDB(ctx)

		if res := database.Where("path = ?", dirName).Find(&files); res.Error != nil {
			ctx.Error(pkg.BuildError(constants.SomethingWentWrong, http.StatusInternalServerError, res.Error.Error()))
			return
		}
		response := pkg.BuildFilesResponse(constants.StatusTextOk, files)
		ctx.IndentedJSON(http.StatusOK, response)
	}
}
