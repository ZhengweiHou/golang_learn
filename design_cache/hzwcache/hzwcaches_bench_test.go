package hzwcache

import (
	"log/slog"
	"testing"
)

func TestXxx(t *testing.T) {

	// 创建字符串类型缓存
	strCache, err := NewCache[string]("bigcache", WithMaxSize(500))
	if err != nil {
		t.Fatal(err)
	}
	strCache.Set("name", "张三")

	// 创建结构体类型缓存
	type User struct {
		Name string
		Age  int
	}
	userCache, err := NewCache[User]("bigcache")
	userCache.Set("u1", User{"李四", 30})

	user, ok := userCache.Get("u1")
	if !ok {
		t.Fatal("get u1 failed")
	}
	slog.Info("user", "user", user)

	boolCache, err := NewCache[bool]("bigcache", WithMaxSize(500))
	if err != nil {
		t.Fatal(err)
	}
	boolCache.Set("b1", true)
	b, ok := boolCache.Get("b1")
	slog.Info("b1", "b1", b, "ok", ok)

}
