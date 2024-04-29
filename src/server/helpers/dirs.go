package helpers

import (
	"fmt"
	"net/http"
	"os"
)

func CreateDir(dirPath, dirName string) (string, int, error) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err == nil {
			return fmt.Sprintf("%s%s", DirCreated, dirName), http.StatusOK, nil
		} else {
			return SomethingWentWrong, http.StatusInternalServerError, err
		}
	} else {
		return DirAlreadyExists, http.StatusConflict, fmt.Errorf("%s %s", DirAlreadyExists, dirName)
	}
}
