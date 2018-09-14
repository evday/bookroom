package models

import (
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID uint64 `json:"id" gorm:"primary_key"`  	
	CreateAt time.Time  `json:"create_time"`  //创建时间
}

var db *gorm.DB

//设置数据库连接
func SetDB() {
	var err error
	var connection string = "root:@tcp(localhost:3306)/bookroom?parseTime=True"
	db, err = gorm.Open("mysql", connection)
	if err != nil {
		panic(err)
	}
}


func GetDB() *gorm.DB {
	return db
}


//自动提交
func AutoMigrate() {
	db.AutoMigrate(&User{}, &Record{},&Room{},&Equipment{},&EquiRoom{})
}