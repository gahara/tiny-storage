package helpers

import (
	"fmt"
	"os"
	"s3/src/internal/constants"
)

func CreateDir(dirPath, dirName string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		err := os.MkdirAll(dirPath, os.ModePerm)
		if err == nil {
			return nil
		} else {
			return err
		}
	} else {
		return fmt.Errorf("%s %s", constants.DirAlreadyExists, dirName)
	}
}
