package hello

import (
	"log/slog"
	"reflect"
	"time"
)

type AopHandler func(method string, args []interface{}, result []interface{})

type GenericProxy struct {
	target   interface{}
	handler  AopHandler
	methodFn map[string]reflect.Value
}

func NewGenericProxy(target interface{}, handler AopHandler) interface{} {
	proxy := &GenericProxy{
		target:   target,
		handler:  handler,
		methodFn: make(map[string]reflect.Value),
	}

	t := reflect.TypeOf(target)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		proxy.methodFn[method.Name] = method.Func
	}

	// 特殊处理IHService接口
	if _, ok := target.(IHService); ok {
		return &ihServiceProxy{proxy}
	}
	return proxy
}

type ihServiceProxy struct {
	proxy *GenericProxy
}

func (p *ihServiceProxy) HelloSvs(msg string) string {
	result := p.proxy.Invoke("HelloSvs", msg)
	return result[0].(string)
}

func (p *GenericProxy) Invoke(method string, args ...interface{}) []interface{} {
	start := time.Now()
	slog.Info("AOP Before", "method", method, "args", args)

	// Convert args to []reflect.Value
	in := make([]reflect.Value, len(args)+1)
	in[0] = reflect.ValueOf(p.target)
	for i, arg := range args {
		in[i+1] = reflect.ValueOf(arg)
	}

	// Call method
	result := p.methodFn[method].Call(in)

	// Convert result to []interface{}
	out := make([]interface{}, len(result))
	for i, r := range result {
		out[i] = r.Interface()
	}

	slog.Info("AOP After",
		"method", method,
		"duration", time.Since(start),
		"result", out,
	)
	return out
}
