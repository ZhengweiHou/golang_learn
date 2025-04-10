module gospringdemo

go 1.24

toolchain go1.24.0

//require github.com/go-spring/spring-core v1.1.4-0.20250403000933-f34afe1ae1d0
//replace github.com/go-spring/spring-core v1.1.4 => /home/houzw/document/git-rep/golang/spring-core
replace github.com/go-spring/spring-core v1.1.4 => /home/houzw/document/git-rep/HOUZW/golang/spring-core

require github.com/go-spring/spring-core v1.1.4

require (
	github.com/expr-lang/expr v1.16.9 // indirect
	github.com/magiconair/properties v1.8.9 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	go.uber.org/mock v0.5.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
