package storage

import (
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	Name string
	Path string
}
