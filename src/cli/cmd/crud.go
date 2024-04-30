package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var defaultHost = "http://localhost:8080"

type File struct {
	StorageName string `json:"storage_name"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	FullPath    string `json:"-"`
}

func GetFile(fileId, host string) {
	if host == "" {
		host = defaultHost
	}
	uri := fmt.Sprintf("%s/%s/%s", host, FILE_ROUTE, fileId)

	resp, err := http.Get(uri)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var file File

	err = ParseResponse(&file, resp)
	if err != nil {
		log.Println("Could not parse response")
		log.Fatalln(err)
	}
	PrettyPrint(file)
}

func AddFile(filePath, host, dir string) {
	if host == "" {
		host = defaultHost
	}

	if filePath == "" {
		log.Fatalln("filename is not provided")
	}

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalln(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(file)

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	log.Println("✅filepath ", filepath.Base(filePath))
	fileContent, err := writer.CreateFormFile("file", filepath.Base(filePath))

	if err != nil {
		log.Fatalln(err)
	}

	err = writer.WriteField("path", dir)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = io.Copy(fileContent, file)

	err = writer.Close()

	if err != nil {
		log.Fatalln(err)
	}

	uri := fmt.Sprintf("%s/%s", host, FILE_ROUTE)

	request, err := http.NewRequest("POST", uri, &body)

	request.Header.Add("Content-Type", writer.FormDataContentType())

	var client http.Client

	response, err := client.Do(request)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var fileResponse File

	err = ParseResponse(&fileResponse, response)
	if err != nil {
		log.Println("Could not parse response")
		log.Fatalln(err)
	}

	PrettyPrint(fileResponse)
}

func ListDir(host, dirname string) {
	if host == "" {
		host = defaultHost
	}

	uri := fmt.Sprintf("%s/%s/%s", host, DIR_ROUTE, dirname)

	resp, err := http.Get(uri)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}
	var files []File

	err = ParseResponse(&files, resp)

	if err != nil {
		log.Println("Could not parse response")
		log.Fatalln(err)
	}

	for _, file := range files {
		PrettyPrint(file)
	}

}
