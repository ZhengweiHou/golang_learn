package slogdemo

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/golang/glog"
)

func TestJSONHandler(t *testing.T) {
	// 创建一个buffer来捕获日志输出
	var buf bytes.Buffer

	// 创建JSONHandler
	handler := NewJSONHandler(&buf)

	// 创建logger
	logger := slog.New(handler)

	// 记录一些日志
	logger.Info("测试信息", "key1", "value1", "key2", 123)
	logger.Warn("测试警告", "user", "admin", "count", 42)
	logger.Error("测试错误", "err", "something went wrong")

	// 打印捕获的日志输出
	t.Logf("日志输出:\n%s", buf.String())

	// 检查输出是否包含预期的内容
	if !bytes.Contains(buf.Bytes(), []byte("测试信息")) {
		t.Error("日志输出不包含预期的消息")
	}
}

func TestSlogHello(t *testing.T) {
	// 记录一些日志
	slog.Info("测试信息", "key1", "value1", "key2", 123)
	slog.Warn("测试警告", "user", "admin", "count", 42)
	slog.Error("测试错误", "err", "something went wrong")

	hlog := slog.Default().WithGroup("hello")
	hlog.Info("测试信息", "key1", "value1", "key2", 123)
	hlog.Warn("测试警告", "user", "admin", "count", 42)
	hlog.Error("测试错误", "err", "something went wrong")

	glog.Infoln("测试信息", "key1", "value1", "key2", 123)
	// glog.Warn("测试警告", "user", "admin", "count", 42)
	glog.Errorln("测试错误", "err", "something went wrong")
}
