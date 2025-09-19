package test

import (
	"context"
	"log/slog"
	"os"
	"testing"

	slogmulti "github.com/samber/slog-multi"
)

func TestSlogMultiRoute(t *testing.T) {
	route := slogmulti.Router()
	mf1 := func(ctx context.Context, r slog.Record) bool {
		level := r.Level
		if level == slog.LevelInfo {
			return true
		}
		return false
	}

	route = route.Add(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}),
		mf1,
	)

	handler := route.Handler()
	loggs := slog.New(handler)

	_loginfo(loggs)

}

func TestSlogGroup(t *testing.T) {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	logg1 := logger.WithGroup("hzw_group1")

	logg1 = logg1.With("hzw_key1", "hzw_val1")

	// _loginfo(logger)
	_loginfo(logg1)

}
