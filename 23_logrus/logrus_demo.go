package main

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	err := initLog()
	if err != nil {
		log.Fatal("日志初始化失败", err.Error())
	}

	logrus.Info("info hello")
	logrus.Debug("debug hello")

	logrus.Debugf("=== %v", "debug \n  hello")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("hello")

	log1 := logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	})
	log1.Info("info hello")

}

func initLog() error {
	logrus.SetLevel(logrus.DebugLevel)
	// logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
		// DisableColors: false,
		FullTimestamp: true,
	})
	logrus.AddHook(NewContextHook(logrus.DebugLevel))
	return nil
}

func NewContextHook(levels ...logrus.Level) logrus.Hook {
	hook := contextHook{
		Field:  "line",
		Skip:   8,
		Depth:  3,
		levels: levels,
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}
	return &hook
}

type contextHook struct {
	Field  string
	Skip   int
	Depth  int
	levels []logrus.Level
}

// Levels implement levels
func (hook contextHook) Levels() []logrus.Level {
	return hook.levels
}

// Fire implement fire
func (hook contextHook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.Field] = findCaller(hook.Skip, hook.Depth)
	return nil
}

func findCaller(skip int, depth int) string {
	file := ""
	line := 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.Contains(file, "logrus@") {
			break
		}
	}

	n := 0
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			n++
			if n >= depth {
				file = file[i+1:]
				break
			}
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	return file, line
}

// -----------------------------------
// ©著作权归作者所有：来自51CTO博客作者wg_FBhEBGaB的原创作品，请联系作者获取转载授权，否则将追究法律责任
// go包之logrus显示日志文件与行号
// https://blog.51cto.com/u_10624715/3234031
