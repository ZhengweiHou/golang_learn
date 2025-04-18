package hzwcache

import (
	"context"
	"errors"
)

var (
	// ErrNoSuchEngine 表示不支持的缓存引擎类型
	ErrNoSuchEngine = errors.New("no such cache engine")
	// constructors    = make(map[string]func(*Config) (ICache[any], error))
)

// fixEmpty 设置默认配置
func (cfg *Config) fixEmpty() {
	if cfg.MaxSizeInMB <= 0 {
		cfg.MaxSizeInMB = 100 // 默认100MB
	}
}

// ICache 泛型缓存接口抽象
type ICache[T any] interface {
	// Get 获取key对应的值
	Get(key string) (T, bool)

	// GetWithLoader 获取key对应的值，如果不存在则调用loader加载
	GetWithLoader(ctx context.Context, key string, loader func() (T, error)) (T, error)

	// Set 设置key-value对
	Set(key string, value T) error

	// SetWithExpire 设置带过期时间的key-value对
	SetWithExpire(key string, value T, ttl int64) error

	// Del 删除指定key
	Del(key string) error

	// Clear 清空所有缓存
	Clear() error
}

type Config struct {
	MaxSizeInMB int64 // 缓存最大大小(MB)
}

// Option 配置选项函数类型
type Option func(cfg *Config)

// WithMaxSize 设置缓存最大大小(MB)
func WithMaxSize(maxSize int64) Option {
	return func(cfg *Config) {
		cfg.MaxSizeInMB = maxSize
	}
}

// NewCache 创建泛型缓存实例
func NewCache[T any](engine string, opts ...Option) (ICache[T], error) {

	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}

	cfg.fixEmpty()

	switch engine {
	case "bigcache":
		return NewBigCache[T](cfg)
	default:
		return nil, ErrNoSuchEngine
	}

	// if f := GetConstructor(engine); f != nil {
	// 	cache, err := f(cfg)
	// 	if err != nil {
	// 		var zero ICache[T]
	// 		return zero, err
	// 	}
	// 	return cache.(ICache[T]), nil
	// }

	// var zero ICache[T]
	// return zero, ErrNoSuchEngine
}
