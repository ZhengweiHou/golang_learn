package main

import (
	"fmt"
	"log"
	"testing"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dburl = "root:root@tcp(127.0.0.1:3316)/dbsynctest?parseTime=true"
)

var Db *sql.DB

func TestDb(t *testing.T) {
	log.Default().Fatalln("hahaha")
	Db = NewDB() // 此时不会创建链接
	Db.Ping()    // 链接测试
	defer CloseDB()
}

func NewDB() *sql.DB {
	DB, _ := sql.Open("mysql", dburl)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//验证连接
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return nil
	}
	fmt.Println("connnect success")
	return DB
}

func CloseDB() {
	Db.Close()
}
