package models

import (
	"github.com/jinzhu/gorm"
)

type Credential struct {
	gorm.Model 
	Username string
	Password string
}