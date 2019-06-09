package models

import (
	"github.com/jinzhu/gorm"
)

type Asset struct {
	gorm.Model
	FilePath string
	UserID uint `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}