package pkg

import "fmt"

type HttpError struct {
	Description string `json:"description"`
	StatusCode  int    `json:"statusCode"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf("description: %s, status %d", e.Description, e.StatusCode)
}

func BuildError(description string, statusCode int) HttpError {
	return HttpError{
		Description: description,
		StatusCode:  statusCode,
	}
}
