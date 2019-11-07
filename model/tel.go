package model

import "github.com/jinzhu/gorm"

type Tel struct {
	gorm.Model
	UserID     int
	Tel        string `gorm:"type:varchar(100);unique_index"`
	Subscribed bool
}
