package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("12313")
	// var db *sqlx.DB
	// //db, err := sqlx.Open("mysql", "username:password@tcp(ip:port)/database?charset=utf8")
	// db, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/dbsynctest?parseTime=true")
	// fmt.Println(db, err)
}
