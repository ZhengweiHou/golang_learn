package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"

	"entdemo/benchmark"
	"entdemo/benchmark/hzwent"
	"entdemo/benchmark/hzwent/student"
)

// 性能测试配置
type BenchmarkConfig struct {
	Threads        int           `json:"threads"`         // 并发线程数
	Iterations     int           `json:"iterations"`      // 每个线程的迭代次数
	WarmupTime     time.Duration `json:"warmup_time"`     // 预热时间
	TestTime       time.Duration `json:"test_time"`       // 测试时间
	BatchSize      int           `json:"batch_size"`      // 批量操作大小
	ReportInterval time.Duration `json:"report_interval"` // 报告间隔
}

// 性能测试结果
type BenchmarkResult struct {
	TestName      string  `json:"test_name"`
	TotalOps      int64   `json:"total_ops"`
	TotalTime     float64 `json:"total_time_ms"` // 总时间(毫秒)
	OpsPerSecond  float64 `json:"ops_per_second"`
	AvgLatency    float64 `json:"avg_latency_ms"` // 平均延迟(毫秒)
	MinLatency    float64 `json:"min_latency_ms"` // 最小延迟(毫秒)
	MaxLatency    float64 `json:"max_latency_ms"` // 最大延迟(毫秒)
	P95Latency    float64 `json:"p95_latency_ms"` // P95延迟(毫秒)
	P99Latency    float64 `json:"p99_latency_ms"` // P99延迟(毫秒)
	ErrorCount    int64   `json:"error_count"`
	MemoryUsageMB float64 `json:"memory_usage_mb"`
}

// 测试运行器
type BenchmarkRunner struct {
	config *BenchmarkConfig
	client *hzwent.Client
	ctx    context.Context
}

func NewBenchmarkRunner(config *BenchmarkConfig, dsn string, dbType string) *BenchmarkRunner {
	ctx := context.Background()
	return &BenchmarkRunner{
		config: config,
		client: benchmark.GetCrudDb(dsn, dbType),
		ctx:    ctx,
	}
}

func main() {
	// 解析命令行参数
	threads := flag.Int("threads", 1, "并发线程数")
	iterations := flag.Int("iterations", 1, "每个线程的迭代次数")
	warmup := flag.Duration("warmup", 5*time.Second, "预热时间")
	testTime := flag.Duration("test-time", 5*time.Second, "测试时间")
	batchSize := flag.Int("batch-size", 100, "批量操作大小")
	reportFile := flag.String("report", "benchmark_report.json", "测试报告文件")
	dbType := flag.String("db-type", "db2", "数据库类型: mysql 或 db2")
	dsn := flag.String("dsn", "HOSTNAME=localhost;DATABASE=testdb;PORT=50001;UID=db2inst1;PWD=db2inst1;AUTHENTICATION=SERVER;CurrentSchema=ENTTEST", "数据库连接字符串")
	cpuProfile := flag.String("cpuprofile", "", "CPU profile 输出文件")
	memProfile := flag.String("memprofile", "", "内存 profile 输出文件")
	flag.Parse()

	config := &BenchmarkConfig{
		Threads:        *threads,
		Iterations:     *iterations,
		WarmupTime:     *warmup,
		TestTime:       *testTime,
		BatchSize:      *batchSize,
		ReportInterval: 5 * time.Second,
	}

	// 启动 CPU profiling
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal("无法创建 CPU profile 文件:", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("无法启动 CPU profiling:", err)
		}
		defer pprof.StopCPUProfile()
		fmt.Printf("CPU profiling 已启动，输出到: %s\n", *cpuProfile)
	}

	runner := NewBenchmarkRunner(config, *dsn, *dbType)
	defer runner.client.Close()

	fmt.Printf("开始性能基准测试...\n")
	fmt.Printf("配置: 线程数=%d, 迭代次数=%d, 预热时间=%v, 测试时间=%v, 批量大小=%d\n",
		config.Threads, config.Iterations, config.WarmupTime, config.TestTime, config.BatchSize)

	// 运行预热
	fmt.Printf("预热阶段...\n")
	runner.warmup()

	// 运行所有性能测试
	results := make(map[string]*BenchmarkResult)

	// 创建性能测试
	fmt.Printf("\n运行创建性能测试...\n")
	results["create"] = runner.runCreateBenchmark()

	// 读取性能测试
	fmt.Printf("\n运行读取性能测试...\n")
	results["read"] = runner.runReadBenchmark()

	// 更新性能测试
	fmt.Printf("\n运行更新性能测试...\n")
	results["update"] = runner.runUpdateBenchmark()

	// 删除性能测试
	fmt.Printf("\n运行删除性能测试...\n")
	results["delete"] = runner.runDeleteBenchmark()

	// // 批量操作性能测试
	// fmt.Printf("\n运行批量操作性能测试...\n")
	// results["batch"] = runner.runBatchBenchmark()

	// // 混合工作负载测试
	// fmt.Printf("\n运行混合工作负载测试...\n")
	// results["mixed"] = runner.runMixedWorkloadBenchmark()

	// 生成报告
	fmt.Printf("\n生成测试报告...\n")
	runner.generateReport(results, *reportFile)

	// 输出内存 profile
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal("无法创建内存 profile 文件:", err)
		}
		defer f.Close()
		runtime.GC() // 获取最新的 GC 统计
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("无法写入内存 profile:", err)
		}
		fmt.Printf("内存 profile 已保存到: %s\n", *memProfile)
	}

	fmt.Printf("性能基准测试完成！报告已保存到: %s\n", *reportFile)
}

