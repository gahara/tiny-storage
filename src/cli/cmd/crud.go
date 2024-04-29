package cmd

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

var defaultHost = "http://localhost:8080"

func Get(fileId string, host string) {
	if host == "" {
		host = defaultHost
	}
	uri := fmt.Sprintf("%s/%s/%s", host, FILE_ROUTE, fileId)
	resp, err := http.Get(uri)
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

	file, err := os.Open(filePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	log.Println("✅filepath ", filepath.Base(filePath))
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

	uri := fmt.Sprintf("%s/%s", host, FILE_ROUTE)

	request, err := http.NewRequest("POST", uri, body)

	request.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	if responseBody, err := io.ReadAll(response.Body); err != nil {
		return "", err
	} else {
		println(string(responseBody))
		return "", nil
	}
}
