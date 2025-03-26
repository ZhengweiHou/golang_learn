package main

import (
	"context"
	"entdemo/hzwent"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEntMysql(t *testing.T) {
	ctx := context.Background()
	client, err := hzwent.Open("mysql", "root:root@tcp(localhost:3306)/test?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// trancaction
	tx, err := client.Tx(ctx)
	tx.OnCommit(func(c hzwent.Committer) hzwent.Committer {
		fmt.Println("======tx OnCommit hook======")
		return c
	})
	// insert stu with tx
	u, err := tx.Student.
		Create().
		SetName("张三").
		SetStuNo("001").
		//		SetStuno("001").
		SetAge(18).
		Save(ctx)
	tx.Commit()
	fmt.Println(u)

	// create stu
	CreateStudent(context.Background(), client)

	// query
	stu, err := QueryStudent(context.Background(), client)
	fmt.Printf("stu:%v\n", stu)

}

func TestEntSchemaMysql(t *testing.T) {
	client, err := hzwent.Open("mysql", "root:root@tcp(localhost:3306)/test?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
