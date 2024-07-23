package runtimegc

import (
	"runtime"
	"runtime/debug"
	"testing"
)

func TestGcWithParm(t *testing.T) {
	debug.SetGCPercent(-1)                       // GOGC=off
	debug.SetMemoryLimit(1024 * 1024 * 1024 * 5) // GOMEMLIMIT=5GiB  限制程序内存使用，单位bytes
	runtime.GOMAXPROCS(15)                       // GOMAXPROCS=15
	runtime.SetBlockProfileRate(1)               // 开启对阻塞操作的跟踪，block
	runtime.SetMutexProfileFraction(1)           // 开启对锁调用的跟踪，mutex
	// debug.SetMaxThreads(20)

	DoGcTest()
}
