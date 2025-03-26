package log

import (
	"fmt"
	"log/slog"
	"testing"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestSlog4zap(t *testing.T) {
	//zapLogger := newZap1()
	//zapLogger := newZap2()
	zapLogger := newZap3()
	defer zapLogger.Sync()

	// 创建一个基于 zap 的 slog 处理器
	handler := NewZapHandler(zapLogger)
	logger := NewZapSlog(handler)
	logger.Info("This is an info log", "key", "value")

	// 使用 slog 进行日志记录
	slog.Info("This is an info log", "key", "value")
	slog.Warn("This is a warning log", "count", 10)
	//	slog.Error("This is an error log", "err", "something went wrong")

	// 记录带有时间的日志
	slog.Info("Log with time", "time", time.Now())
}

func newZap() *zap.Logger {
	return newZap1()
}

func newZap1() *zap.Logger {
	// 创建一个启用调用者信息的 zap 日志实例
	zapConfig := zap.NewProductionConfig()
	zapConfig.Encoding = "console"
	zapConfig.DisableCaller = false
	zapConfig.EncoderConfig.CallerKey = "caller" // 设置 caller 字段名
	zapLogger, _ := zapConfig.Build()
	return zapLogger
}

func newZap2() *zap.Logger {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	zapConfig := &zap.Config{
		//Level:             zap.NewAtomicLevelAt(zapcore.Level(level)),
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       true,
		DisableCaller:     false,
		DisableStacktrace: true,
		Sampling:          &zap.SamplingConfig{Initial: 100, Thereafter: 100},
		Encoding:          "json",
		EncoderConfig:     encoderConfig,
		OutputPaths:       []string{"stderr"},
		ErrorOutputPaths:  []string{"stderr"},
	}
	l, err := zapConfig.Build(zap.AddCallerSkip(3))
	if err != nil {
		fmt.Printf("zap build logger fail err=%v", err)
		return nil
	}
	return l
}

func newZap3() *zap.Logger {
	conf := viper.New()
	conf.SetDefault("hello", "hello")
	//	conf.SetDefault("log.log_file_name", "hzwzaptest.log")
	conf.SetDefault("log.log_level", "debug")
	// conf.SetDefault("log.encoding", "console")
	//conf.SetDefault("env", "prod")

	zaplog := NewZapLog(conf)
	return zaplog
}
