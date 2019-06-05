package models 

import (
    "github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Content string `form:"content" json:"content"`
	Author int `form:"author" json:"author"`
	Article int `form:"article" json:"article"`
}

