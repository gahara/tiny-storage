package pkg

import "fmt"

type HttpError struct {
	Message         string `json:"description"`
	StatusCode      int    `json:"statusCode"`
	LongDescription string `json:"-"`
}

func (e HttpError) Error() string {
	return fmt.Sprintf("⛔ message: %s, status %d, description: %s", e.Message, e.StatusCode, e.LongDescription)
}

func BuildError(message string, statusCode int, desc ...string) HttpError {
	longDescription := ""
	if len(desc) != 0 {
		longDescription = desc[0]
	}

	return HttpError{
		Message:         message,
		StatusCode:      statusCode,
		LongDescription: longDescription,
	}
}
