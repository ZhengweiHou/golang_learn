package hzwcache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/allegro/bigcache/v3"
)

type HBigCache[T any] struct {
	cache *bigcache.BigCache
}

// func init() {
// 	RegisterEngine("bigcache", func(cfg *Config) (ICache[any], error) {
// 		return NewBigCache[any](cfg)
// 	})
// }

func NewBigCache[T any](cfg *Config) (ICache[T], error) {
	if cfg.MaxSizeInMB <= 0 {
		return nil, errors.New("MaxSizeInMB must be positive")
	}
	if cfg.MaxSizeInMB > 1<<31-1 {
		return nil, errors.New("MaxSizeInMB too large")
	}

	// TODO bigcache配置项调整取自Config，并增加相关option
	bcfg := bigcache.Config{
		Shards:             1024,
		LifeWindow:         0, // 永不过期
		CleanWindow:        0, // 不自动清理
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		HardMaxCacheSize:   int(cfg.MaxSizeInMB),
		StatsEnabled:       true,
		Verbose:            true,
	}

	cache, initErr := bigcache.New(context.Background(), bcfg)
	if initErr != nil {
		return nil, initErr
	}

	hbc := &HBigCache[T]{
		cache: cache,
	}
	return hbc, nil
}

func (hbc *HBigCache[T]) Get(key string) (T, bool) {
	var zero T
	entry, err := hbc.cache.Get(key)
	if err != nil {
		if errors.Is(err, bigcache.ErrEntryNotFound) {
			return zero, false
		}
		return zero, false
	}

	var value T
	switch any(zero).(type) {
	case string:
		value = any(string(entry)).(T)
	case bool:
		// bool类型特殊处理
		b, err := strconv.ParseBool(string(entry))
		if err != nil {
			return zero, false
		}
		value = any(b).(T)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		// 数值类型直接类型断言
		value = any(entry).(T)
	default:
		// 复杂类型需要反序列化
		if err := json.Unmarshal(entry, &value); err != nil {
			return zero, false
		}
	}
	return value, true
}

func (hbc *HBigCache[T]) GetWithLoader(ctx context.Context, key string, loader func() (T, error)) (T, error) {
	if val, ok := hbc.Get(key); ok {
		return val, nil
	}

	val, err := loader()
	if err != nil {
		var zero T
		return zero, err
	}

	if err := hbc.Set(key, val); err != nil {
		var zero T
		return zero, err
	}
	return val, nil
}

func (hbc *HBigCache[T]) Set(key string, value T) error {
	var data []byte
	switch v := any(value).(type) {
	case string:
		data = []byte(v)
	case bool:
		// bool类型特殊处理
		data = []byte(strconv.FormatBool(v))
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		// 数值类型直接存储
		data = []byte(fmt.Sprintf("%v", v))
	default:
		// 复杂类型需要序列化
		var err error
		data, err = json.Marshal(value)
		if err != nil {
			return err
		}
	}
	return hbc.cache.Set(key, data)
}

func (hbc *HBigCache[T]) SetWithExpire(key string, value T, ttl int64) error {
	// BigCache不支持单个key的TTL，统一使用全局LifeWindow
	return hbc.Set(key, value)
}
func (hbc *HBigCache[T]) Del(key string) error {
	return hbc.cache.Delete(key)
}

func (hbc *HBigCache[T]) Clear() error {
	return hbc.cache.Reset()
}
