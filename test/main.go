package main

import (
	"database/sql"                     //标准库
	_ "github.com/go-sql-driver/mysql" //我们使用的mysql，需要导入相应驱动包，否则会报错, _ 表示不使用相关函数但是会自动执行init方法
	"log"
)

// 定义一个全局对象db
var db *sql.DB

func initDB() {
	var err error
	// 设置一下dns charset:编码方式 parseTime:是否解析time类型 loc:时区
	dsn := "root:123456@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=True&loc=Local"
	// 打开mysql驱动
	db, err = sql.Open("mysql", dsn) //不会检验用户名和密码是否正确
	if err != nil {
		log.Fatalln(err)
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Printf("Open failed,err: %v", err)
	}
	log.Println("DB connect success")
	return
}
func main() {
	//初始化连接
	initDB()
}
