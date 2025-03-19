package gormdemo

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getMysqlDb() *gorm.DB {
	con := "root:root@tcp(localhost:3306)/test?parseTime=True"
	dialector := mysql.Open(con)
	conf := &gorm.Config{
		PrepareStmt: true,
		//		DryRun:      true, // 生成SQL而不执行
		Logger: logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Info}),
	}

	db, err := gorm.Open(dialector, conf)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}
	return db

}

func TestMysqlQuery(t *testing.T) {
	db := getMysqlDb()

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

func TestMysqlInsert(t *testing.T) {
	db := getMysqlDb()

	// insert a student
	//	student := Student{
	//		StuNo: "111",
	//		Name:  "Alice",
	//		Age:   18,
	//	}
	//	result := db.Create(&student)
	//	if result.Error != nil {
	//		log.Fatalf("insert err:%v", result.Error)
	//	}
	//	log.Println(student)

	// insert a teacher
	teacher := Teacher{
		Name: "Tom3",
		Age:  30,
	}

	result := db.Create(&teacher)
	if result.Error != nil {
		log.Fatalf("insert err:%v", result.Error)
	}
	log.Println(teacher)

}
