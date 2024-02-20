package cmd

import (
	"io"
	"log"
	"net/http"
)

var defaultHost = "http://localhost:8080/"

func Get(fileId string, host string) {
	if host == "" {
		host = defaultHost
	}
	resp, err := http.Get(host + fileId)
	if err != nil {
		log.Fatalln(err)
	}

	if body, err := io.ReadAll(resp.Body); err != nil {
		log.Fatalln(err)
	} else {
		println(string(body))
	}
}