// 预热阶段
func (r *BenchmarkRunner) warmup() {
	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < r.config.Threads; i++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			for time.Since(start) < r.config.WarmupTime {
				// 执行一些基础操作进行预热
				stu, err := r.client.Student.
					Create().
					SetName(fmt.Sprintf("warmup_%d", threadID)).
					SetAge(20 + threadID).
					Save(r.ctx)
				if err == nil {
					r.client.Student.DeleteOneID(stu.ID).Exec(r.ctx)
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Printf("预热完成\n")
}

// 运行创建性能测试
func (r *BenchmarkRunner) runCreateBenchmark() *BenchmarkResult {
	var statsMutex sync.Mutex
	var latencies []time.Duration
	var errorCount int64
	var totalOps int64

	startTime := time.Now()
	stopTime := startTime.Add(r.config.TestTime)

	var wg sync.WaitGroup
	for i := 0; i < r.config.Threads; i++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			for time.Now().Before(stopTime) {
				opStart := time.Now()

				_, err := r.client.Student.
					Create().
					SetName(fmt.Sprintf("student_%d_%d", threadID, time.Now().UnixNano())).
					SetAge(20 + threadID).
					Save(r.ctx)

				if err != nil {
					statsMutex.Lock()
					errorCount++
					statsMutex.Unlock()
					continue
				}

				latency := time.Since(opStart)

				statsMutex.Lock()
				latencies = append(latencies, latency)
				totalOps++
				statsMutex.Unlock()
			}
		}(i)
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	return r.calculateResult("创建性能测试", latencies, totalOps, errorCount, totalTime)
}

// 运行读取性能测试
func (r *BenchmarkRunner) runReadBenchmark() *BenchmarkResult {
	// 先创建一些测试数据
	testStudents := r.createTestDataSequential(1000)

	var statsMutex sync.Mutex
	var latencies []time.Duration
	var errorCount int64
	var totalOps int64

	startTime := time.Now()
	stopTime := startTime.Add(r.config.TestTime)

	var wg sync.WaitGroup
	for i := 0; i < r.config.Threads; i++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			for time.Now().Before(stopTime) {
				studentID := testStudents[threadID%len(testStudents)]

				opStart := time.Now()
				// _, err := r.client.Student.Query().Where(student.IDEQ(studentID)).First(r.ctx)
				_, err := r.client.Student.Query().Where(student.IDEQ(studentID)).All(r.ctx)
				if err != nil {
					statsMutex.Lock()
					errorCount++
					statsMutex.Unlock()
					continue
				}

				latency := time.Since(opStart)

				statsMutex.Lock()
				latencies = append(latencies, latency)
				totalOps++
				statsMutex.Unlock()
			}
		}(i)
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	return r.calculateResult("读取性能测试", latencies, totalOps, errorCount, totalTime)
}

