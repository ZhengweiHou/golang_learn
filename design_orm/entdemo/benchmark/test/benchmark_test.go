package test

import (
	"context"
	"entdemo/benchmark"
	"entdemo/benchmark/hzwent"
	"entdemo/benchmark/hzwent/student"
	"fmt"
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
func BenchmarkEntCreate(b *testing.B) {
	client := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer client.Close()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Student.
			Create().
			SetName(fmt.Sprintf("benchmark_%d", i)).
			SetAge(20 + i).
			Save(ctx)
		if err != nil {
			b.Fatalf("Create failed: %v", err)
		}
	}
}

// 基准测试：读取操作
func BenchmarkEntRead(b *testing.B) {
	client := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer client.Close()
	ctx := context.Background()

	// 先创建测试数据
	var testStudents []int
	for i := 0; i < 100; i++ {
		stu, err := client.Student.
			Create().
			SetName(fmt.Sprintf("test_data_%d", i)).
			SetAge(20 + i).
			Save(ctx)
		if err != nil {
			b.Fatalf("Create test data failed: %v", err)
		}
		testStudents = append(testStudents, stu.ID)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		studentID := testStudents[i%len(testStudents)]
		_, err := client.Student.Query().Where(student.IDEQ(studentID)).All(ctx)
		if err != nil {
			b.Fatalf("Read failed: %v", err)
		}
	}
}

// 基准测试：更新操作
func BenchmarkEntUpdate(b *testing.B) {
	client := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer client.Close()
	ctx := context.Background()

	// 先创建测试数据
	var testStudents []int
	for i := 0; i < 100; i++ {
		stu, err := client.Student.
			Create().
			SetName(fmt.Sprintf("test_data_%d", i)).
			SetAge(20 + i).
			Save(ctx)
		if err != nil {
			b.Fatalf("Create test data failed: %v", err)
		}
		testStudents = append(testStudents, stu.ID)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		studentID := testStudents[i%len(testStudents)]
		stus, err := client.Student.Query().Where(student.IDEQ(studentID)).All(ctx)
		if err != nil {
			b.Fatalf("Read failed: %v", err)
		}
		if len(stus) == 0 {
			b.Fatalf("No student found")
		}
		stu := stus[0]
		_, err = stu.Update().
			SetAge(stu.Age + 1).
			Save(ctx)
		if err != nil {
			b.Fatalf("Update failed: %v", err)
		}
	}
}

// 基准测试：删除操作
func BenchmarkEntDelete(b *testing.B) {
	client := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer client.Close()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		stu, err := client.Student.
			Create().
			SetName(fmt.Sprintf("delete_benchmark_%d", i)).
			SetAge(20 + i).
			Save(ctx)
		if err != nil {
			b.Fatalf("Create failed: %v", err)
		}
		err = client.Student.DeleteOneID(stu.ID).Exec(ctx)
		if err != nil {
			b.Fatalf("Delete failed: %v", err)
		}
	}
}

// 基准测试：批量创建
func BenchmarkEntBatchCreate(b *testing.B) {
	client := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer client.Close()
	ctx := context.Background()

	batchSize := 100
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		builders := make([]*hzwent.StudentCreate, batchSize)
		for j := 0; j < batchSize; j++ {
			builders[j] = client.Student.
				Create().
				SetName(fmt.Sprintf("batch_%d_%d", i, j)).
				SetAge(20 + j)
		}
		_, err := client.Student.CreateBulk(builders...).Save(ctx)
		if err != nil {
			b.Fatalf("Batch create failed: %v", err)
		}
	}
}

// 性能对比测试：MySQL vs DB2
func BenchmarkEntMySQLvsDB2(b *testing.B) {
	// MySQL 基准测试
	b.Run("MySQL", func(b *testing.B) {
		client := benchmark.GetCrudDb(mysqlDSN, "mysql")
		defer client.Close()
		ctx := context.Background()

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_, err := client.Student.
				Create().
				SetName(fmt.Sprintf("mysql_benchmark_%d", i)).
				SetAge(20 + i).
				Save(ctx)
			if err != nil {
				b.Fatalf("MySQL create failed: %v", err)
			}
		}
	})

	// DB2 基准测试（需要 DB2 数据库可用）
	// b.Run("DB2", func(b *testing.B) {
	// 	client := benchmark.GetCrudDb(db2DSN, "db2")
	// 	defer client.Close()
	// 	ctx := context.Background()

	// 	b.ResetTimer()
	// 	for i := 0; i < b.N; i++ {
	// 		_, err := client.Student.
	// 			Create().
	// 			SetName(fmt.Sprintf("db2_benchmark_%d", i)).
	// 			SetAge(20 + i).
	// 			Save(ctx)
	// 		if err != nil {
	// 			b.Fatalf("DB2 create failed: %v", err)
	// 		}
	// 	}
	// })
}

// 并发基准测试
func BenchmarkEntConcurrentCreate(b *testing.B) {
	client := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer client.Close()
	ctx := context.Background()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		counter := 0
		for pb.Next() {
			_, err := client.Student.
				Create().
				SetName(fmt.Sprintf("concurrent_%d_%d", time.Now().UnixNano(), counter)).
				SetAge(20 + counter).
				Save(ctx)
			if err != nil {
				b.Fatalf("Concurrent create failed: %v", err)
			}
			counter++
		}
	})
}

// 内存使用基准测试
func BenchmarkEntMemoryUsage(b *testing.B) {
	client := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer client.Close()
	ctx := context.Background()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := client.Student.
			Create().
			SetName(fmt.Sprintf("memory_benchmark_%d", i)).
			SetAge(20 + i).
			Save(ctx)
		if err != nil {
			b.Fatalf("Create failed: %v", err)
		}
	}
}

// Ent 性能基准测试总结
func BenchmarkEntPerformanceSummary(b *testing.B) {
	// 这个测试用于展示 Ent 的整体性能特征
	client := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer client.Close()
	ctx := context.Background()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 创建
		stu, err := client.Student.
			Create().
			SetName(fmt.Sprintf("summary_benchmark_%d", i)).
			SetAge(20 + i).
			Save(ctx)
		if err != nil {
			b.Fatalf("Create failed: %v", err)
		}

		// 读取
		_, err = client.Student.Query().Where(student.IDEQ(stu.ID)).All(ctx)
		if err != nil {
			b.Fatalf("Read failed: %v", err)
		}

		// 更新
		_, err = stu.Update().
			SetAge(stu.Age + 1).
			Save(ctx)
		if err != nil {
			b.Fatalf("Update failed: %v", err)
		}

		// 删除
		err = client.Student.DeleteOneID(stu.ID).Exec(ctx)
		if err != nil {
			b.Fatalf("Delete failed: %v", err)
		}
	}
}

// 复杂查询基准测试
func BenchmarkEntComplexQuery(b *testing.B) {
	client := benchmark.GetCrudDb(mysqlDSN, "mysql")
	defer client.Close()
	ctx := context.Background()

	// 创建测试数据
	for i := 0; i < 1000; i++ {
		_, err := client.Student.
			Create().
			SetName(fmt.Sprintf("complex_test_%d", i)).
			SetAge(20 + i%50).
			Save(ctx)
		if err != nil {
			b.Fatalf("Create test data failed: %v", err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 复杂查询：年龄在 20-40 之间的学生，按年龄降序排列
		_, err := client.Student.
			Query().
			Where(
				student.AgeGTE(20),
				student.AgeLTE(40),
			).
			Order(student.ByAge()).
			All(ctx)
		if err != nil {
			b.Fatalf("Complex query failed: %v", err)
		}
	}
}
