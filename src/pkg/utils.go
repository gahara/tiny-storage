package pkg

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"reflect"
)

func ParseResponse[T any](target *T, response *http.Response) error {
	return json.NewDecoder(response.Body).Decode(&target)
}

func PrettyPrint(target interface{}) {
	prettyTarget, _ := json.MarshalIndent(target, "", "\t")
	fmt.Println(string(prettyTarget))
}

func isSliceOrArray[T any](item T) bool {
	t := reflect.TypeOf(item).Kind()
	if t == reflect.Array || t == reflect.Slice {
		return true
	}
	return false
}

func DeconStrucMultipartForm(form *multipart.Form) (*multipart.FileHeader, string, error) {
	file := form.File["file"]
	dir := form.Value["path"]

	if file != nil && dir != nil {
		return file[0], dir[0], nil
	}

	return nil, "", fmt.Errorf("%s %s", "File or path not sent")
}
