package test

import (
	"fmt"
	"gormdemo/benchmark"
	"testing"

	_ "github.com/ibmdb/go_ibm_db"
)

// var concrud = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
var concrud = "root:root@tcp(localhost:3306)/test?parseTime=True"

// var dbtype = "db2"
var dbtype = "mysql"

var StuID = "22"

// 增
func TestStudentCreate(t *testing.T) {
	db := benchmark.GetCrudDb(concrud, dbtype)

	student := benchmark.Student{
		Name: "sirius",
		Age:  20,
	}

	result := db.Create(&student)
	if result.Error != nil {
		t.Fatalf("Create failed: %v", result.Error)
	}
	fmt.Printf("Affected: %v Created student: %v\n", result.RowsAffected, student)

}

// 查
func TestStudentRead(t *testing.T) {
	db := benchmark.GetCrudDb(concrud, dbtype)

	var student benchmark.Student
	result := db.First(&student, "ID = ?", StuID)
	if result.Error != nil {
		t.Fatalf("Read failed: %v", result.Error)
	}
	fmt.Printf("Read student: %+v\n", student)
}

// 改
func TestStudentUpdate(t *testing.T) {
	db := benchmark.GetCrudDb(concrud, dbtype)

	var student benchmark.Student
	db.First(&student, "ID = ?", StuID)

	student.Age = student.Age + 1
	result := db.Save(&student)
	if result.Error != nil {
		t.Fatalf("Update failed: %v", result.Error)
	}
}

// 删
func TestStudentDelete(t *testing.T) {
	db := benchmark.GetCrudDb(concrud, dbtype)

	var student benchmark.Student
	db.First(&student, "ID = ?", StuID)

	result := db.Delete(&student)
	if result.Error != nil {
		t.Fatalf("Delete failed: %v", result.Error)
	}

}
