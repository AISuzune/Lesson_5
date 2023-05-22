package dao

import (
	"Lesson_5/example/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func initDB() {
	var err error
	dsn := "root:123456@tcp(127.0.0.1:3306)/school?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := dsn.DSN
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)
	log.Println("connect success:")
}

func insert(st model.Student) {
	res := db.Table("student").Create(&st)
	if res.Error != nil {
		log.Println("insert err:", res.Error)
	}
	log.Println("insert success")
}
