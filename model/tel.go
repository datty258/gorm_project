package model

import "github.com/jinzhu/gorm"

type Tel struct {
	gorm.Model
	UserID     int
	Tel        string `gorm:"type:varchar(100);unique_index"`
	Subscribed bool
}

func GetTel(UserID interface{}) (Tel, error) {
	var tel Tel
	result := DB.Where("user_id = ?", UserID).First(&tel)
	return tel, result.Error
}
