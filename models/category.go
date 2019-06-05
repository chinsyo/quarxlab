package models

import (
    "github.com/jinzhu/gorm"
)

type Category struct {
    gorm.Model
    Name string `form:"name" json:"name"`
    Articles []Article `form:"-" json:"-"`
}
