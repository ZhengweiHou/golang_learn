package test

import (
	"fmt"
	"gormdemo/benchmark"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ibmdb/go_ibm_db"
)

var (
	// 数据库配置
	db2DSN   = "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST"
	mysqlDSN = "root:root@tcp(localhost:3306)/test?parseTime=True"
)

// 基准测试：创建操作
func BenchmarkGormCreate(b *testing.B) {
	db := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		student := benchmark.Student{
			Name: fmt.Sprintf("benchmark_%d", i),
			Age:  20 + i,
		}
		if err := db.Create(&student).Error; err != nil {
			b.Fatalf("Create failed: %v", err)
		}
	}
}

// 基准测试：读取操作
func BenchmarkGormRead(b *testing.B) {
	db := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 先创建测试数据
	var testStudents []benchmark.Student
	for i := 0; i < 100; i++ {
		student := benchmark.Student{
			Name: fmt.Sprintf("test_data_%d", i),
			Age:  20 + i,
		}
		if err := db.Create(&student).Error; err != nil {
			b.Fatalf("Create test data failed: %v", err)
		}
		testStudents = append(testStudents, student)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var student benchmark.Student
		studentID := testStudents[i%len(testStudents)].ID
		if err := db.First(&student, studentID).Error; err != nil {
			b.Fatalf("Read failed: %v", err)
		}
	}
}

// 基准测试：更新操作
func BenchmarkGormUpdate(b *testing.B) {
	db := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 先创建测试数据
	var testStudents []benchmark.Student
	for i := 0; i < 100; i++ {
		student := benchmark.Student{
			Name: fmt.Sprintf("test_data_%d", i),
			Age:  20 + i,
		}
		if err := db.Create(&student).Error; err != nil {
			b.Fatalf("Create test data failed: %v", err)
		}
		testStudents = append(testStudents, student)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var student benchmark.Student
		studentID := testStudents[i%len(testStudents)].ID
		if err := db.First(&student, studentID).Error; err != nil {
			b.Fatalf("Read failed: %v", err)
		}
		student.Age = student.Age + 1
		if err := db.Save(&student).Error; err != nil {
			b.Fatalf("Update failed: %v", err)
		}
	}
}

// 基准测试：删除操作
func BenchmarkGormDelete(b *testing.B) {
	db := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		student := benchmark.Student{
			Name: fmt.Sprintf("delete_benchmark_%d", i),
			Age:  20 + i,
		}
		if err := db.Create(&student).Error; err != nil {
			b.Fatalf("Create failed: %v", err)
		}
		if err := db.Delete(&student).Error; err != nil {
			b.Fatalf("Delete failed: %v", err)
		}
	}
}

// 基准测试：批量创建
func BenchmarkGormBatchCreate(b *testing.B) {
	db := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	batchSize := 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		students := make([]benchmark.Student, batchSize)
		for j := 0; j < batchSize; j++ {
			students[j] = benchmark.Student{
				Name: fmt.Sprintf("batch_%d_%d", i, j),
				Age:  20 + j,
			}
		}
		if err := db.Create(&students).Error; err != nil {
			b.Fatalf("Batch create failed: %v", err)
		}
	}
}

// 性能对比测试：MySQL vs DB2
func BenchmarkGormMySQLvsDB2(b *testing.B) {
	// MySQL 基准测试
	b.Run("MySQL", func(b *testing.B) {
		db := benchmark.GetCrudDb(mysqlDSN, "mysql")
		defer func() {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			student := benchmark.Student{
				Name: fmt.Sprintf("mysql_benchmark_%d", i),
				Age:  20 + i,
			}
			if err := db.Create(&student).Error; err != nil {
				b.Fatalf("MySQL create failed: %v", err)
			}
		}
	})

	// DB2 基准测试（需要 DB2 数据库可用）
	// b.Run("DB2", func(b *testing.B) {
	// 	db := benchmark.GetCrudDb(db2DSN, "db2")
	// 	defer func() {
	// 		sqlDB, _ := db.DB()
	// 		sqlDB.Close()
	// 	}()

	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		student := benchmark.Student{
	// 			Name: fmt.Sprintf("db2_benchmark_%d", i),
	// 			Age:  20 + i,
	// 		}
	// 		if err := db.Create(&student).Error; err != nil {
	// 			b.Fatalf("DB2 create failed: %v", err)
	// 		}
	// 	}
	// })
}

// 并发基准测试
func BenchmarkGormConcurrentCreate(b *testing.B) {
	db := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			student := benchmark.Student{
				Name: fmt.Sprintf("concurrent_%d_%d", time.Now().UnixNano(), counter),
				Age:  20 + counter,
			}
			if err := db.Create(&student).Error; err != nil {
				b.Fatalf("Concurrent create failed: %v", err)
			}
			counter++
		}
	})
}

// 内存使用基准测试
func BenchmarkGormMemoryUsage(b *testing.B) {
	db := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		student := benchmark.Student{
			Name: fmt.Sprintf("memory_benchmark_%d", i),
			Age:  20 + i,
		}
		if err := db.Create(&student).Error; err != nil {
			b.Fatalf("Create failed: %v", err)
		}
	}
}