// 运行更新性能测试
func (r *BenchmarkRunner) runUpdateBenchmark() *BenchmarkResult {
	// 先创建一些测试数据
	testStudents := r.createTestDataSequential(1000)

	var statsMutex sync.Mutex
	var latencies []time.Duration
	var errorCount int64
	var totalOps int64

	startTime := time.Now()
	stopTime := startTime.Add(r.config.TestTime)

	var wg sync.WaitGroup
	for i := 0; i < r.config.Threads; i++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			for time.Now().Before(stopTime) {
				studentID := testStudents[threadID%len(testStudents)]

				opStart := time.Now()
				//stu, err := r.client.Student.Query().Where(student.IDEQ(studentID)).First(r.ctx)
				stus, err := r.client.Student.Query().Where(student.IDEQ(studentID)).All(r.ctx)
				if err != nil {
					statsMutex.Lock()
					errorCount++
					statsMutex.Unlock()
					continue
				}
				if len(stus) == 0 {
					continue
				}

				stu := stus[0]

				_, err = stu.Update().
					SetAge(stu.Age + 1).
					Save(r.ctx)
				if err != nil {
					statsMutex.Lock()
					errorCount++
					statsMutex.Unlock()
					continue
				}

				latency := time.Since(opStart)

				statsMutex.Lock()
				latencies = append(latencies, latency)
				totalOps++
				statsMutex.Unlock()
			}
		}(i)
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	return r.calculateResult("更新性能测试", latencies, totalOps, errorCount, totalTime)
}

// 运行删除性能测试
func (r *BenchmarkRunner) runDeleteBenchmark() *BenchmarkResult {
	var statsMutex sync.Mutex
	var latencies []time.Duration
	var errorCount int64
	var totalOps int64

	startTime := time.Now()
	stopTime := startTime.Add(r.config.TestTime)

	var wg sync.WaitGroup
	for i := 0; i < r.config.Threads; i++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			for time.Now().Before(stopTime) {
				// 先创建再删除
				stu, err := r.client.Student.
					Create().
					SetName(fmt.Sprintf("delete_test_%d_%d", threadID, time.Now().UnixNano())).
					SetAge(20 + threadID).
					Save(r.ctx)

				if err != nil {
					statsMutex.Lock()
					errorCount++
					statsMutex.Unlock()
					continue
				}

				opStart := time.Now()
				err = r.client.Student.DeleteOneID(stu.ID).Exec(r.ctx)
				if err != nil {
					statsMutex.Lock()
					errorCount++
					statsMutex.Unlock()
					continue
				}

				latency := time.Since(opStart)

				statsMutex.Lock()
				latencies = append(latencies, latency)
				totalOps++
				statsMutex.Unlock()
			}
		}(i)
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	return r.calculateResult("删除性能测试", latencies, totalOps, errorCount, totalTime)
}

// 运行批量操作性能测试
func (r *BenchmarkRunner) runBatchBenchmark() *BenchmarkResult {
	var statsMutex sync.Mutex
	var latencies []time.Duration
	var errorCount int64
	var totalOps int64

	startTime := time.Now()
	stopTime := startTime.Add(r.config.TestTime)

	var wg sync.WaitGroup
	for i := 0; i < r.config.Threads; i++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			for time.Now().Before(stopTime) {
				opStart := time.Now()

				// 批量创建
				builders := make([]*hzwent.StudentCreate, r.config.BatchSize)
				for j := 0; j < r.config.BatchSize; j++ {
					builders[j] = r.client.Student.
						Create().
						SetName(fmt.Sprintf("batch_%d_%d_%d", threadID, j, time.Now().UnixNano())).
						SetAge(20 + j)
				}

				_, err := r.client.Student.CreateBulk(builders...).Save(r.ctx)
				if err != nil {
					statsMutex.Lock()
					errorCount++
					statsMutex.Unlock()
					continue
				}

				latency := time.Since(opStart)

				statsMutex.Lock()
				latencies = append(latencies, latency)
				totalOps += int64(r.config.BatchSize)
				statsMutex.Unlock()
			}
		}(i)
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	return r.calculateResult("批量操作性能测试", latencies, totalOps, errorCount, totalTime)
}

