package main

import (
	"Lesson_5/homework/lv1/model"
	"database/sql" //标准库
	"fmt"
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

	// 创建数据库
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS student")
	if err != nil {
		log.Fatalln(err)
	}

	// 切换到 student 数据库
	_, err = db.Exec("USE student")
	if err != nil {
		log.Fatalln(err)
	}

	// 创建数据表 student
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS student (
			id INT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(20),
			age INT
		)
	`)
	if err != nil {
		log.Fatalln(err)
	}
	return
}

// 插入十条记录
func insertStudents() {
	students := []model.Student{
		{ID: 1, Name: "zty", Age: 20},
		{ID: 2, Name: "ty", Age: 21},
		{ID: 3, Name: "Zty", Age: 19},
		{ID: 4, Name: "yxh", Age: 18},
		{ID: 5, Name: "YuanShen", Age: 18},
		{ID: 6, Name: "YUAN", Age: 18},
		{ID: 7, Name: "XinHao", Age: 18},
		{ID: 8, Name: "wx", Age: 20},
		{ID: 9, Name: "MJ", Age: 22},
		{ID: 10, Name: "hhz", Age: 19},
	}

	for _, student := range students {
		_, err := db.Exec("INSERT INTO student (id, name, age) VALUES (?, ?, ?)", student.ID, student.Name, student.Age)
		if err != nil {
			fmt.Printf("insert failed, err:%v\n", err)
			return
		}
		log.Println("insert success")
	}
}

// 全部读出并打印
func readStudents() {
	rows, err := db.Query("SELECT id, name, age FROM student")
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var students []model.Student
	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		students = append(students, student)
	}

	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", student.ID, student.Name, student.Age)
	}
}

func main() {
	//初始化连接
	initDB()
	insertStudents()
	readStudents()
}
