package main

//line <generated>:1
import _ "otel_demo/otel_pkg/rules"

//line <generated>:1
import (
	_ "go.opentelemetry.io/otel"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	_ "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	_ "go.opentelemetry.io/otel/exporters/prometheus"
)

//line main.go:3:1
import (
	// "internal/runtime/exithook"
	"os"
	"otel_demo/hellodo"

	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	os.Setenv("hello", "world")

	// 1
	hello := &hellodo.Hello{}
	hello.HelloOTEL3("hello otel3")

	trace.NewTracerProvider()

	os.Exit(0)
	// main.go:4:2: use of internal package internal/runtime/exithook not allowed
	// exithook.Add(exithook.Hook{
	// 	F: func() {
	// 		// 2
	// 		hello.HelloOTEL3("hello otel3")
	// 	},
	// })
}