// 运行混合工作负载测试
func (r *BenchmarkRunner) runMixedWorkloadBenchmark() *BenchmarkResult {
	testStudents := r.createTestDataSequential(1000)

	var statsMutex sync.Mutex
	var latencies []time.Duration
	var errorCount int64
	var totalOps int64

	startTime := time.Now()
	stopTime := startTime.Add(r.config.TestTime)

	var wg sync.WaitGroup
	for i := 0; i < r.config.Threads; i++ {
		wg.Add(1)
		go func(threadID int) {
			defer wg.Done()
			operation := 0
			for time.Now().Before(stopTime) {
				var latency time.Duration
				var err error

				switch operation % 4 {
				case 0: // 创建
					opStart := time.Now()
					_, err = r.client.Student.
						Create().
						SetName(fmt.Sprintf("mixed_%d_%d", threadID, time.Now().UnixNano())).
						SetAge(20 + threadID).
						Save(r.ctx)
					latency = time.Since(opStart)

				case 1: // 读取
					studentID := testStudents[threadID%len(testStudents)]
					opStart := time.Now()
					_, err = r.client.Student.Query().Where(student.IDEQ(studentID)).First(r.ctx)
					latency = time.Since(opStart)

				case 2: // 更新
					studentID := testStudents[threadID%len(testStudents)]
					opStart := time.Now()
					stu, err1 := r.client.Student.Query().Where(student.IDEQ(studentID)).First(r.ctx)
					if err1 == nil {
						_, err = stu.Update().
							SetAge(stu.Age + 1).
							Save(r.ctx)
					} else {
						err = err1
					}
					latency = time.Since(opStart)

				case 3: // 删除
					stu, err1 := r.client.Student.
						Create().
						SetName(fmt.Sprintf("mixed_delete_%d_%d", threadID, time.Now().UnixNano())).
						SetAge(20 + threadID).
						Save(r.ctx)
					if err1 == nil {
						opStart := time.Now()
						err = r.client.Student.DeleteOneID(stu.ID).Exec(r.ctx)
						latency = time.Since(opStart)
					} else {
						err = err1
					}
				}

				if err != nil {
					statsMutex.Lock()
					errorCount++
					statsMutex.Unlock()
				} else {
					statsMutex.Lock()
					latencies = append(latencies, latency)
					totalOps++
					statsMutex.Unlock()
				}

				operation++
			}
		}(i)
	}

	wg.Wait()
	totalTime := time.Since(startTime)

	return r.calculateResult("混合工作负载测试", latencies, totalOps, errorCount, totalTime)
}

// 创建测试数据（批量方式，可能有问题）
func (r *BenchmarkRunner) createTestData(count int) []int {
	students := make([]*hzwent.StudentCreate, count)
	for i := 0; i < count; i++ {
		students[i] = r.client.Student.
			Create().
			SetName(fmt.Sprintf("test_data_%d", i)).
			SetAge(20 + i)
	}

	createdStudents, err := r.client.Student.CreateBulk(students...).Save(r.ctx)
	if err != nil {
		log.Printf("创建测试数据失败: %v", err)
		log.Panic("创建测试数据失败")
		return []int{}
	}

	studentIDs := make([]int, len(createdStudents))
	for i, stu := range createdStudents {
		studentIDs[i] = stu.ID
	}

	return studentIDs
}

// 创建测试数据（逐个创建方式）
func (r *BenchmarkRunner) createTestDataSequential(count int) []int {
	studentIDs := make([]int, count)

	for i := 0; i < count; i++ {
		stu, err := r.client.Student.
			Create().
			SetName(fmt.Sprintf("test_data_%d", i)).
			SetAge(20 + i).
			Save(r.ctx)

		if err != nil {
			log.Printf("创建测试数据失败: %v", err)
			return []int{}
		}

		studentIDs[i] = stu.ID
	}

	return studentIDs
}

