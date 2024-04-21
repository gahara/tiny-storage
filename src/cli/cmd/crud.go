package cmd

import (
	"bytes"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var defaultHost = "http://localhost:8080/"

func Get(fileId string, host string) {
	if host == "" {
		host = defaultHost
	}
	resp, err := http.Get(host + FILE_ROUTE + fileId)
	if err != nil {
		log.Fatalln(err)
	}

	if body, err := io.ReadAll(resp.Body); err != nil {
		log.Fatalln(err)
	} else {
		println(string(body))
	}
}

func Post(filePath, host, dir string) (string, error) {
	if host == "" {
		host = defaultHost
	}

	if filePath == "" {
		return "", errors.New("Filename is not provided")
	}

	if dir != "" && dir[len(dir)-1] != '/' {
		dir = dir + "/"
	}

	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	log.Println("filepath✅", filepath.Base(filePath))
	fileContent, err := writer.CreateFormFile("file", filepath.Base(filePath))

	if err != nil {
		return "", err
	}

	err = writer.WriteField("path", dir)

	if err != nil {
		return "", err
	}

	_, err = io.Copy(fileContent, file)

	err = writer.Close()

	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("POST", host+FILE_ROUTE, body)

	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if responseBody, err := io.ReadAll(response.Body); err != nil {
		return "", err
	} else {
		println(string(responseBody))
		return "", nil
	}
}
