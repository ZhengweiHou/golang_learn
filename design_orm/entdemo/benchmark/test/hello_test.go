package test

import (
	"context"
	"entdemo/benchmark"
	"entdemo/benchmark/hzwent/student"
	"log"
	"testing"

	_ "github.com/ibmdb/go_ibm_db"
)

// var concrud = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
var concrud = "root:root@tcp(localhost:3306)/test?parseTime=True"

// var dbtype = "db2"
var dbtype = "mysql"
var StuID = 64

// 增
func TestStudentCreate(t *testing.T) {
	cli := benchmark.GetCrudDb(concrud, dbtype)
	defer cli.Close()

	ctx := context.Background()
	// create one
	istu, err := cli.Student.
		Create().
		SetName("sirius").
		SetAge(18).
		Save(ctx) // SaveX(ctx) 会panic
	if err != nil {
		log.Fatalf("failed creating user: %v", err)
	}
	log.Printf("student create one: id:%v  %v ", istu.ID, istu)
}

// 查
func TestStudentRead(t *testing.T) {
	cli := benchmark.GetCrudDb(concrud, dbtype)
	defer cli.Close()

	ctx := context.Background()
	stu, err := cli.Student.Query().Where(student.IDEQ(StuID)).First(ctx)
	if err != nil {
		log.Fatalf("failed reading user: %v", err)
	}
	log.Printf("student read one: id:%v  %v ", stu.ID, stu)

}

// 改
func TestStudentUpdate(t *testing.T) {
	cli := benchmark.GetCrudDb(concrud, dbtype)
	defer cli.Close()

	ctx := context.Background()
	stu, err := cli.Student.Query().Where(student.IDEQ(StuID)).First(ctx)
	if err != nil {
		log.Fatalf("failed reading user: %v", err)
	}
	log.Printf("student read one: id:%v  %v ", stu.ID, stu)

	// update one
	// SQL: UPDATE student SET NAME = ? WHERE ID = ?
	ustu, err := stu.Update().
		SetAge(stu.Age + 1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed updating user: %v", err)
	}
	log.Printf("2: student update one: %v", ustu)
}

// 删
func TestStudentDelete(t *testing.T) {
	cli := benchmark.GetCrudDb(concrud, dbtype)
	defer cli.Close()

	ctx := context.Background()
	stu, err := cli.Student.Query().Where(student.IDEQ(StuID)).First(ctx)
	if err != nil {
		log.Fatalf("failed reading user: %v", err)
	}
	log.Printf("student read one: id:%v  %v ", stu.ID, stu)

	// delete one
	// SQL: DELETE FROM student WHERE ID = ?
	err = cli.Student.DeleteOneID(stu.ID).Exec(ctx)
	if err != nil {
		log.Fatalf("failed deleting user: %v", err)
	}
	log.Printf("3: student delete one: %v", stu)
}
