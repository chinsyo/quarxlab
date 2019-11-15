package models

import (
	"github.com/jinzhu/gorm"
)

type Asset struct {
	gorm.Model
	FilePath string `form:"file_path" json:"file_path"`
	Author   string `form:"author" json:"author"`
	Visitor  string `form:"visitor" json:"visitor"`
	// UserID   uint `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}
