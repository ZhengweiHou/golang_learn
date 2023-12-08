package main

import (
	"fmt"
	"log"
	"testing"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dburl = "root:root@tcp(127.0.0.1:3306)/dbsynctest?parseTime=true"
)

var Db *sql.DB

func TestDb(t *testing.T) {
	log.Default().Fatalln("hahaha")
	Db = NewDB() // 此时不会创建链接
	Db.Ping()    // 链接测试
	defer CloseDB()
}

func TestAffect(t *testing.T) {
	db, _ := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?parseTime=true")

	// insert one
	result, err := db.Exec("INSERT INTO test.test (id, name, age) values (99, '99', 99)", make([]any, 0)...)
	if err != nil {
		t.Fatal(err)
	}
	aft1, err := result.RowsAffected()
	fmt.Println(aft1)

	// insert multiple
	r2, err := db.Exec("INSERT INTO test.test (id, name, age) values (0, '0', 0),(1, '1', 1),(2, '2', 2),(3, '3', 3),(4, '4', 4)", make([]any, 0)...)
	if err != nil {
		t.Fatal(err)
	}
	aft2, err := r2.RowsAffected()
	fmt.Println(aft2)

	// delete
	r3, err := db.Exec("DELETE FROM test.test WHERE id<?", 100)
	if err != nil {
		t.Fatal(err)
	}
	aft3, err := r3.RowsAffected()
	fmt.Println(aft3)

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
