package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `form:"username" json:"username"`
	Assets []Asset  `form:"-" json:"-"`
}
