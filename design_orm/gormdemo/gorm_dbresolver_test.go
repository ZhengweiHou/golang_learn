package gormdemo

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

var (
	con1 = "root:root@tcp(localhost:3306)/test?parseTime=True"
	con2 = "root:root@tcp(localhost:3306)/test2?parseTime=True"
)

func getResolverDb() *gorm.DB {
	conf := &gorm.Config{
		// Logger: logger.New(log.New(log.Writer(), "\r\n", log.LstdFlags), logger.Config{LogLevel: logger.Info}),
	}

	db, err := gorm.Open(
		mysql.Open(con1),
		conf,
	)
	if err != nil {
		log.Fatalf("open err: %v", err)
	}

	resolver := dbresolver.Register(
		dbresolver.Config{
			Sources: []gorm.Dialector{
				mysql.Open(con1),
				mysql.Open(con2),
			},
			Replicas: []gorm.Dialector{
				mysql.Open(con2),
			},
			Policy:            dbresolver.RandomPolicy{},
			TraceResolverMode: true,
		},
	).Register(
		dbresolver.Config{
			Sources: []gorm.Dialector{
				mysql.Open(con2),
			},
			Replicas: []gorm.Dialector{
				mysql.Open(con1),
				mysql.Open(con2),
			},
		},
		"secondary", // 定义为后线库
	).Register(
		dbresolver.Config{
			Sources: []gorm.Dialector{
				mysql.Open(con2),
			},
		},
		&Student2{}, // 指定表访问con2
	)

	db.Use(
		resolver,
	)
	return db
}

func TestDbResolverCreate(t *testing.T) {
	db := getResolverDb()
	for i := 0; i < 10; i++ {
		stu := Student{Name: fmt.Sprintf("default_%d", i)}
		db.Create(&stu) // 分别落入con1, con2

		stu = Student{Name: fmt.Sprintf("write_%d", i)}
		db.Clauses(dbresolver.Write).Create(&stu) // Write 使用source， 分别落入con1, con2

		stu = Student{Name: fmt.Sprintf("read_%d", i)}
		db.Clauses(dbresolver.Read).Create(&stu) // Read 使用replicas， 预期落入con2，实际落入con1

		stu2 := Student2{Name: fmt.Sprintf("STU2_%d", i)}
		db.Create(&stu2) // student2 表 指定落入con2

		// 指定后线库操作
		stu = Student{Name: fmt.Sprintf("SECONE_STU_%d", i)}
		db.Clauses(dbresolver.Use("secondary")).Create(&stu) // 指定后线库操作， 落入con2

	}
}

func TestDbResolverQuery(t *testing.T) {
	db := getResolverDb()
	for i := 0; i < 5; i++ {
		stu2 := Student2{}
		db.Clauses(dbresolver.Use("secondary")).First(&stu2, 1) // 查询默认使用read Replicas
		fmt.Println("secondary:", stu2)
	}
	fmt.Println("==========")

	for i := 0; i < 5; i++ {
		stu2 := Student2{}
		db.Clauses(dbresolver.Use("secondary"), dbresolver.Write).First(&stu2, 1) // 后线库查询
		fmt.Println("secondary write:", stu2)
	}
	fmt.Println("==========")

	for i := 0; i < 5; i++ {
		stu2 := Student2{}
		db.Clauses(dbresolver.Use("secondary"), dbresolver.Read).First(&stu2, 1) // 后线库查询
		fmt.Println("secondary read:", stu2)
	}
}
