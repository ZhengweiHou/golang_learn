package util

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/davecgh/go-spew/spew"
)

var config = spew.ConfigState{
	MaxDepth: 15, // 限制递归深度
	Indent:   "  ",
}

func TestLog(msg string, obj any) {

	// if !slog.Default().Enabled(nil, slog.LevelDebug) {
	// 	return
	// }

	pijson, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		msgstring := config.Sdump(obj)
		slog.Info(fmt.Sprintf("%s:%s", msg, msgstring))
	} else {
		slog.Info(fmt.Sprintf("%s:%s", msg, pijson))
	}
}
