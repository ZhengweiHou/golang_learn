package gormdemo

import (
	"fmt"
	"log"
	"sync"
	"testing"

	_ "github.com/ibmdb/go_ibm_db"
	"gorm.io/driver/ibmdb"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var con = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"

func getDb() *gorm.DB {
	dialector := ibmdb.Open(con)

	conf := &gorm.Config{
		PrepareStmt: false,
		//		DryRun:      true, // 生成SQL而不执行
		Logger: logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Warn}),
	}

	db, err := gorm.Open(dialector, conf)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}

	return db

}

func TestQuery(t *testing.T) {
	db := getDb()

	// find all students
	//var allStudents []Student
	//result := db.Find(&allStudents)
	//if result.Error != nil {
	//	log.Fatalf("find err:%v", result.Error)
	//}
	//for _, student := range allStudents {
	//	log.Println(student)
	//}

	// first student
	//stu := Student{}
	//db.First(&stu)
	//log.Println(stu)

	//stu2 := Student{}
	//db.First(&stu2)
	//log.Println(stu2)

	stus := []Student{}
	db.Limit(3).Find(&stus)
	log.Println(stus)
}

func TestInsert(t *testing.T) {
	db := getDb()
	// insert a student
	student := Student{
		Name: "Alice",
		Age:  18,
	}

	result := db.Create(&student)
	if result.Error != nil {
		log.Fatalf("insert err:%v", result.Error)
	}
	log.Println(student)
}

func TestConcurrentInsert(t *testing.T) {
	db := getDb()

	// 并发数量
	concurrency := 10
	var swg, wg sync.WaitGroup
	swg.Add(1)
	wg.Add(concurrency)

	// 用于存储生成的主键ID
	idChan := make(chan *Student, concurrency)

	// 并发插入
	for i := 0; i < concurrency; i++ {
		go func(i int) {
			defer wg.Done()

			student := Student{
				Name: fmt.Sprintf("Student-%d", i),
				Age:  18 + i,
			}
			swg.Wait()
			result := db.Create(&student)
			if result.Error != nil {
				t.Errorf("insert err: %v", result.Error)
				return
			}

			idChan <- &student
		}(i)
	}

	swg.Done()
	// 等待所有goroutine完成
	wg.Wait()
	close(idChan)

	for stu := range idChan {
		id := stu.ID
		stu2 := Student{}
		db.First(&stu2, id)
		if stu.Name != stu2.Name {
			t.Errorf("name not equal, name1:%v,name2:%v", stu.Name, stu2.Name)
		}
		log.Printf("id1:%v,id2:%v,name1:%v,name2:%v", stu.ID, stu2.ID, stu.Name, stu2.Name)
	}

	// 删除测试数据 S00开头的学号
	db.Where("stu_no like ?", "S%").Delete(&Student{})

}

func TestPrepareStmt(t *testing.T) {
	db, err := gorm.Open(
		ibmdb.Open(con),
		&gorm.Config{
			PrepareStmt: true,
			//			DryRun:      true,
		},
	)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}

	tx := db.Session(&gorm.Session{PrepareStmt: true})
	stu := Student{
		Name: "hzw",
		Age:  18,
	}
	tx.First(&stu, 1)
	tx.Find(&stu)
	//	tx.Model(&stu).Update("AGE", 20)

	stmtManger, ok := tx.ConnPool.(*gorm.PreparedStmtDB)
	if !ok {
		log.Fatalf("not prepared stmt db")
	}

	stmt := stmtManger.Stmts
	for sql, stmt := range stmt {
		log.Printf("sql:%v,stmt:%v", sql, stmt)
		stmt.Close()
	}
	stmtManger.Close()

}
