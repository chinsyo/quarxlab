package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Content   string `form:"content" json:"content"`
	AuthorID  uint   `form:"author" json:"author"`
	ArticleID uint   `gorm:"foreignkey:ArticleID;association_foreignkey:ID"`
}
