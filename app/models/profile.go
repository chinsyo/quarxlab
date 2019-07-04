package models

import (
	"github.com/jinzhu/gorm"
)

type Profile struct {
	gorm.Model
	Nickname string `form:"nickname" json:"nickname"`
	Birthday string `form:"birthday" json:"birthday"`
	Gender   uint   `form:"gender" json:"gender"`
	UserID   uint   `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}
