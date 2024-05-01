package customTypes

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
