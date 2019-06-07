package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Open() error{
	var err error
	db, err = gorm.Open("mysql", "root:123456@/zgoband?charset=utf8&parseTime=True&loc=Local")
	if(err != nil) {
		panic(err.Error())
	}
	db.LogMode(true)
	return nil
}

func Close() {
	if(db == nil) {
		fmt.Println("db is nil")
		return
	}
	err := db.Close()
	if(err != nil) {
		fmt.Println(err.Error())
	}
}