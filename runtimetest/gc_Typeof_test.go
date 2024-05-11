package runtimetest

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"testing"
	"time"
)

var typeofTypes []reflect.Type
var values [][]byte

func TestTypeof_gc(t *testing.T) {
	typeofTypes = make([]reflect.Type, 0)
	values = make([][]byte, 0)

	// 创建一个输出文件
	bfgc, _ := os.Create("beforgc.heap")
	defer bfgc.Close()
	afgc, _ := os.Create("aftergc.heap")
	defer afgc.Close()
	tracef, _ := os.Create("trace.out")

	trace.Start(tracef)
	defer trace.Stop()

	for i := 0; i < 2048; i++ {
		bt := make([]byte, 1024)
		typeofTypes = append(typeofTypes, reflect.TypeOf(bt))
		values = append(values, bt)
	}
	pprof.WriteHeapProfile(bfgc)
	runtime.GC()
	time.Sleep(time.Second)
	pprof.WriteHeapProfile(afgc)

	fmt.Printf("%d", len(typeofTypes))
	fmt.Printf("%d", len(typeofTypes))
}
