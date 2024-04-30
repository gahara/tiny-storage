package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseResponse[T any](target *T, response *http.Response) error {
	return json.NewDecoder(response.Body).Decode(&target)
}

func PrettyPrint(target interface{}) {
	prettyTarget, _ := json.MarshalIndent(target, "", "\t")
	fmt.Println(string(prettyTarget))
}
