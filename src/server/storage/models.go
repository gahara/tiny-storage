package storage

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model `json:"-"`
	Name       string `json:"name"`
	Path       string `json:"path"`
}
