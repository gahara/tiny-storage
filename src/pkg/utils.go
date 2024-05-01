package pkg

import (
	"encoding/json"
	"fmt"
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
