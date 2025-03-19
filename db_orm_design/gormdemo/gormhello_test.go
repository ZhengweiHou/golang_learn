package gormdemo

import (
	"fmt"
	"log"
	"testing"

	_ "github.com/ibmdb/go_ibm_db"
	"gorm.io/driver/ibmdb"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func TestGormBasehello(t *testing.T) {
	dialector := mysql.Open("root:root@tcp(localhost:3306)/test?parseTime=True")
	dialector = sqlite.Open("test.db")
	dialector = sqlserver.Open("sqlserver://gorm:LoremIpsum86@localhost:9930?database=sqlserver")
	dialector = postgres.Open("user=gorm password=gorm DB.name=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai")

	fmt.Printf("%v", dialector)
}

func TestGormhello(t *testing.T) {
	con := "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	dialector := ibmdb.Open(con)

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatalf("open err: %v", err)
	}

	// find all students
	var allStudents []Student
	result := db.Find(&allStudents)
	if result.Error != nil {
		log.Fatalf("find err:%v", result.Error)
	}
	for _, student := range allStudents {
		log.Println(student)
	}
}

func TestGormhello2(t *testing.T) {
	dialector := mysql.Open("root:root@tcp(localhost:3306)/test?parseTime=True")

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// find all students
	var allStudents []Student
	result := db.Find(&allStudents)
	if result.Error != nil {
		log.Fatalf("find err:%v", result.Error)
	}
	for _, student := range allStudents {
		log.Println(student)
	}

}
