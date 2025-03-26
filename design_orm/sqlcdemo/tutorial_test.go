package tutorial

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"sqlc/tutorial"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestTororialQuery(t *testing.T) {
	ctx := context.Background()

	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test?parseTime=True")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalf("err %v", err)
	}

	// new queries
	queries := tutorial.New(db)

	// insert
	cstup := tutorial.CreateStudentParams{
		StudentNo: "11",
		Name:      "李四",
		Age:       28,
	}
	queries.CreateStudent(ctx, cstup)

	// query by id
	stu, err := queries.GetStudentById(ctx, 15)
	fmt.Println(stu)

	queries.GetStudentLikeNo(ctx, "%11%")

}

//func TestTororialQueryB(t *testing.B) {
//	ctx := context.Background()
//
//	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test?parseTime=True")
//	//db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
//	if err != nil {
//		log.Fatalf("err %v", err)
//	}
//
//	// new queries
//	queries := tutorial.New(db)
//
//	// query by id
//	stu, err := queries.GetStudentById(ctx, 7)
//	fmt.Println(stu)
//
//}
