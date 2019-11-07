package model

import (
	"github.com/jinzhu/gorm"
)

type Address struct {
	gorm.Model
	Name   string
	Tel    string
	Detail string
	UserID uint
}
