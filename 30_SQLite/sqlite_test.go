package main

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func Test1(t *testing.T) {
	db, err := sql.Open("sqlite3", "sqlitedb.db")
	if err != nil {
		// 处理连接错误
	}     
	defer db.Close()

	// 建表
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER
	)
	`)
	if err != nil {
		// 处理建表错误
	}

	// 插入
	_, err = db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 25)
	if err != nil {
		// 处理插入数据错误
	}

	// 查询
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		// 处理查询错误
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		err := rows.Scan(&id, &name, &age)
		fmt.Printf("id:%v,name:%v,age:%v\n", id, name, age)
		if err != nil {
			// 处理扫描错误
		}
		// 处理每行数据
	}

	if rows.Err() != nil {
		// 处理迭代错误
	}

	// 更新
	_, err = db.Exec("UPDATE users SET age = ? WHERE name = ?", 30, "Alice")
	if err != nil {
		// 处理更新数据错误
	}

	// 删除
	// _, err = db.Exec("DELETE FROM users WHERE id > ?", 0)
	// if err != nil {
	// 	// 处理删除数据错误
	// }

}
