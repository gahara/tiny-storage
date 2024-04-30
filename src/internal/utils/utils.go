package utils

import (
	"fmt"
	"net/http"
	"s3/src/internal/constants"
)

func DirExistsRequest(host, dirname string) error {
	uri := fmt.Sprintf("%s/%s/%s", host, constants.DIR_ROUTE, dirname)

	resp, err := http.Get(uri)

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%s %s", constants.DirDoesNotExist, dirname)
		return err
	}

	return nil
}
