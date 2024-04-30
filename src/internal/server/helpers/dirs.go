package helpers

import (
	"fmt"
	"net/http"
	"os"
	"s3/src/internal/constants"
)

func CreateDir(dirPath, dirName string) (string, int, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err == nil {
			return fmt.Sprintf("%s %s", constants.DirCreated, dirName), http.StatusOK, nil
		} else {
			return constants.SomethingWentWrong, http.StatusInternalServerError, err
		}
	} else {
		return constants.DirAlreadyExists, http.StatusConflict, fmt.Errorf("%s %s", constants.DirAlreadyExists, dirName)
	}
}
