package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"s3/src/internal/customTypes"
	"s3/src/internal/utils"
	"s3/src/pkg"
)

var defaultHost = "http://localhost:8080"

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
	var fileResp customTypes.FilesResponse

	err = pkg.ParseResponse(&fileResp, resp)
	if err != nil {
		log.Println("Could not parse response")
		log.Fatalln(err)
	}
	pkg.PrettyPrint(fileResp.Results.Data)
}

func AddFile(filePath, host, dir string) {
	if host == "" {
		host = defaultHost
	}

	err := utils.DirExistsRequest(host, dir)

	if err != nil {
		log.Fatalln("dir does not exist")
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

	var fileResponse customTypes.FilesResponse

	err = pkg.ParseResponse(&fileResponse, response)
	if err != nil {
		log.Println("Could not parse response")
		log.Fatalln(err)
	}

	pkg.PrettyPrint(fileResponse.Results.Data)
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

	var fileResp customTypes.FilesResponse

	err = pkg.ParseResponse(&fileResp, resp)

	if err != nil {
		log.Println("Could not parse response")
		log.Fatalln(err)
	}
	files := fileResp.Results.Data

	for _, file := range files {
		pkg.PrettyPrint(file)
	}

}

func AddDir(host, dirName string) {
	if host == "" {
		host = defaultHost
	}

	uri := fmt.Sprintf("%s/%s", host, DIR_ROUTE)

	values := map[string]string{"name": dirName}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post(uri, "application/json", bytes.NewBuffer(jsonValue))

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Println(bodyString)
}
