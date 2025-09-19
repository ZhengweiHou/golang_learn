package slogdemo

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"time"
)

// JSONHandler 是一个实现slog.Handler接口的JSON格式日志处理器
type JSONHandler struct {
	w io.Writer
}

// NewJSONHandler 创建一个新的JSONHandler
func NewJSONHandler(w io.Writer) *JSONHandler {
	return &JSONHandler{w: w}
}

// Handle 处理日志记录
func (h *JSONHandler) Handle(_ context.Context, r slog.Record) error {
	// 构建日志条目
	entry := map[string]interface{}{
		"time":    r.Time.Format(time.RFC3339),
		"level":   r.Level.String(),
		"message": r.Message,
	}

	// 添加属性
	r.Attrs(func(attr slog.Attr) bool {
		entry[attr.Key] = attr.Value.Any()
		return true
	})

	// 编码为JSON
	enc := json.NewEncoder(h.w)
	enc.SetIndent("", "  ")
	return enc.Encode(entry)
}

// Enabled 返回是否处理指定级别的日志
func (h *JSONHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return true
}

// WithAttrs 返回一个带有新属性的新handler
func (h *JSONHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// 简化实现，实际使用时可能需要更复杂的处理
	return h
}

// WithGroup 返回一个带有新组的新handler
func (h *JSONHandler) WithGroup(name string) slog.Handler {
	// 简化实现，实际使用时可能需要更复杂的处理
	return h
}
