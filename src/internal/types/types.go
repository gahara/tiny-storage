package types

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model  `json:"-"`
	StorageName string `json:"storage_name"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	FullPath    string `json:"-"`
}

type Response[T any] struct {
	ResponseKey     string
	ResponseMessage string
	Data            T
}
