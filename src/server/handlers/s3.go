package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"s3/src/server/helpers"
	"s3/src/server/storage"
)

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

	if err := ctx.SaveUploadedFile(file, storagePath+dirName+fileNameForStorage); err != nil {
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

func GetFiles(ctx *gin.Context) {
	var files []storage.File
	database := helpers.GetDB(ctx)

	if res := database.Find(&files); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res.Error)
		return
	}
	ctx.IndentedJSON(http.StatusOK, &files)
}

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

func MakeDir(ctx *gin.Context) {
	dirName := ctx.Query("dn")

	if _, err := os.Stat(dirName + "/"); os.IsNotExist(err) {
		err := os.MkdirAll(dirName, os.ModeDir)
		if err == nil {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Created dir with name " + dirName})
		} else {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
	} else {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Dir already exists"})
		return
	}

}

func GetDir(ctx *gin.Context) {
	dirName := ctx.Param("name")

	if _, err := os.Stat(dirName + "/"); !os.IsNotExist(err) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Dir does not exist"})
		return
	} else {
		// return all from db where path = dirName
	}
}
