package gormdemo

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/ibmdb"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var concrud = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"

func getCrudDb() *gorm.DB {
	dialector := ibmdb.Open(concrud)

	conf := &gorm.Config{
		PrepareStmt: false,
		//		DryRun:      true, // 生成SQL而不执行
		Logger: logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Info}),
	}

	db, err := gorm.Open(dialector, conf)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}

	return db
}

func TestStudentCreate(t *testing.T) {
	db := getCrudDb()

	student := Student{
		Name: "sirius",
		Age:  20,
	}

	result := db.Create(&student)
	if result.Error != nil {
		t.Fatalf("Create failed: %v", result.Error)
	}
	fmt.Printf("Affected: %v Created student: %v\n", result.RowsAffected, student)

	hzw := Hzw{
		Name: "hzw",
		Age:  20,
	}
	db.Create(&hzw)
}

func TestStudentRead(t *testing.T) {
	db := getCrudDb()

	var student Student
	result := db.First(&student, "STU_NO = ?", "2023001")
	if result.Error != nil {
		t.Fatalf("Read failed: %v", result.Error)
	}
	fmt.Printf("Read student: %+v\n", student)
}

func TestStudentUpdate(t *testing.T) {
	db := getCrudDb()

	var student Student
	db.First(&student, "STU_NO = ?", "2023001")

	student.Age = 21
	result := db.Save(&student)
	if result.Error != nil {
		t.Fatalf("Update failed: %v", result.Error)
	}

	var updatedStudent Student
	db.First(&updatedStudent, student.ID)
	if updatedStudent.Age != 21 {
		t.Errorf("Expected age 21, got %d", updatedStudent.Age)
	}
	fmt.Printf("Updated student: %+v\n", updatedStudent)
}

func TestStudentDelete(t *testing.T) {
	db := getCrudDb()

	var student Student
	db.First(&student, "STU_NO = ?", "2023001")

	result := db.Delete(&student)
	if result.Error != nil {
		t.Fatalf("Delete failed: %v", result.Error)
	}

	var deletedStudent Student
	result = db.First(&deletedStudent, student.ID)
	if result.Error == nil {
		t.Error("Expected record to be deleted")
	}
	fmt.Printf("Deleted student ID: %d\n", student.ID)
}

func TestStudentOptimisticLocking(t *testing.T) {
	db := getCrudDb()

	// 创建初始记录
	student := Student{
		Name: "李四",
		Age:  22,
	}
	db.Create(&student)

	// 获取第一个版本
	var student1 Student
	db.First(&student1, student.ID)
	version1 := student1.Version

	// 获取第二个版本
	var student2 Student
	db.First(&student2, student.ID)
	version2 := student2.Version

	// 确保两个版本相同
	if version1 != version2 {
		t.Errorf("Expected versions to match, got %d and %d", version1, version2)
	}

	// 更新第一个实例
	student1.Age = 23
	db.Save(&student1)

	// 尝试更新第二个实例 - 应该失败
	student2.Age = 24
	result := db.Save(&student2)
	if result.Error == nil {
		t.Error("Expected optimistic lock error but got none")
	} else {
		fmt.Printf("Optimistic lock error: %v\n", result.Error)
	}

	// 清理
	db.Delete(&student)
}
func TestStuUpsert(t *testing.T) {
	db := getMysqlDb()
	//	db := getHzwDb()

	stus := []*Student{
		{Name: "1111", Age: 1111},
		{Name: "2222", Age: 2222},
	}
	tx := db.Save(stus) // save列表时为merge操作
	if tx.Error != nil {
		log.Fatalf("save err:%v", tx.Error)
	}

	for _, hzw := range stus {
		jstr, err := json.Marshal(hzw)
		if err != nil {
			log.Fatalf("json marshal failed: %v", err)
		}
		fmt.Printf("save: %s\n", jstr)
	}

}
