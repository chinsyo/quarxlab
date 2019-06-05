package models

import (
    "github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title string `form:"title" json:"title"`
	Image string `form:"image" json:"image"`
	Content string `form:"content" json:"content"`
	AuthorId uint `form:"author_id" json:"author_id"`
	CategoryId uint `form:"category_id" json:"category_id"`
}