// 计算性能结果
func (r *BenchmarkRunner) calculateResult(testName string, latencies []time.Duration, totalOps int64, errorCount int64, totalTime time.Duration) *BenchmarkResult {
	if len(latencies) == 0 {
		return &BenchmarkResult{
			TestName:   testName,
			ErrorCount: errorCount,
		}
	}

	// 计算延迟统计
	var totalLatency time.Duration
	var minLatency, maxLatency time.Duration = latencies[0], latencies[0]

	for _, latency := range latencies {
		totalLatency += latency
		if latency < minLatency {
			minLatency = latency
		}
		if latency > maxLatency {
			maxLatency = latency
		}
	}

	avgLatency := totalLatency / time.Duration(len(latencies))

	// 计算百分位延迟
	sortedLatencies := make([]time.Duration, len(latencies))
	copy(sortedLatencies, latencies)

	// 简单的排序实现
	for i := 0; i < len(sortedLatencies)-1; i++ {
		for j := i + 1; j < len(sortedLatencies); j++ {
			if sortedLatencies[i] > sortedLatencies[j] {
				sortedLatencies[i], sortedLatencies[j] = sortedLatencies[j], sortedLatencies[i]
			}
		}
	}

	// 计算百分位延迟
	p95Index := int(float64(len(sortedLatencies)) * 0.95)
	p99Index := int(float64(len(sortedLatencies)) * 0.99)

	p95Latency := sortedLatencies[p95Index]
	p99Latency := sortedLatencies[p99Index]

	// 计算每秒操作数
	opsPerSecond := float64(totalOps) / totalTime.Seconds()

	// 获取内存使用情况
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	memoryUsageMB := float64(m.Alloc) / 1024 / 1024

	// 将时间单位转换为毫秒
	totalTimeMs := float64(totalTime) / float64(time.Millisecond)
	avgLatencyMs := float64(avgLatency) / float64(time.Millisecond)
	minLatencyMs := float64(minLatency) / float64(time.Millisecond)
	maxLatencyMs := float64(maxLatency) / float64(time.Millisecond)
	p95LatencyMs := float64(p95Latency) / float64(time.Millisecond)
	p99LatencyMs := float64(p99Latency) / float64(time.Millisecond)

	return &BenchmarkResult{
		TestName:      testName,
		TotalOps:      totalOps,
		TotalTime:     totalTimeMs,
		OpsPerSecond:  opsPerSecond,
		AvgLatency:    avgLatencyMs,
		MinLatency:    minLatencyMs,
		MaxLatency:    maxLatencyMs,
		P95Latency:    p95LatencyMs,
		P99Latency:    p99LatencyMs,
		ErrorCount:    errorCount,
		MemoryUsageMB: memoryUsageMB,
	}
}

// 生成测试报告
func (r *BenchmarkRunner) generateReport(results map[string]*BenchmarkResult, filename string) {
	// 控制台输出
	fmt.Printf("\n=== 性能基准测试报告 ===\n")
	for testName, result := range results {
		if result.TotalOps == 0 {
			fmt.Printf("\n%s: 无有效数据\n", testName)
			continue
		}

		fmt.Printf("\n%s:\n", testName)
		fmt.Printf("  总操作数: %d\n", result.TotalOps)
		fmt.Printf("  总时间: %.2f ms\n", result.TotalTime)
		fmt.Printf("  每秒操作数: %.2f\n", result.OpsPerSecond)
		fmt.Printf("  平均延迟: %.2f ms\n", result.AvgLatency)
		fmt.Printf("  最小延迟: %.2f ms\n", result.MinLatency)
		fmt.Printf("  最大延迟: %.2f ms\n", result.MaxLatency)
		fmt.Printf("  P95延迟: %.2f ms\n", result.P95Latency)
		fmt.Printf("  P99延迟: %.2f ms\n", result.P99Latency)
		fmt.Printf("  错误数: %d\n", result.ErrorCount)
		fmt.Printf("  内存使用: %.2f MB\n", result.MemoryUsageMB)
	}

	// JSON 文件输出
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("创建报告文件失败: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		log.Printf("写入报告文件失败: %v", err)
	}
}
