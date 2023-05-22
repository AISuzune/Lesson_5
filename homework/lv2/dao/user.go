package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// GetDB 返回一个连接到 MySQL 数据库的实例
func GetDB() (*sql.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("DB connect success")
	return db, nil
}

// AddUser 向数据库中添加用户
func AddUser(username, password string) error {
	db, err := GetDB()
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	// 准备插入语句
	insertQuery := "INSERT INTO users (username, password) VALUES (?, ?)"

	// 执行插入操作
	_, err = db.Exec(insertQuery, username, password)
	if err != nil {
		fmt.Printf("Add failed, err:%v\n", err)
		return err
	}
	log.Println("Add success")
	return nil
}

// SelectUser 检查用户是否存在于数据库中
func SelectUser(username string) (bool, error) {
	db, err := GetDB()
	if err != nil {
		return false, err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	// 准备查询语句
	selectQuery := "SELECT COUNT(*) FROM users WHERE username = ?"
	var count int

	// 执行查询并将结果存储到 count 变量中
	err = db.QueryRow(selectQuery, username).Scan(&count)
	if err != nil {
		return false, err
	}

	// 如果 count 大于 0，则用户存在
	// 若没有这个用户返回 false，反之返回 true
	if count > 0 {
		return true, nil
	}
	return false, nil
}

// SelectPasswordFromUsername 根据用户名从数据库中获取密码
func SelectPasswordFromUsername(username string) (string, error) {
	db, err := GetDB()
	if err != nil {
		return "", err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	// 准备查询语句
	selectQuery := "SELECT password FROM users WHERE username = ?"
	var password string

	// 执行查询并将结果存储到 password 变量中
	err = db.QueryRow(selectQuery, username).Scan(&password)
	if err != nil {
		return "", err
	}

	return password, nil
}
