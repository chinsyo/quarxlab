package models

import (
    "github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" form:"username" json:"username"`
	Password string `from:"password" json:"password"`
	Profile Profile
	ProfileID uint
	Assets []Asset  `form:"-" json:"-"`
}
