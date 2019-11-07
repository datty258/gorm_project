package model

import "github.com/jinzhu/gorm"

type Language struct {
	gorm.Model
	Name string `gorm:"index:idx_name_code"` 
	Code string `gorm:"index:idx_name_code"`
}
