package main

import (
	"fmt"
	"log"
	"testing"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dburl = "root:root@tcp(127.0.0.1:3306)/test?parseTime=true"
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

func TestInsertZ(t *testing.T) {
	db := NewDB()
	defer db.Close()

	// 测试插入单条数据
	result, err := db.Exec("INSERT INTO z (name1, name2, name3, age, notnull, DATE) VALUES (?, ?, ?, ?, ?, ?)",
		"测试1", "test1", []byte("binary1"), 20, "notnull", time.Now())
	if err != nil {
		t.Fatal(err)
	}
	id, _ := result.LastInsertId()
	t.Logf("Inserted record with id: %d", id)

	// 测试插入多条数据
	result, err = db.Exec("INSERT INTO z (name1, name2, name3, age, notnull, DATE) VALUES (?, ?, ?, ?, ?, ?), (?, ?, ?, ?, ?, ?)",
		"测试2", "test2", []byte("binary2"), 21, "notnull", time.Now(),
		"测试3", "test3", []byte("binary3"), 22, "notnull", time.Now())
	if err != nil {
		t.Fatal(err)
	}
	rows, _ := result.RowsAffected()
	t.Logf("Inserted %d records", rows)
}

func TestSelectZ(t *testing.T) {
	db := NewDB()
	defer db.Close()

	// 测试查询单条记录
	var id int
	var name1, name2, name3Str string
	var name3 []byte
	var age int
	var notnull string
	var date time.Time

	err := db.QueryRow("SELECT id, name1, name2, name3, age, notnull, DATE FROM z WHERE id = ?", 1).Scan(
		&id, &name1, &name2, &name3Str, &age, &notnull, &date)
	if err != nil {
		t.Fatal(err)
	}
	// name3Str = string(name3) // 将字节数组转换为UTF-8字符串
	fmt.Printf("name3: %s\n", name3Str)
	t.Logf("Selected record: id=%d, name1=%s, name2=%s, name3=%s, age=%d, notnull=%s, date=%v",
		id, name1, name2, name3Str, age, notnull, date)

	// 测试查询多条记录
	rows, err := db.Query("SELECT id, name1, name2, name3, age, notnull, DATE FROM z WHERE age > ?", 20)
	if err != nil {
		t.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name1, &name2, &name3, &age, &notnull, &date)
		if err != nil {
			t.Fatal(err)
		}
		name3Str := string(name3) // 将字节数组转换为UTF-8字符串
		t.Logf("Selected record: id=%d, name1=%s, name2=%s, name3=%s, age=%d",
			id, name1, name2, name3Str, age)
	}
}

func TestUpdateZ(t *testing.T) {
	db := NewDB()
	defer db.Close()

	// 测试更新数据
	result, err := db.Exec("UPDATE z SET age = ?, name1 = ? WHERE id = ?",
		30, "更新后的名字", 1)
	if err != nil {
		t.Fatal(err)
	}
	rows, _ := result.RowsAffected()
	t.Logf("Updated %d records", rows)
}

func TestDeleteZ(t *testing.T) {
	db := NewDB()
	defer db.Close()

	// 测试删除数据
	result, err := db.Exec("DELETE FROM z WHERE id = ?", 1)
	if err != nil {
		t.Fatal(err)
	}
	rows, _ := result.RowsAffected()
	t.Logf("Deleted %d records", rows)
}

/*
CREATE TABLE `z` (
	`id` int NOT NULL AUTO_INCREMENT,
	`name1` varchar(20) COLLATE DEFAULT NULL,
	`name2` varchar(20) CHARACTER SET latin1 COLLATE latin1_swedish_ci DEFAULT NULL,
	`name3` varbinary(20) DEFAULT NULL,
	`age` int DEFAULT NULL,
	`notnull` varchar(9) CHARACTER SET utf8mb4 COLLATE NOT NULL,
	`DATE` datetime DEFAULT NULL,
	PRIMARY KEY (`id`)
  )
*/
