package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database() {
	db, err := gorm.Open("mysql", "root:root@/gorm_project?parseTime=true")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
	migration()
}
