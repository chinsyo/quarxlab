package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title string `form:"title" json:"title"`
	Image string `form:"image" json:"image"`
	Content string `form:"content" json:"content"`
	Comments []Comment `form:"-" json:"-"`
	AuthorId uint `form:"author_id" json:"author_id"`
	CategoryId uint `gorm:"foreignkey:ArticleID;association_foreignkey:ID" form:"category_id" json:"category_id"`
}
